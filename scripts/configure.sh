#!/bin/bash

set -e

# Function to get user input
get_input() {
    read -p "$1: " value
    echo $value
}

# Configure client
configure_client() {
    echo "Configuring client..."
    server_addr=$(get_input "Enter server address (e.g., 192.168.1.100:8000)")
    local_port=$(get_input "Enter local port to forward")
    remote_port=$(get_input "Enter remote port to forward to")

    echo "#!/bin/bash" > run_client.sh
    echo "bin/client -server $server_addr -local $local_port -remote $remote_port" >> run_client.sh
    chmod +x run_client.sh

    echo "Client configuration completed. Run ./run_client.sh to start the client."
}

# Configure server
configure_server() {
    echo "Configuring server..."
    listen_addr=$(get_input "Enter address to listen on (e.g., :8000)")

    echo "#!/bin/bash" > run_server.sh
    echo "bin/server -listen $listen_addr" >> run_server.sh
    chmod +x run_server.sh

    echo "Server configuration completed. Run ./run_server.sh to start the server."
}

# Main menu
while true; do
    echo "Select an option:"
    echo "1. Configure client"
    echo "2. Configure server"
    echo "3. Exit"
    read -p "Enter your choice (1-3): " choice

    case $choice in
        1) configure_client ;;
        2) configure_server ;;
        3) exit 0 ;;
        *) echo "Invalid choice. Please try again." ;;
    esac
done
