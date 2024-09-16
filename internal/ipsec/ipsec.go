package ipsec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"net"
)

type IPSec struct {
	key   []byte
	block cipher.Block
}

func NewIPSec() *IPSec {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate key: %v", err))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(fmt.Sprintf("Failed to create AES cipher: %v", err))
	}

	return &IPSec{
		key:   key,
		block: block,
	}
}

func (i *IPSec) EstablishConnection(conn net.Conn) error {
	// In a real implementation, this would involve a key exchange protocol
	// For simplicity, we'll just use a pre-shared key
	return nil
}

func (i *IPSec) Encrypt(plaintext []byte) ([]byte, error) {
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(i.block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func (i *IPSec) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(i.block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
