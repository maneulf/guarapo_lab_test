package handlers

import (
	"crypto/rand"
	"encoding/hex"
)

type Base struct {
	Usernames map[string]string
}

func (b *Base) generateToken(length int) string {
	token := make([]byte, length)
	if _, err := rand.Read(token); err != nil {
		return ""
	}
	return hex.EncodeToString(token)
}
