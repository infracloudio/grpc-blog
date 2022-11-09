package main

import (
	"context"
	"fmt"
	greetpb "github.com/infracloudio/grpc-blog/greet_app/internal/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
)

const SSL_CA_CERT_PATH = "internal/ssl/ca.crt"

func main() {
	fmt.Println("Starting greet client")

	opts := grpc.WithInsecure()

	certFilePath := SSL_CA_CERT_PATH
	if len(os.Getenv("SSL_CA_CERT_PATH")) > 0 {
		certFilePath = os.Getenv("SSL_CA_CERT_PATH")
	}
	certFile := certFilePath
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("Failed to loading ca trust certificates: %v", sslErr)
		return
	}

	opts = grpc.WithTransportCredentials(creds)

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
