package main

import (
	"flag"
	"log"

	"ipsec-port-forward/internal/client"
)

func main() {
	serverAddr := flag.String("server", "", "Server address")
	localPort := flag.Int("local", 0, "Local port to forward")
	remotePort := flag.Int("remote", 0, "Remote port to forward to")
	flag.Parse()

	if *serverAddr == "" || *localPort == 0 || *remotePort == 0 {
		log.Fatal("Please provide server address, local port, and remote port")
	}

	c, err := client.NewClient(*serverAddr)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	err = c.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer c.Close()

	err = c.ForwardPort(*localPort, *remotePort)
	if err != nil {
		log.Fatalf("Failed to set up port forwarding: %v", err)
	}

	log.Printf("Port forwarding established: local:%d -> remote:%d", *localPort, *remotePort)
	select {} // Keep the program running
}
