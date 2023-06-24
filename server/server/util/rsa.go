package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func generateRSA() ([]byte, []byte, error) {
	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Convert private key to DER format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	// Convert public key to PEM PKCS1 format
	publicKeyBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})
	if publicKeyBytes == nil {
		return nil, nil, errors.New("failed to encode public key to PEM format")
	}

	return privateKeyBytes, publicKeyBytes, nil
}
