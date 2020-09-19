package identity

import (
	"bytes"
	"testing"
)

func TestNewPasswordIdentity(t *testing.T) {
	uname := "ec2-user"
	pwd := []byte("test")

	pwdi, err := NewPasswordIdentity(uname, pwd)

	if err != nil {
		t.Errorf(err.Error())
	}
	if pwdi == nil {
		t.Errorf("expected value, got nil")
	}
	if pwdi.GetUsername() != uname {
		t.Errorf("expected %s, got %s", uname, pwdi.GetUsername())
	}
	if bytes.Compare(pwdi.GetSecret(), pwd) != 0 {
		t.Errorf("expected %v, got %v", pwd, pwdi.GetSecret())
	}
}

func TestNewPrivateKeyIdentity(t *testing.T) {
	uname := "ec2-user"
	key := []byte("test")

	pki, err := NewPrivateKeyIdentity(uname, key)

	if err != nil {
		t.Errorf(err.Error())
	}
	if pki == nil {
		t.Errorf("expected value, got nil")
	}
	if pki.GetUsername() != uname {
		t.Errorf("expected %s, got %s", uname, pki.GetUsername())
	}
	if bytes.Compare(pki.GetSecret(), key) != 0 {
		t.Errorf("expected %v, got %v", key, pki.GetSecret())
	}
}
