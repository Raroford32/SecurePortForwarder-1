package utils

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

func LogError(message string, err error) {
	log.Printf("%s: %v", message, err)
}

func WriteString(conn net.Conn, s string) error {
	length := uint16(len(s))
	if err := binary.Write(conn, binary.BigEndian, length); err != nil {
		return err
	}
	_, err := conn.Write([]byte(s))
	return err
}

func ReadString(conn net.Conn) (string, error) {
	var length uint16
	if err := binary.Read(conn, binary.BigEndian, &length); err != nil {
		return "", err
	}
	buffer := make([]byte, length)
	_, err := io.ReadFull(conn, buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}
