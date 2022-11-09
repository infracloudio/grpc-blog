package main

import (
	"context"
	"fmt"
	greetpb "github.com/infracloudio/grpc-blog/greet_app/internal/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting greet client")

	opts := grpc.WithInsecure()

	servAddr := ":50059"
	if len(os.Getenv("GRPC_SRV_ADDR")) > 0 {
		servAddr = os.Getenv("GRPC_SRV_ADDR")
	}
	cc, err := grpc.Dial(servAddr, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Println("starting unary.....")
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
