## Go GRPC Demo

- Install Go Protoc Compiler (mac)
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
$ go mod tidy
```