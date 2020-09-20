package client

import (
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/host"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/identity"
	"golang.org/x/crypto/ssh"
)

// Client is a wrapper for an SSH client. The Client dials into a remote server
// and can be used to create an SSH session
type Client struct {
	*ssh.Client
}

// New creates a new SSH client and dials into the remote host. error is not nil if there is an error during this process.
func New(h *host.Host, i identity.Identity) (*Client, error) {
	// determine authmethod based on identity type
	var am ssh.AuthMethod
	switch i.(type) {
	case *identity.PasswordIdentity:
		am = ssh.Password(string(i.GetSecret()))
	case *identity.PrivateKeyIdentity:
		// get signer from raw key
		k, err := ssh.ParsePrivateKey(i.GetSecret())
		if err != nil {
			return nil, err
		}
		am = ssh.PublicKeys(k)
	}

	// Create the client config
	conf := &ssh.ClientConfig{
		User: i.GetUsername(),
		Auth: []ssh.AuthMethod{
			am,
		},
	}
}
