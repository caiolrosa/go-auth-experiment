package main

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestJWTExpiration(t *testing.T) {
	expected := time.Now().Add(time.Hour * 24)
	got := JWTExpiration()

	if expected.Unix() != got.Unix() {
		t.Errorf("\nExpected: %v\nGot: %v", expected.Unix(), got.Unix())
	}
}

func TestClaims(t *testing.T) {
	expectedID := "123"
	expected := &jwt.StandardClaims{
		Id:        expectedID,
		ExpiresAt: JWTExpiration().Unix(),
	}
	got := claims(expectedID)

	if expected.Id != got.Id || expected.ExpiresAt != got.ExpiresAt {
		t.Errorf("\nExpected: %v\nGot: %v", expected, got)
	}
}
