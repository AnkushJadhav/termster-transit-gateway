package sessionmanager

import (
	"fmt"
	"io"
	"sync"

	"github.com/AnkushJadhav/termster-transit-gateway/pkg/utils"

	"github.com/AnkushJadhav/termster-transit-gateway/pkg/host"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/identity"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/ssh/client"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/ssh/session"
	"github.com/AnkushJadhav/termster-transit-gateway/pkg/ssh/terminal"
)

// SessionManager manages all active sessions and creates new ones
type SessionManager struct {
	sync.Mutex
	store map[string]*session.Session
}

// New creates a new SessionManager
func New() *SessionManager {
	store := make(map[string]*session.Session)
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
	for {
		uid = utils.GenerateUID()
		err = sm.addSessionToStore(uid, sess)
		if err == nil {
			break
		}
	}

	return uid, nil
}

func (sm *SessionManager) addSessionToStore(id string, sess *session.Session) error {
	sm.Lock()
	defer sm.Unlock()

	if _, present := sm.store[id]; present {
		return fmt.Errorf("store already has session with id %s", id)
	}

	sm.store[id] = sess
	return nil
}
