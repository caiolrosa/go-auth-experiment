package main

import (
	"guardian-api/models"
	"guardian-api/services"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

const (
	cookieName       = "hsvr_auth"
	guardianLoginURL = "http://localhost:8080/"
)

// ErrorResponse represents an http error message response
type ErrorResponse struct {
	Message string `json:"message"`
}

// HandleHealthCheck GET /api/healthcheck
func (app *App) HandleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service available")
}

// HandleAuth tries to auth the user with cookie and redirect if not possible
func (app *App) HandleAuth(c echo.Context) error {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		log.Error(err)
		return c.Redirect(
			http.StatusTemporaryRedirect,
			guardianLoginURL,
		)
	}

	uid, err := app.JWTService.ParseToken(cookie.Value)
	if err != nil {
		log.Errorf("Unable to parse token: %s", cookie.Value)
		return c.Redirect(
			http.StatusTemporaryRedirect,
			guardianLoginURL,
		)
	}

	intUID, err := strconv.Atoi(uid)
	if err != nil {
		log.Errorf("Unable to convert %s to integer", uid)
		return c.Redirect(
			http.StatusTemporaryRedirect,
			guardianLoginURL,
		)
	}

	authUser, err := app.AuthenticationService.AuthenticateUserByID(int64(intUID))
	if err != nil {
		log.Error(err)
		return c.Redirect(
			http.StatusTemporaryRedirect,
			guardianLoginURL,
		)
	}

	return c.JSON(http.StatusOK, authUser)
}

// HandleLogin POST /api/login
func (app *App) HandleLogin(c echo.Context) error {
	reqUser := &models.User{}

	if err := c.Bind(reqUser); err != nil {
		return err
	}

	authUser, status, err := app.AuthenticationService.AuthenticateUser(reqUser.Email, reqUser.Password)
	if err != nil {
		return c.JSON(
			status,
			&ErrorResponse{Message: err.Error()},
		)
	}

	token, err := app.JWTService.IssueToken(strconv.FormatInt(authUser.ID, 10))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&ErrorResponse{Message: err.Error()},
		)
	}

	c.SetCookie(&http.Cookie{
		Name:     cookieName,
		Value:    token,
		Expires:  services.JWTExpiration(),
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, authUser)
}

// HandleLogout DELETE /api/login
func (app *App) HandleLogout(c echo.Context) error {
	pastExpire := time.Now().Add(-7 * time.Hour * 24)
	c.SetCookie(&http.Cookie{
		Name:    cookieName,
		Value:   "delete",
		Expires: pastExpire,
	})
	return c.JSON(http.StatusOK, map[string]string{"message": "Logout complete"})
}

// HandleRegister POST /api/register
func (app *App) HandleRegister(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	if !user.Valid() {
		return c.JSON(
			http.StatusUnprocessableEntity,
			&ErrorResponse{Message: "Invalid user"},
		)
	}

	if err := user.EncryptPassword(); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&ErrorResponse{Message: "Unable to create user"},
		)
	}

	newUser, err := app.UserRepository.Save(*user)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&ErrorResponse{Message: err.Error()},
		)
	}

	token, err := app.JWTService.IssueToken(strconv.FormatInt(newUser.ID, 10))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&ErrorResponse{Message: err.Error()},
		)
	}

	c.SetCookie(&http.Cookie{
		Name:     cookieName,
		Value:    token,
		Expires:  services.JWTExpiration(),
		HttpOnly: true,
	})
	return c.JSON(http.StatusOK, newUser)
}

// HandleEditUser PUT /api/user/:id/edit
func (app *App) HandleEditUser(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Endpoint not implemented")
}
