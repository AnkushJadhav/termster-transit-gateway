package utils

import (
	"testing"
)

func TestGenerateUID(t *testing.T) {
	explen := 36
	uuid := GenerateUID()

	if uuid == "" {
		t.Errorf("expected value, got nil")
	}
	if len(uuid) != explen {
		t.Errorf("expected length of uuid to be %d, got %d", explen, len(uuid))
	}
}