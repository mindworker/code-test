package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Session stores users session data and expected expiry, when the
// session should be marked for garbage collection
type Session struct {
	Data       Data
	Expiration int64
}

// Data defines the data we store for a session
type Data struct {
	WebsiteURL         string          `json:"websiteUrl,omitempty"`
	SessionID          string          `json:"sessionId,omitempty"`
	ResizeFrom         Dimension       `json:"resizeFrom,omitempty"`
	ResizeTo           Dimension       `json:"resizeTo,omitempty"`
	CopyAndPaste       map[string]bool `json:"copyAndPaste,omitempty"`       // map[fieldId]true
	FormCompletionTime int             `json:"formCompletionTime,omitempty"` // Seconds
}

// Dimension stores window dimensions
//
// @TODO not sure why but test requires Width and Height to be strings
type Dimension struct {
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
}

// MakeSessionID creates a random session ID in the following format.
// XXXXXX-XXXXXX-XXXXXXXXX
func MakeSessionID() (string, error) {
	// Create random 21 bytes
	buf := make([]byte, 21)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", err
	}

	// Use standard Base64 RFC 4648, since it doesn't contain
	// dashes
	b64 := base64.StdEncoding.EncodeToString(buf)

	// Add dashes (-) to form the desired format
	return fmt.Sprintf("%s-%s-%s", b64[0:7], b64[7:13], b64[13:21]), nil
}
