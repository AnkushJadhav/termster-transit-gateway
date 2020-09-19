package identity

// An Identity represents a credential set used to authenticate with a remote host
type Identity interface {
	GetUsername() string
	GetSecret() []byte
}
