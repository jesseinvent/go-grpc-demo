## Go gRPC Demo

### gRPC client - server communication modes:

- Unary (request-request RPC)
- Client-side streaming
- Server-side streaming
- Client-server bi-directional streaming

## Setting up demo

- Clone repo

```
$ git clone https://github.com/jesseinvent/go-grpc-demo
```

- Install Go Protoc Compiler (Mac)

```
$ brew install protoc-gen-go
```

```
$ brew install protoc-gen-go-grpc
```

- Compile Proto file

```
$ protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

- Install packages

```
$ go mod download
```

```
$ go mod tidy
```
