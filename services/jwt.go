package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// IJWTService exposes jwt parse, issue and expiration methods
type IJWTService interface {
	IssueToken(id string) (string, error)
	ParseToken(token string) (string, error)
}

// JWTService implements IJWTService
type JWTService struct{}

// IssueToken return a signed JWT token
func (j *JWTService) IssueToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims(id))
	return token.SignedString(signKey())
}

// ParseToken returns an error if not authorized
func (j *JWTService) ParseToken(token string) (string, error) {
	claims := &jwt.StandardClaims{}
	validatedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return signKey(), nil
	})

	if !validatedToken.Valid {
		log.Errorf("Invalid token: %s", validatedToken.Raw)
		return "", jwt.ErrSignatureInvalid
	}

	return claims.Id, err
}

// JWTExpiration provides the hsvr jwt expiration time
func JWTExpiration() time.Time {
	oneDay := time.Hour * 24
	return time.Now().Add(oneDay)
}

func claims(id string) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Id:        id,
		ExpiresAt: JWTExpiration().Unix(),
	}
}

func signKey() []byte {
	return []byte("mustbereplaced")
}
