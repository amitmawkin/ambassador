from typing import Any, Dict, Iterable, List, Optional, Tuple, TYPE_CHECKING
from typing import cast as typecast

from ..config import Config
from .irbasemapping import IRBaseMapping
from .irbasemappinggroup import IRBaseMappingGroup
from .irhost import IRHost
from .irresource import IRResource
from .irtlscontext import IRTLSContext

if TYPE_CHECKING:
    from .ir import IR


class IRListener (IRResource):
    """
    IRListener is Ambassador's concept of a listener.

    Note that an IRListener is a _very different beast_ than an Envoy listener.
    An IRListener is pretty straightforward right now: a port, whether TLS is
    required, whether we're doing Proxy protocol.

    NOTE WELL: at present, all IRListeners are considered to be equals --
    specifically, every Mapping is assumed to belong to every IRListener. That
    may change later, but that's why you don't see Mappings associated with an
    IRListener when creating the IRListener.

    An Envoy listener, by contrast, is something more akin to an environment
    definition. See V2Listener for more.
    """
    @staticmethod
    def helper_contexts(res: IRResource, k: str) -> Tuple[str, dict]:
        return k, { ctx_key: ctx.as_dict() for ctx_key, ctx in res[k].items() }

    def __init__(self, ir: 'IR', aconf: Config,
                 service_port: int,
                 require_tls: bool,
                 use_proxy_proto: bool,
                 redirect_listener: bool = False,

                 rkey: str="ir.listener",
                 kind: str="IRListener",
                 name: str="ir.listener",
                 **kwargs) -> None:
        # print("IRListener __init__ (%s %s %s)" % (kind, name, kwargs))

        super().__init__(
            ir=ir, aconf=aconf, rkey=rkey, kind=kind, name=name,
            service_port=service_port,
            require_tls=require_tls,
            use_proxy_proto=use_proxy_proto,
            **kwargs)

        self.redirect_listener: bool = redirect_listener
        self.require_tls: bool = require_tls
        self.add_dict_helper('tls_contexts', IRListener.helper_contexts)


class IRListenerSet:
    def __init__(self, ir: 'IR', name: str, host: Optional[IRHost]) -> None:
        self.ir = ir
        self.name = name
        self.host = host
        self.contexts: List[IRTLSContext] = []
        self.listeners: List[IRListener] = []
        self.groups: Dict[str, IRBaseMappingGroup] = {}

        self.ir.logger.info(f"ListenerSet {self.name}: init, host {self.host.name if self.host else '-none-'}")

    def match(self, resource: IRResource) -> bool:
        if not self.host:
            # Good to go.
            self.ir.logger.info(f"ListenerSet {self.name}: fallback match for {resource}")
            return True

        if self.host.matches(resource):
            self.ir.logger.info(f"ListenerSet {self.name}: match for {resource}")
            return True

        self.ir.logger.info(f"ListenerSet {self.name}: no match for {resource}")
        return False

    def consider_context(self, ctx: IRTLSContext) -> bool:
        self.ir.logger.debug(f"ListenerSet {self.name}: consider_context {ctx}")

        if self.match(ctx):
            self.ir.logger.debug(f"ListenerSet {self.name}: add {ctx}")
            self.contexts.append(ctx)
            return True

        return False

    def consider_mapping(self, aconf: Config, mapping: IRBaseMapping) -> bool:
        self.ir.logger.debug(f"ListenerSet {self.name}: consider_mapping {mapping}")

        if self.match(mapping):
            if mapping.group_id not in self.groups:
                group_name = "GROUP: %s" % mapping.name
                group_class = mapping.group_class()
                group = group_class(ir=self.ir, aconf=aconf,
                                    location=mapping.location,
                                    name=group_name,
                                    mapping=mapping)

                self.groups[group.group_id] = group
            else:
                group = self.groups[mapping.group_id]
                group.add_mapping(aconf, mapping)

            return True

        return False

    def ordered_groups(self) -> Iterable[IRBaseMappingGroup]:
        return reversed(sorted(self.groups.values(), key=lambda x: x['group_weight']))

    def as_dict(self) -> Dict[str, Any]:
        return {
            'name': self.name,
            'host': [ self.host.as_dict() if self.host else None ],
            'contexts': [ ctx.as_dict() for ctx in self.contexts ],
            'groups': [ group.as_dict() for group in self.ordered_groups() ],
            'listeners': [ l.as_dict() for l in self.listeners ]
        }


