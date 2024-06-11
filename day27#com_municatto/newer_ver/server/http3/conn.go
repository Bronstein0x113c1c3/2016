package http3

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"net"

	grpcquic "github.com/coremedic/grpc-quic"
	"github.com/quic-go/quic-go"
)

func tls_setup(p string) (*tls.Config, error) {
	// Generate a new RSA key
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Printf("failed to generate RSA key: %s", err)
		return nil, err
	}

	// Create a certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		// Set other fields of the certificate as required
	}

	// Create a certificate using the template
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		log.Printf("failed to create certificate: %s", err)
		return nil, err
	}

	// Encode the certificate and key to PEM format
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	// Load the X509 key pair from PEM blocks
	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Printf("failed to load X509 key pair from PEM: %s", err)
		return nil, err
	}
	protos := []string{}
	protos = append(protos, p)
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   protos,
	}, nil
}
func NewConn(p string, endpoint string) (net.Listener, error) {
	tls, err := tls_setup(p)
	if err != nil {
		return nil, err
	}
	quic_tunnel, err := quic.ListenAddr(endpoint, tls, nil)
	if err != nil {
		return nil, err
	}
	listener := grpcquic.Listen(*quic_tunnel)
	return listener, nil
}
