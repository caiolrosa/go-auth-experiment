package main

import (
	"net/http"
	"strconv"
	"time"

	"guardian-api/user"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

const (
	cookieName       = "hsvr_auth"
	guardianLoginURL = "http://localhost:8080/"
)

// HandleHealthCheck GET /api/healthcheck
func (app *App) HandleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service available")
}

// HandleAuth tries to auth the user with cookie and redirect if not possible
func (app *App) HandleAuth(c echo.Context) error {
	reqUser := &user.User{
		DBClient: app.dbClient,
	}

	cookie, err := c.Cookie(cookieName)
	if err == nil {
		uid, err := authToken(cookie.Value)
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

		reqUser.ID = int64(intUID)
		authUser, err := reqUser.FindByID()
		if err != nil {
			log.Error(err)
			return c.Redirect(
				http.StatusTemporaryRedirect,
				guardianLoginURL,
			)
		}

		return c.JSON(http.StatusOK, authUser)
	}

	return c.Redirect(
		http.StatusTemporaryRedirect,
		guardianLoginURL,
	)
}

// HandleLogin POST /api/login
func (app *App) HandleLogin(c echo.Context) error {
	reqUser := &user.User{
		DBClient: app.dbClient,
	}

	if err := c.Bind(reqUser); err != nil {
		return err
	}

	authUser, err := reqUser.FindByEmail()
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{"message": "User not found"},
		)
	}

	if err = authUser.Authenticate(reqUser.Password); err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{"message": "Incorrect password"},
		)
	}

	token, err := IssueToken(strconv.FormatInt(authUser.ID, 10))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": err.Error()},
		)
	}

	c.SetCookie(&http.Cookie{
		Name:     cookieName,
		Value:    token,
		Expires:  JWTExpiration(),
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
	user := &user.User{
		DBClient: app.dbClient,
	}
	if err := c.Bind(user); err != nil {
		return err
	}

	if !user.Valid() {
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{"message": "Invalid user"},
		)
	}

	if err := user.EncryptPassword(); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Unable to create user"},
		)
	}

	if err := user.Save(); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": err.Error()},
		)
	}

	token, err := IssueToken(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"message": err.Error()},
		)
	}

	c.SetCookie(&http.Cookie{
		Name:     cookieName,
		Value:    token,
		Expires:  JWTExpiration(),
		HttpOnly: true,
	})
	return c.JSON(http.StatusOK, user)
}

// HandleEditUser PUT /api/user/:id/edit
func (app *App) HandleEditUser(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Endpoint not implemented")
}

func authToken(token string) (string, error) {
	id, err := ParseToken(token)
	if err != nil {
		return "", err
	}

	return id, nil
}
