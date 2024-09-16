package client

import (
	"fmt"
	"net"

	"ipsec-port-forward/internal/ipsec"
	"ipsec-port-forward/internal/portforward"
	"ipsec-port-forward/internal/utils"
)

type Client struct {
	serverAddr string
	conn       net.Conn
	ipsec      *ipsec.IPSec
}

func NewClient(serverAddr string) (*Client, error) {
	return &Client{
		serverAddr: serverAddr,
		ipsec:      ipsec.NewIPSec(),
	}, nil
}

func (c *Client) Connect() error {
	var err error
	c.conn, err = net.Dial("tcp", c.serverAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %v", err)
	}

	err = c.ipsec.EstablishConnection(c.conn)
	if err != nil {
		return fmt.Errorf("failed to establish IPsec connection: %v", err)
	}

	return nil
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) ForwardPort(localPort, remotePort int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", localPort))
	if err != nil {
		return fmt.Errorf("failed to listen on local port: %v", err)
	}

	go func() {
		for {
			localConn, err := listener.Accept()
			if err != nil {
				utils.LogError("Failed to accept local connection", err)
				continue
			}

			go c.handleConnection(localConn, remotePort)
		}
	}()

	return nil
}

func (c *Client) handleConnection(localConn net.Conn, remotePort int) {
	defer localConn.Close()

	remoteConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.serverAddr, remotePort))
	if err != nil {
		utils.LogError("Failed to connect to remote port", err)
		return
	}
	defer remoteConn.Close()

	portforward.Forward(localConn, remoteConn)
}