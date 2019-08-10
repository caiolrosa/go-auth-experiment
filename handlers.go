package main

import (
	"net/http"

	"guardian-api/user"

	"github.com/labstack/echo/v4"
)

// HandleHealthCheck GET /api/healthcheck
func (app *App) HandleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service available")
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
			map[string]string{"error": "Incorrect password"},
		)
	}

	return c.JSON(http.StatusOK, authUser)
}

// HandleLogout DELETE /api/login
func (app *App) HandleLogout(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Endpoint not implemented")
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

	return c.JSON(http.StatusOK, user)
}

// HandleEditUser PUT /api/user/:id/edit
func (app *App) HandleEditUser(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Endpoint not implemented")
}
