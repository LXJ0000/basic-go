package handler

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	encrypted, _ := bcrypt.GenerateFromPassword([]byte("Hello@123"), bcrypt.DefaultCost)
	t.Log(string(encrypted))
}
