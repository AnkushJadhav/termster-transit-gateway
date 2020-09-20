package session

import (
	"io"

	"github.com/AnkushJadhav/termster-transit-gateway/pkg/ssh/client"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/ssh/terminal"
	"golang.org/x/crypto/ssh"
)

// Session is an ssh session to a remote server
type Session struct {
	s *ssh.Session
}

// StreamConfig is used to configure the read and write streams for the session
type StreamConfig struct {
	in  io.ReadCloser
	out io.WriteCloser
	err io.WriteCloser
}

// New creates a SSH session on the remote server connected to by client c
func New(c *client.Client, sc *StreamConfig) (*Session, error) {
	// create a ssh session from the client
	s, err := c.NewSession()
	if err != nil {
		return nil, err
	}

	sess := &Session{
		s,
	}

	// set the streams for communication
	sess.setStreamConfig(sc)

	return sess, nil
}

// StartShell requests for a PTY on the session and starts a shell that waits until the session is closed
func (s *Session) StartShell(conf terminal.Config) error {
	// request for a PTY on the session
	err := s.s.RequestPty(conf.Name, conf.Height, conf.Width, conf.Modes)
	if err != nil {
		return err
	}

	err = s.s.Shell()
	if err != nil {
		return err
	}

	s.s.Wait()
	return nil
}

// Close closes the session
func (s *Session) Close() error {
	return s.s.Close()
}

// setStreamConfig configures the stdin, stdout and stderr for the session
func (s *Session) setStreamConfig(sc *StreamConfig) {
	s.s.Stdin = sc.in
	s.s.Stdout = sc.out
	s.s.Stderr = sc.err
}
