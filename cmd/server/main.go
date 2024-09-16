package main

import (
	"flag"
	"log"

	"ipsec-port-forward/internal/server"
)

func main() {
	listenAddr := flag.String("listen", ":8000", "Address to listen on")
	flag.Parse()

	s, err := server.NewServer(*listenAddr)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	log.Printf("Starting server on %s", *listenAddr)
	err = s.Start()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
