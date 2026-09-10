package certs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	caCN = "AURA Intercept CA"
	org  = "AURA"
)

var leafMu sync.Mutex

func caDir(root string) (string, error) {
	dir := filepath.Join(root, "ca")
	return dir, os.MkdirAll(dir, 0o755)
}

func loadOrCreateKey(path string) (*rsa.PrivateKey, error) {
	if data, err := os.ReadFile(path); err == nil {
		block, _ := pem.Decode(data)
		if block == nil {
			return nil, fmt.Errorf("invalid key PEM: %s", path)
		}
		if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
			return key, nil
		}
		k, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaKey, ok := k.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("not an RSA key: %s", path)
		}
		return rsaKey, nil
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	if err := os.WriteFile(path, pem.EncodeToMemory(block), 0o600); err != nil {
		return nil, err
	}
	return key, nil
}

func randomSerial() *big.Int {
	n, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return big.NewInt(time.Now().UnixNano())
	}
	if n.Sign() == 0 {
		n = big.NewInt(1)
	}
	return n
}

func EnsureCA(root string) (*x509.Certificate, *rsa.PrivateKey, error) {
	dir, err := caDir(root)
	if err != nil {
		return nil, nil, err
	}
	keyPath := filepath.Join(dir, "ca-key.pem")
	certPath := filepath.Join(dir, "ca.crt")
	key, err := loadOrCreateKey(keyPath)
	if err != nil {
		return nil, nil, err
	}
	if data, err := os.ReadFile(certPath); err == nil {
		block, _ := pem.Decode(data)
		if block == nil {
			return nil, nil, fmt.Errorf("invalid CA cert PEM")
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, nil, err
		}
		return cert, key, nil
	}
	now := time.Now()
	tpl := &x509.Certificate{
		SerialNumber:          randomSerial(),
		Subject:               pkix.Name{CommonName: caCN, Organization: []string{org}},
		NotBefore:             now.Add(-24 * time.Hour),
		NotAfter:              now.Add(3650 * 24 * time.Hour),
		IsCA:                  true,
		MaxPathLen:            1,
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}
	der, err := x509.CreateCertificate(rand.Reader, tpl, tpl, &key.PublicKey, key)
	if err != nil {
		return nil, nil, err
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, nil, err
	}
	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	if err := os.WriteFile(certPath, pemBytes, 0o644); err != nil {
		return nil, nil, err
	}
	return cert, key, nil
}

func leafPaths(root, hostname string) (string, string, error) {
	var b strings.Builder
	for _, ch := range hostname {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '.' || ch == '-' || ch == '_' {
			b.WriteRune(ch)
		} else {
			b.WriteByte('_')
		}
	}
	safe := b.String()
	if len(safe) > 200 {
		safe = safe[:200]
	}
	if safe == "" {
		safe = "host"
	}
	dir := filepath.Join(root, "leaf")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", "", err
	}
	return filepath.Join(dir, safe+".crt"), filepath.Join(dir, safe+".key"), nil
}

func EnsureLeaf(root, hostname string) (*tls.Certificate, error) {
	leafMu.Lock()
	defer leafMu.Unlock()
	certPath, keyPath, err := leafPaths(root, hostname)
	if err != nil {
		return nil, err
	}
	if _, err1 := os.Stat(certPath); err1 == nil {
		if _, err2 := os.Stat(keyPath); err2 == nil {
			c, err := tls.LoadX509KeyPair(certPath, keyPath)
			if err == nil {
				return &c, nil
			}
		}
	}
	caCert, caKey, err := EnsureCA(root)
	if err != nil {
		return nil, err
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	cn := hostname
	if len(cn) > 64 {
		cn = cn[:64]
	}
	tpl := &x509.Certificate{
		SerialNumber:          randomSerial(),
		Subject:               pkix.Name{CommonName: cn, Organization: []string{org}},
		NotBefore:             now.Add(-24 * time.Hour),
		NotAfter:              now.Add(825 * 24 * time.Hour),
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	if ip := net.ParseIP(hostname); ip != nil {
		tpl.IPAddresses = []net.IP{ip}
	} else {
		tpl.DNSNames = []string{hostname}
	}
	der, err := x509.CreateCertificate(rand.Reader, tpl, caCert, &key.PublicKey, caKey)
	if err != nil {
		return nil, err
	}
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err := os.WriteFile(certPath, certPEM, 0o644); err != nil {
		return nil, err
	}
	if err := os.WriteFile(keyPath, keyPEM, 0o600); err != nil {
		return nil, err
	}
	c, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func CAPath(root string) string {
	return filepath.Join(root, "ca", "ca.crt")
}
