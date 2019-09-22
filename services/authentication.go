package services

import (
	"guardian-api/models"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// IAuthenticationService provides authentication methods
type IAuthenticationService interface {
	AuthenticateUser(email string, password string) (models.User, int, error)
	AuthenticateUserByID(id int64) (models.User, error)
	AuthenticateUserWithCookie(cookie *http.Cookie) (models.User, error)
}

// AuthenticationService implements IAuthenticationService
type AuthenticationService struct {
	*UserRepository
	*JWTService
}

// AuthenticateUser auths a user by email and password
func (a *AuthenticationService) AuthenticateUser(email string, password string) (models.User, int, error) {
	user, err := a.UserRepository.FindByEmail(email)
	if err != nil {
		return user, http.StatusNotFound, err
	}

	if err = user.Authenticate(password); err != nil {
		return user, http.StatusUnauthorized, err
	}
	return user, http.StatusOK, nil
}

// AuthenticateUserByID auths a user by email and password
func (a *AuthenticationService) AuthenticateUserByID(id int64) (models.User, error) {
	return a.UserRepository.FindByID(id)
}

// AuthenticateUserWithCookie auths the user based on the cookie and returns a user struct
func (a *AuthenticationService) AuthenticateUserWithCookie(cookie *http.Cookie) (models.User, error) {
	user := models.User{}

	uid, err := a.JWTService.ParseToken(cookie.Value)
	if err != nil {
		log.Errorf("Unable to parse token: %s", cookie.Value)
		return user, err
	}

	intUID, err := strconv.Atoi(uid)
	if err != nil {
		log.Errorf("Unable to convert %s to integer", uid)
		return user, err
	}

	user, err = a.AuthenticateUserByID(int64(intUID))
	if err != nil {
		log.Error(err)
		return user, err
	}

	return user, nil
}
