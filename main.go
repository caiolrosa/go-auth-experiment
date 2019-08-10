package main

import (
	"net/http"

	"guardian-api/db"
	"guardian-api/db/migrations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// App holds all necessary application information and interfaces
type App struct {
	dbClient db.API
}

func main() {
	app := &App{
		dbClient: &db.Client{},
	}
	migrations.Migrate(app.dbClient)

	server := echo.New()
	setupMiddlewares(server)
	setupRoutes(app, server)

	server.Logger.Fatal(server.Start(":3000"))
}

func setupMiddlewares(server *echo.Echo) {
	server.Use(
		middleware.Logger(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:8080"},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodDelete,
				http.MethodPut,
			},
		}),
		middleware.Recover(),
	)
}

func setupRoutes(app *App, server *echo.Echo) {
	server.GET("/api/healthcheck", app.HandleHealthCheck)
	server.POST("/api/login", app.HandleLogin)
	server.DELETE("/api/login", app.HandleLogout)
	server.POST("/api/register", app.HandleRegister)
	server.PUT("/api/user/:id/edit", app.HandleEditUser)
}
