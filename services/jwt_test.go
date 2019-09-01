package services

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestJWTExpiration(t *testing.T) {
	expected := time.Now().Add(time.Hour * 24)

	got := JWTExpiration()

	assert.Equal(t, expected.Unix(), got.Unix())
}

func TestClaims(t *testing.T) {
	expectedID := "123"
	expected := &jwt.StandardClaims{
		Id:        expectedID,
		ExpiresAt: JWTExpiration().Unix(),
	}

	got := claims(expectedID)

	assert.Equal(t, expected, got)
}
