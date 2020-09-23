package sessionmanager

import (
	"fmt"
	"io"
	"sync"

	"github.com/AnkushJadhav/termster-transit-gateway/pkg/host"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/identity"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/ssh/client"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/ssh/session"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/ssh/terminal"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/utils"
)

type payload struct {
	c *client.Client
	s *session.Session
}

// SessionManager manages all active sessions and creates new ones
type SessionManager struct {
	sync.Mutex
	store map[string]*payload
}

// New creates a new SessionManager
func New() *SessionManager {
	store := make(map[string]*payload)
	sm := &SessionManager{
		store: store,
	}

	return sm
}

// CreateSession creates a new session and adds it to the session manager's store
func (sm *SessionManager) CreateSession(h *host.Host, i identity.Identity, instream io.ReadCloser, outstream, errstream io.WriteCloser) (string, error) {
	// create a new SSH client using host and identity
	c, err := client.New(h, i)
	if err != nil {
		return "", err
	}

	// create a new streamconfig for the session
	sc := &session.StreamConfig{
		In:  instream,
		Out: outstream,
		Err: errstream,
	}
	// create a new SSH session on the client
	sess, err := session.New(c, sc)
	if err != nil {
		return "", err
	}
	termconf := terminal.GetTerminalConfig()
	err = sess.StartShell(termconf)
	if err != nil {
		return "", err
	}

	// retry in case duplicate id is generated
	uid := ""
	retry := 2
	for {
		uid = utils.GenerateUID()
		err = sm.addSessionToStore(uid, c, sess)
		if err == nil {
			break
		} else {
			if retry == 0 {
				return "", fmt.Errorf("failed to add session to store, exhausted addition attempts - %d", (retry + 1))
			}
		}
		retry--
	}

	return uid, nil
}

// TerminateSession closes a session and removes it from the store
func (sm *SessionManager) TerminateSession(uid string) error {
	// get the session from the store
	sess := sm.store[uid]
	if sess == nil {
		return fmt.Errorf("session with uid %s does not exist", uid)
	}

	// close the session
	sess.s.Close()
	// close the client
	sess.c.Close()

	// remove from the store
	sm.removeSessionFromStore(uid)
	return nil
}

func (sm *SessionManager) addSessionToStore(id string, c *client.Client, sess *session.Session) error {
	sm.Lock()
	defer sm.Unlock()

	if _, present := sm.store[id]; present {
		return fmt.Errorf("store already has session with id %s", id)
	}

	sm.store[id] = &payload{
		c: c,
		s: sess,
	}
	return nil
}

func (sm *SessionManager) removeSessionFromStore(id string) {
	delete(sm.store, id)
}
