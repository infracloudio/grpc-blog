package main

import (
	"context"
	"fmt"
	greetpb "github.com/infracloudio/grpc-blog/grpc-loadbalancing/proto"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting greet client")

	opts := grpc.WithInsecure()
	serverHost := "localhost"
	serverPort := "50051"

	if host, ok := os.LookupEnv("GRPC_SERVER_HOST"); ok {
		serverHost = host
	}

	if port, ok := os.LookupEnv("GRPC_SERVER_PORT"); ok {
		serverPort = port
	}

	servAddr := fmt.Sprintf("dns:///%s:%s", serverHost, serverPort)

	cc, err := grpc.Dial(
		servAddr,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithBlock(),
		opts,
	)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)

	var reqCount int
	if count, ok := os.LookupEnv("REQUEST_COUNT"); ok {
		reqCount, _ = strconv.Atoi(count)
	}

	for i := 0; i < reqCount; i++ {
		doUnary(c)
	}
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hitesh",
			LastName:  "Pattanayak",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet rpc : %v", err)
	}
	log.Printf("reponse from Greet rpc: %v", res.Result)
}
