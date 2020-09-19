package identity

// PrivateKeyIdentity is an Identity with Private Key authentication method
type PrivateKeyIdentity struct {
	username string
	key      []byte
}

// NewPrivateKeyIdentity creates a new identity with username u and private key k
func NewPrivateKeyIdentity(u string, k []byte) (*PrivateKeyIdentity, error) {
	pki := &PrivateKeyIdentity{
		username: u,
		key:      k,
	}

	return pki, nil
}

// GetUsername returns the username used to create pki
func (pki *PrivateKeyIdentity) GetUsername() string {
	return pki.username
}

// GetSecret returns the key used to create pki
func (pki *PrivateKeyIdentity) GetSecret() []byte {
	return pki.key
}
