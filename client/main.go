package main

import (
	"io"
	"fmt"
	"log"
	"time"
	"context"

	pb "github.com/jesseinvent/go-grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()));

	if err != nil {
		log.Fatalf("did not connect: %v", err);
	}

	defer conn.Close(); 

	client := pb.NewGreetServiceClient(conn);

	names := &pb.NamesList{
		Names: []string{"Jesse", "John", "Doe"}, 
	} 

	// callSayHello(client);
 
	// callSayHelloServerStreaming(client, names);

	// callSayHelloClientStreaming(client, names);

	callSayHelloBidirectionalStreaming(client, names);
}

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second);

	defer cancel();

	res, err := client.SayHello(ctx, &pb.NoParam{});

	if err != nil {
		log.Fatalf("could not greet: %v", err);
	}

	fmt.Printf("response: %s", res.Message);
}

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {

	fmt.Println("Streaming started");

	stream, err := client.SayHelloServerStreaming(context.Background(), names);

	if err != nil {
		log.Fatalf("could not send names: %v", err);
	}

	for {
		message, err := stream.Recv();

		if err == io.EOF {
			break;
		}

		if err != nil {
			log.Fatalf("error while streaming: %v", err); 
		}

		fmt.Printf(message.Message);
	}

	fmt.Println("Streaming finished..");
}

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {

	fmt.Println("Client streaming started");

	stream, err := client.SayHelloClientStreaming(context.Background());

	if err != nil {
		log.Fatalf("could not send names: %v", err);
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest {
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending: %v", err);
		}

		fmt.Printf("Send the request with name: %s\n", name);

		time.Sleep(2 * time.Second);
	}

	res, err := stream.CloseAndRecv();

	fmt.Println("Client streaming finished"); 

	if err != nil {
		log.Fatalf("error while receiving: %v", err);
	}

	fmt.Println(res.Messages);
}


func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	fmt.Println("Bidirectional Streaming started");

	stream, err := client.SayHelloBidirectionalStreaming(context.Background());

	if err != nil {
		log.Fatalf("could not send names %v", err);
	}

	channel := make(chan struct{});

	go func() {
		for {
			message, err := stream.Recv();

			if err == io.EOF {
				break;
			}

			if err != nil {
				log.Fatalf("error while streaming %v", err);
			}

			fmt.Printf("Recieved response: %v\n", message.Message);
		}

		close(channel);
	}();

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending request: %v", err);
		}

		time.Sleep(2 * time.Second); 
	}

	stream.CloseSend();

	<-channel;

	fmt.Println("Birectional Streaming finshed..");
}