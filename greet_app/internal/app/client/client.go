package main

import (
	"context"
	"fmt"
	greetpb "github.com/infracloudio/grpc-blog/greet_app/internal/pkg/proto"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting greet client")

	rpcCreds := oauth.NewOauthAccess(&oauth2.Token{AccessToken: "client-x-id"})
	trnCreds, err := credentials.NewClientTLSFromFile("internal/app/client/cert/public.crt", "localhost")
	if err != nil {
		log.Fatalln(err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(trnCreds), grpc.WithPerRPCCredentials(rpcCreds)}

	servAddr := ":50059"
	if len(os.Getenv("GRPC_SRV_ADDR")) > 0 {
		servAddr = os.Getenv("GRPC_SRV_ADDR")
	}
	cc, err := grpc.Dial(servAddr, opts...)
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
