package main

import (
	"fmt"
	"log"

	"github.com/hacdan/receipt-processor-challenge/api"
)

func main() {
	listenAddr := ":8080"
	server := api.NewServer(listenAddr)
	fmt.Printf("Starting server on localhost%s\n", listenAddr)
	log.Fatal(server.Start())
}
