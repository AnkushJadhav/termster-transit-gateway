package dualstream

import (
	"io"
)

type band struct {
	r io.ReadCloser
	w io.WriteCloser
}

// DualStream is a two-way communication mechanism that provides two bands to support the two-way communication.
// Band1 reader streams to the Band1 writer and the same applies for Band2
type DualStream struct {
	band1 *band
	band2 *band
}

// New creates a new DualStream
func New() *DualStream {
	b1r, b1w := io.Pipe()
	b1 := &band{
		r: b1r,
		w: b1w,
	}

	b2r, b2w := io.Pipe()
	b2 := &band{
		r: b2r,
		w: b2w,
	}

	return &DualStream{
		band1: b1,
		band2: b2,
	}
}

// Band1Reader returns the reader for band1 of the DualStream d
func (d *DualStream) Band1Reader() io.ReadCloser {
	return d.band1.r
}

// Band1Writer returns the writer for band1 of the DualStream d
func (d *DualStream) Band1Writer() io.WriteCloser {
	return d.band1.w
}

// Band2Reader returns the reader for band1 of the DualStream d
func (d *DualStream) Band2Reader() io.ReadCloser {
	return d.band2.r
}

// Band2Writer returns the writer for band1 of the DualStream d
func (d *DualStream) Band2Writer() io.WriteCloser {
	return d.band2.w
}
