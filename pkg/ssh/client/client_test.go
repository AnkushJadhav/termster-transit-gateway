package client

import (
	"testing"

	"github.com/AnkushJadhav/termster-transit-gateway/pkg/host"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/identity"
)

func TestNew(t *testing.T) {
	a := "127.0.0.1"
	p := 22
	h := host.New(a, p)

	uname := "lead school"
	pwd := []byte("lead@123")
	pwdi, _ := identity.NewPasswordIdentity(uname, pwd)

	c, err := New(h, pwdi)
	if err != nil {
		t.Errorf(err.Error())
	}
	if c == nil {
		t.Errorf("expected value, but got nil")
	}
}
