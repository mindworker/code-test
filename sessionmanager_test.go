package main

import (
	"reflect"
	"testing"
	"time"
)

func TestSessionManagersCreationAndUpdate(t *testing.T) {
	// Create manager and new session
	m := NewSessionManager(1 * time.Second)
	sID, err := m.CreateSession()
	if err != nil {
		t.Error("Error CreateSession:", err)
	}

	data, err := m.GetSessionData(sID)
	if err != nil {
		t.Error("Error GetSessionData:", err)
	}

	// refData is used as reference to check data integrety
	refData := &Data{
		SessionID:    sID,
		CopyAndPaste: make(map[string]bool),
	}

	// Check if saved data is as expected
	if !reflect.DeepEqual(data, refData) {
		t.Error("SessionData should only have SessionID filled")
	}

	// This is only to ensure a different expiration time
	time.Sleep(1 * time.Second)

	// Modify and update data
	data.WebsiteURL = "longhoang.de"
	refData.WebsiteURL = "longhoang.de"
	err = m.UpdateSessionData(sID, data)
	if err != nil {
		t.Error("Error UpdateSessionData:", err)
	}

	// Retrieve data from manager again
	data, err = m.GetSessionData(sID)
	if err != nil {
		t.Error("Error GetSessionData:", err)
	}

	// Check if the rest of the fields are still the same
	if !reflect.DeepEqual(data, refData) {
		t.Error("SessionData should only have SessionID and WebsiteURL filled")
	}
}

func TestSessionManagersCleaner(t *testing.T) {
	m := NewSessionManager(1 * time.Second)
	sID, err := m.CreateSession()
	if err != nil {
		t.Error("Error CreateSession:", err)
	}

	_, err = m.GetSessionData(sID)
	if err != nil {
		t.Error("Should be able to get session data for", sID, ", but:", err)
	}

	// Note that the cleaner is only running every 5s
	time.Sleep(7 * time.Second)
	_, err = m.GetSessionData(sID)
	if err != ErrSessionNotFound {
		t.Error("Should return ErrSessionNotFound instead of", err)
	}
}
