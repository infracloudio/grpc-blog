package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"strconv"
)

func main() {
	var reqCount int
	if count, ok := os.LookupEnv("REQUEST_COUNT"); ok {
		reqCount, _ = strconv.Atoi(count)
	}

	for i:=0; i<reqCount; i++ {
		makeRequest()
	}
}

func makeRequest() {
	client := resty.New()

	serverHost := "localhost"
	serverPort := "9090"

	if host, ok := os.LookupEnv("SERVER_HOST"); ok {
		serverHost = host
	}

	if port, ok := os.LookupEnv("SERVER_PORT"); ok {
		serverPort = port
	}

	url := fmt.Sprintf("http://%s:%s/", serverHost, serverPort)
	resp, err := client.R().Get(url)

	if err != nil {
		log.Fatalf("error while calling rest api : %v", err)
	}
	log.Printf("reponse from rest api: %v", resp.String())
}