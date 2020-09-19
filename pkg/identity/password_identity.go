package identity

// PasswordIdentity is an Identity with Private Key authentication method
type PasswordIdentity struct {
	username string
	password []byte
}

// NewPasswordIdentity creates a new identity with username u and password p
func NewPasswordIdentity(u string, p []byte) (*PasswordIdentity, error) {
	pwdi := &PasswordIdentity{
		username: u,
		password: p,
	}

	return pwdi, nil
}

// GetUsername returns the username used to create pwdi
func (pwdi *PasswordIdentity) GetUsername() string {
	return pwdi.username
}

// GetSecret returns the password used to create pwdi
func (pwdi *PasswordIdentity) GetSecret() []byte {
	return pwdi.password
}
