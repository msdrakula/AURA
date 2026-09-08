package certs

import (
	"crypto/tls"
	"fmt"
)

// Authority loads or creates a local CA and per-host leaf certificates.
type Authority struct {
	Root string
}

// NewAuthority uses root as the certificate data directory.
func NewAuthority(root string) *Authority {
	return &Authority{Root: root}
}

// Ensure loads an existing CA or writes a new one under Root.
func (a *Authority) Ensure() error {
	if a == nil || a.Root == "" {
		return fmt.Errorf("certs: empty data directory")
	}
	_, _, err := EnsureCA(a.Root)
	if err != nil {
		return fmt.Errorf("failed to ensure CA: %w", err)
	}
	return nil
}

// Leaf returns a server certificate for hostname, minting one if needed.
func (a *Authority) Leaf(hostname string) (*tls.Certificate, error) {
	if a == nil {
		return nil, fmt.Errorf("certs: nil authority")
	}
	crt, err := EnsureLeaf(a.Root, hostname)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure leaf for %s: %w", hostname, err)
	}
	return crt, nil
}

// CAFile is the PEM path served to the UI as the downloadable CA.
func (a *Authority) CAFile() string {
	if a == nil {
		return ""
	}
	return CAPath(a.Root)
}
