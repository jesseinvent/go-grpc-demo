## Go gRPC Demo

Basic demo of the gRPC client - server communication modes in Go.

### gRPC client - server communication modes:

- Unary (request-request RPC)
- Client-side streaming
- Server-side streaming
- Client-server bi-directional streaming

## Setting up demo

Clone repo:

```
$ git clone https://github.com/jesseinvent/go-grpc-demo
```

Install Go Protoc Compiler (Mac)

```
$ brew install protoc-gen-go
```

```
$ brew install protoc-gen-go-grpc
```

[Click to download for other OS](https://grpc.io/docs/protoc-installation/)

Compile Proto file:

```
$ protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

Install packages:

```
$ go mod download
```

```
$ go mod tidy
```

Run server:
```
$ cd server
$ go run main.go
```

Run client:
```
$ cd client
$ go run main.go
```

`Note: Run both server and client processes on different terminal windows` 