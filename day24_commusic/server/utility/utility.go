package utility

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/dhowden/tag"
)

func Readinginfo(dir string) {
	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	meta, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	res := ""
	res += fmt.Sprintf("name: %v \n", meta.Title())
	res += fmt.Sprintf("artist name: %v \n", meta.Artist())
	res += fmt.Sprintf("album name: %v \n", meta.Album())
	i, _ := meta.Track()
	res += fmt.Sprintf("track: %v \n", i)
	res += fmt.Sprintf("published year: %v", meta.Year())
	fmt.Println(res)
}

func GenerateTLSConfig(p string) (*tls.Config, error) {
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
