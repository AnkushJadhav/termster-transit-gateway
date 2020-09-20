package host

import (
	"testing"
)

func TestNew(t *testing.T) {
	a := "127.0.0.1"
	p := 22
	s := "127.0.0.1:22"

	h := New(a, p)

	if h == nil {
		t.Errorf("expected value, got nil")
	}
	if h.String() != s {
		t.Errorf("expected %s, got %s", s, h.String())
	}
}
