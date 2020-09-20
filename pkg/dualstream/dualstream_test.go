package dualstream

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	d := New()

	if d == nil {
		t.Errorf("expected value for dual stream, but got nil")
	}
	if d.Band1Reader() == nil {
		t.Errorf("expected value for dual stream band1 reader, but got nil")
	}
	if d.Band1Writer() == nil {
		t.Errorf("expected value for dual stream band1 writer, but got nil")
	}
	if d.Band2Reader() == nil {
		t.Errorf("expected value for dual stream band2 reader, but got nil")
	}
	if d.Band2Writer() == nil {
		t.Errorf("expected value for dual stream band2 writer, but got nil")
	}

	teststr1 := "test hello band1!"
	go func() {
		fmt.Fprint(d.Band1Writer(), teststr1)
		d.Band1Writer().Close()
	}()
	b1 := new(bytes.Buffer)
	b1.ReadFrom(d.Band1Reader())
	if teststr1 != b1.String() {
		t.Errorf("inconsistency on band1, expected %s, got %s", teststr1, b1.String())
	}

	teststr2 := "test hello band2!"
	go func() {
		d.Band2Writer().Write([]byte(teststr2))
		d.Band2Writer().Close()
	}()
	b2 := new(bytes.Buffer)
	b2.ReadFrom(d.Band2Reader())
	if teststr2 != b2.String() {
		t.Errorf("inconsistency on band2, expected %s, got %s", teststr2, b2.String())
	}
}
