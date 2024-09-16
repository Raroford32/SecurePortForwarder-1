#!/bin/bash

set -e

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    echo "You can download Go from https://golang.org/dl/"
    exit 1
fi

# Function to get user input
get_input() {
    read -p "$1: " value
    echo $value
}

# Install dependencies
echo "Installing dependencies..."
go get -u golang.org/x/crypto/...

# Function to configure and create run script for server
configure_server() {
    echo "Configuring server..."
    listen_addr=$(get_input "Enter address to listen on (e.g., :8000)")

    echo "#!/bin/bash" > run_server.sh
    echo "bin/server -listen $listen_addr" >> run_server.sh
    chmod +x run_server.sh

    echo "Server configuration completed."
    echo "To start the server, run: ./run_server.sh"
}

# Function to configure and create run script for client
configure_client() {
    echo "Configuring client..."
    server_addr=$(get_input "Enter server address (e.g., 192.168.1.100:8000)")
    local_port=$(get_input "Enter local port to forward")
    remote_port=$(get_input "Enter remote port to forward to")

    echo "#!/bin/bash" > run_client.sh
    echo "bin/client -server $server_addr -local $local_port -remote $remote_port" >> run_client.sh
    chmod +x run_client.sh

    echo "Client configuration completed."
    echo "To start the client, run: ./run_client.sh"
}

# Main menu
echo "Welcome to the IPsec Port Forwarding System Setup"
echo "Please choose an option:"
echo "1. Install Server"
echo "2. Install Client"
read -p "Enter your choice (1 or 2): " choice

case $choice in
    1)
        echo "Installing server..."
        go build -o bin/server cmd/server/main.go
        configure_server
        ;;
    2)
        echo "Installing client..."
        go build -o bin/client cmd/client/main.go
        configure_client
        ;;
    *)
        echo "Invalid choice. Exiting."
        exit 1
        ;;
esac

echo "Installation and configuration completed successfully."
echo "To start the application, run the appropriate script (run_server.sh or run_client.sh)."
