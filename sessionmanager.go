package main

import (
	"errors"
	"log"
	"runtime"
	"sync"
	"time"
)

// SessionManager keeps track of all sessions from creation, updating
// to destroying. Since we should not hold sessions for ever in
// memory, a cleaner is implemented that concurrently sweeps for
// sessions that have passed a specified expiry. The removal of old
// sessions are needed, otherwise the backend will eventually run out
// of memory
type SessionManager struct {
	sessions  map[string]Session
	expiresIn time.Duration
	mu        sync.Mutex
	stop      chan bool
}

// NewSessionManager creates new SessionManager which keeps sessions a
// minimum of specified sessionExpiresIn duration. Note that the
// longest time a session stays in memory is the sessionExpiresIn
// duration plus the interval between clean calls. See Cleaner() for
// more
func NewSessionManager(sessionExpiresIn time.Duration) *SessionManager {
	m := &SessionManager{
		sessions:  make(map[string]Session),
		expiresIn: sessionExpiresIn,
	}

	go m.Cleaner()
	runtime.SetFinalizer(m, stopCleaner)

	return m
}

// CreateSession creates a new session and returns the sessionID
func (m *SessionManager) CreateSession() (string, error) {
	sessionID, err := MakeSessionID()
	if err != nil {
		return "", err
	}

	m.sessions[sessionID] = Session{
		Data:       Data{},
		Expiration: time.Now().Add(m.expiresIn).Unix(),
	}

	return sessionID, nil
}

// GetSessionData returns data related to session if sessionID is
// found, errors otherwise
func (m *SessionManager) GetSessionData(sessionID string) (*Data, error) {
	data, ok := m.sessions[sessionID]
	if !ok {
		return nil, errors.New("SessionID does not exists")
	}

	return &data.Data, nil
}

// UpdateSession updates the session with new sessionData and renews
// expiration time as well
func (m *SessionManager) UpdateSession(sessionID string, sessionData Data) error {
	// Check if session actually exist
	//
	// Note that we do not need to use a mutex lock here, as we
	// are simply reading
	_, ok := m.sessions[sessionID]
	if !ok {
		return errors.New("SessionID does not exists")
	}

	// Update session
	m.mu.Lock()
	m.sessions[sessionID] = Session{
		// Renew expiration
		Expiration: time.Now().Add(m.expiresIn).Unix(),
		Data:       sessionData,
	}
	m.mu.Unlock()

	return nil
}

//////////////////////////////////////////////////////////////////////
/// Code related to cleaning expired sessions

// Cleaner cleans expired sessions every 5 seconds
func (m *SessionManager) Cleaner() {
	m.stop = make(chan bool)

	tick := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-tick.C:
			m.Clean()
		case <-m.stop:
			tick.Stop()
			log.Println("Cleaner stopped")
			return
		}
	}
}

// stopCleaner is called by runtime finalizer to stop the cleaner as
// soon as the SessionManager gets destroyed
func stopCleaner(m *SessionManager) {
	m.stop <- true
}

// Clean removes expired sessions
func (m *SessionManager) Clean() {
	now := time.Now().Unix()

	m.mu.Lock()
	for k, v := range m.sessions {
		if now > v.Expiration {
			delete(m.sessions, k)
		}
	}
	m.mu.Unlock()

	log.Println(m.sessions)
}
