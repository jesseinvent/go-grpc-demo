package main

import (
	"io"
	"time"
	"fmt"
	"log"
	"net"
	"context"

	pb "github.com/jesseinvent/go-grpc-demo/proto"

	"google.golang.org/grpc"
)

const(
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", port);

	if err != nil {
		log.Fatalf("Failed to start the server %v", err);
	}

	grpcServer := grpc.NewServer();

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{});

	fmt.Printf("server started at %s\n", listener.Addr());

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err);
	}

}

// Unary
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
	
	return &pb.HelloResponse {
		Message: "Hello",
	}, nil;
}

// Server streaming
func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	fmt.Printf("got request with names: %v", req.Names);

	for _, name := range req.Names {
		res := &pb.HelloResponse{Message: fmt.Sprintf("Hello %v", name)};

		if err := stream.Send(res); err != nil {
			return err;
		} 

		time.Sleep(2 * time.Second); 
	}

	return nil;
}

// Client streaming
func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string;

	for {
		req, err := stream.Recv();

		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages});
		}

		if err != nil {
			return err;
		}

		fmt.Printf("Got request with name: %v\n", req.Name);

		messages = append(messages, fmt.Sprintf("Hello %v", req.Name));
	}
}

// Bidirectional stream


func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv();

		if err == io.EOF {
			return nil;
		}

		if err != nil {
			return err;
		}

		fmt.Printf("Received request: %v\n", req.Name);

		res := &pb.HelloResponse{Message: fmt.Sprintf("Hello %v", req.Name)};

		if err := stream.Send(res); err != nil {
			return err;
		}
	}
}
