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

#### Mobile clients
Built with React Native.

*Take one*
Tried to connect a RN application straight to a backing gRPC service. Didn't work, apparently RN is not node (duh) and gRPC node module needs some dependencies (e.g. path) that doesn't exist in RN land. Refs [Github](https://github.com/grpc/grpc/issues/7038) & [Stack overflow](http://stackoverflow.com/questions/36203549/grpc-on-react-native). It still might work though if I use the native gRPC clients for each platform but that would defeat the purpose of sharing logic in the first place. I
guess I need to find out how much RN and node deviates to see if there is any idea in sharing logic between web and mobile clients.

### Client <-> Public API
- Probably will be JSON over HTTP 1.1 for compatability reasons.
- Maybe also gRPC, possibly through grpc-gateway project.
- Sessions/auth via JWTs.

### Public API (Perimeter)
Should provide a single entrypoint, both for simplicity of clients and to minimize attack surface.  Should handle authN/Z, ratelimiting and such, everything to keep internal services as light weight as possible. Sometimes even just proxy requests to internal services.

### Public API <-> Internal services & Internal services <-> Internal services
gRPC (http2 with protobuf) payloads to reduce the additional cost of decoupling and spending time on the network between services instead of within app bounds (in a monolithic architecture). gRPC is also natively language independent.
- Encrypted communication via mutual TLS
- Possibly signed JWTs to grant/limit access to services and know more about the receiving part

### Internal services
Should be kept as small as possible, avoiding unneccesary abstractions and many layers. Possibly tested via it's exposed gRPC API instead of e.g. unit testing to keep implementations even simpler. Should fit in your head as Dan North should have said.
