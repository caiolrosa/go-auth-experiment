package services

import (
	"guardian-api/models"
	"net/http"
)

// IAuthenticationService provides authentication methods
type IAuthenticationService interface {
	AuthenticateUser(email string, password string) (models.User, int, error)
	AuthenticateUserByID(id int64) (models.User, error)
}

// AuthenticationService implements IAuthenticationService
type AuthenticationService struct {
	*UserRepository
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
