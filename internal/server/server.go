package server

import (
	"fmt"
	"net"

	"ipsec-port-forward/internal/ipsec"
	"ipsec-port-forward/internal/portforward"
	"ipsec-port-forward/internal/utils"
)

type Server struct {
	listenAddr string
	ipsec      *ipsec.IPSec
}

func NewServer(listenAddr string) (*Server, error) {
	return &Server{
		listenAddr: listenAddr,
		ipsec:      ipsec.NewIPSec(),
	}, nil
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			utils.LogError("Failed to accept connection", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	err := s.ipsec.EstablishConnection(conn)
	if err != nil {
		utils.LogError("Failed to establish IPsec connection", err)
		return
	}

	for {
		remoteAddr, err := utils.ReadString(conn)
		if err != nil {
			utils.LogError("Failed to read remote address", err)
			return
		}

		remoteConn, err := net.Dial("tcp", remoteAddr)
		if err != nil {
			utils.LogError("Failed to connect to remote address", err)
			continue
		}

		go portforward.Forward(conn, remoteConn)
	}
}