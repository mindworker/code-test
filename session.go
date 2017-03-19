package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

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
