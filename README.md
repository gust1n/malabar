# Lab, playground and PoC of a (micro)services architecture
To play, learn and show...

## Goals
Provide a fast, secure and not too hard way of decoupling responsibilities and maximize logic reuse. Should be:
- Language independent
- Easily scalable
- All communication encrypted

## Current ideas

### Clients
Provide both mobile and web clients to cover the most usual needs of modern applications. Use a shared node client core which is then deployed to iOS/Android via react native and web via some node webserver (probably Express).

### Client <-> Public API
- Probably will be JSON over HTTP 1.1 for compatability reasons.
- Sessions/auth via JWTs.

### Public API (Perimeter)
Should provide a single entrypoint, both for simplicity of clients and to minimize attack surface.  Should handle authN/Z, ratelimiting and such, everything to keep internal services as light weight as possible. Sometimes even just proxy requests to internal services.

### Public API <->Internal services & Internal services <-> Internal services
gRPC (http2 with protobuf) payloads to reduce the additional cost of decoupling and spending time on the network between services instead of within app bounds (in a monolithic architecture). gRPC is also natively language independent.

### Internal services
Should be kept as small as possible, avoiding unneccesary abstractions and many layers. Possibly tested via it's exposed gRPC API instead of e.g. unit testing to keep implementations even simpler. Should fit in your head as Dan North should have said.
