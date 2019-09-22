package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// UIDContextKey returns the echo uid context key
func (app *App) UIDContextKey() string {
	return "id"
}

// AuthMiddleware sets up the JWT authentication middleware
func (app *App) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(cookieName)
		if err != nil {
			log.Error(err)
			return c.Redirect(
				http.StatusTemporaryRedirect,
				guardianLoginURL,
			)
		}

		authUser, err := app.AuthenticationService.AuthenticateUserWithCookie(cookie)
		if err != nil {
			log.Error(err)
			return c.Redirect(
				http.StatusTemporaryRedirect,
				guardianLoginURL,
			)
		}

		c.Set(app.UIDContextKey(), authUser.ID)
		return next(c)
	}
}