class ListenerFactory:
    @classmethod
    def load_all(cls, ir: 'IR', aconf: Config) -> None:
        amod = ir.ambassador_module

        # OK. First things first: go separate everything out by Host, if that's a thing.
        listener_sets: List[IRListenerSet] = []

        for host in ir.get_hosts() or []:
            listener_sets.append(IRListenerSet(ir, host.name, host))

        listener_sets.append(IRListenerSet(ir, "fallback", None))

        # Split up the TLSContexts.
        for ctx in ir.get_tls_contexts():
            if ctx.is_active():
                found = False

                # Add this context to all the host-based listener_sets that make
                # sense...
                for lset in listener_sets[:-1]:
                    if lset.consider_context(ctx):
                        found = True

                # ...and then to the fallback IFF nothing else matched.
                if not found:
                    listener_sets[-1].consider_context(ctx)

        # OK. Now let's start figuring out how to put each listener_set together.
        for lset in listener_sets:
            # OK, so start by assuming that we're not redirecting cleartext...
            redirect_cleartext_from: Optional[int] = None
            redirection_context: Optional[IRTLSContext] = None

            # ...that we have no TLS termination contexts...
            contexts = {}

            # ...and that the primary listener is being defined by the Ambassador module,
            # rather than by a TLSContext.
            primary_location = amod.location
            primary_context: Optional[IRTLSContext] = None

            # Finally, make a call about whether we'll allow overriding that location later.
            #
            # Here's the story: if there are no Hosts or TLS contexts in play, the primary
            # listener is implicitly defined by the Ambassador module -- that's where
            # the service port is specified, so that's where the user needs to go to
            # make changes to the listener. ("The Ambassador module" might be implicit
            # too, of course, but still, it's the best we have in this case.)
            #
            # Now suppose there is at least one Host or TLS termination context in play.
            # If the user did _not_ supply an Ambassador module, the existence of the context
            # will flip the primary listener from port 8080 to port 8443, and that means
            # that it's the _TLS context_ that's really defining the service port. So
            # we flip the listener's location to the context. Same logic applies for a
            # Host that defines a port (once that becomes a thing).
            #
            # If they supply an Ambassador module _and_ a TLS termination context, we
            # leave the listener's location as the Ambassador module, because in that
            # case it'll be the Ambassador module that wins again.
            #
            # XXX Well. In theory. Need to doublecheck to see what happens if they
            # supplied an Ambassador module that doesn't define the service port...
            override_location = bool(amod.location == '--internal--')

            for ctx in lset.contexts:
                if ctx.kind != 'IRTLSContext':
                    continue

                ctx = typecast(IRTLSContext, ctx)

                if ctx.get('hosts', None):
                    # This is a termination context.
                    contexts[ctx.name] = ctx

                    ctx_kind = "a" if primary_context else "the primary"
                    ir.logger.info(f"ListenerSet {lset.name}: ctx {ctx.name} is {ctx_kind} termination context")

                    primary_context = ctx

                    if override_location:
                        # We'll need to override the location of the primary listener with
                        # this later.
                        primary_location = ctx.location

                    if ctx.get('redirect_cleartext_from', None):
                        # We are indeed redirecting cleartext.
                        redirect_cleartext_from = ctx.redirect_cleartext_from

                        # XXX if 'location'??
                        if 'location' in ctx:
                            redirection_context = ctx

                        ir.logger.debug(f"ListenerSet {lset.name}: ctx {ctx.name} sets redirect_cleartext_from {redirect_cleartext_from}")

            # OK, handle the simple case first: if we're in debug mode, just always
            # fire up multiprotocol listeners on ports 8080 and 8443. This means neither
            # requires TLS, but both have a full set of termination contexts.

            if ir.wizard_allowed:
                ir.logger.info(f"ListenerSet {lset.name}: wizard allowed, going multiprotocol")

                lset.listeners = [
                    IRListener(
                        ir=ir, aconf=aconf, location=amod.location,
                        service_port=8080,
                        require_tls=False,
                        use_proxy_proto=False
                    ),
                    IRListener(
                        ir=ir, aconf=aconf, location=amod.location,
                        service_port=8443,
                        require_tls=False,
                        use_proxy_proto=False
                    )
                ]

                for listener in lset.listeners:
                    # If this listener_set has TLS contexts, all the listeners need to get them
                    # all.
                    #
                    # If we _don't_ have TLS contexts... well, that's kind of interesting
                    # if we're in debug mode, because it shouldn't really happen. Go ahead
                    # and fire things up without any, and trust that the rest of the
                    # system will be supplying something Soon(tm).

                    if contexts:
                        listener['tls_contexts'] = contexts
            else:
                ir.logger.info(f"ListenerSet {lset.name}: setting up primary listener")

                # We're not in debug mode here, so set up our primary listener. If we have TLS termination
                # contexts, we'll attach them to this listener, and it'll do TLS on whatever service port
                # the user has configured.
                primary_listener = IRListener(
                    ir=ir, aconf=aconf, location=primary_location,
                    service_port=amod.service_port,
                    require_tls=amod.get('x_forwarded_proto_redirect', False),
                    use_proxy_proto=amod.use_proxy_proto
                )

                if primary_context:
                    primary_listener.sourced_by(primary_context)

                if contexts:
                    primary_listener['tls_contexts'] = contexts

                # If x_forwarded_proto_redirect is set, then we enable require_tls in primary_listener,
                # which in turn sets require_ssl to true in envoy aconf. Once set, then all requests
                # that contain X-FORWARDED-PROTO set to https, are processes normally by envoy. In all
                # the other cases, including X-FORWARDED-PROTO set to http, a 301 redirect response to
                # https://host is sent
                if primary_listener.require_tls:
                    ir.logger.debug("x_forwarded_proto_redirect is set to true, enabling 'require_tls' in listener")

                if 'use_remote_address' in amod:
                    primary_listener.use_remote_address = amod.use_remote_address

                if 'xff_num_trusted_hops' in amod:
                    primary_listener.xff_num_trusted_hops = amod.xff_num_trusted_hops

                if 'server_name' in amod:
                    primary_listener.server_name = amod.server_name

                lset.listeners.append(primary_listener)

                if redirect_cleartext_from:
                    ir.logger.info(f"ListenerSet {lset.name}: setting up redirect listener")
                    # We're redirecting cleartext. This means a second listener that has no TLS contexts,
                    # and does nothing but redirects.
                    new_listener = IRListener(
                        ir=ir, aconf=aconf, location=redirection_context.location,
                        service_port=redirect_cleartext_from,
                        use_proxy_proto=amod.use_proxy_proto,
                        # Note: no TLS context here, this is a cleartext listener.
                        # We can set require_tls True because we can let the upstream
                        # tell us about that.
                        require_tls=True,
                        redirect_listener=True
                    )

                    if 'use_remote_address' in amod:
                        new_listener.use_remote_address = amod.use_remote_address

                    if 'xff_num_trusted_hops' in amod:
                        new_listener.xff_num_trusted_hops = amod.xff_num_trusted_hops

                    if 'server_name' in amod:
                        new_listener.server_name = amod.server_name

                    lset.listeners.append(new_listener)

            ir.add_listener_set(lset)

    @classmethod
    def finalize(cls, ir: 'IR', aconf: Config) -> None:
        pass
