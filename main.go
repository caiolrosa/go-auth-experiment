package main

import (
	"net/http"

	"guardian-api/db/migrations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	migrations.Migrate()

	server := echo.New()
	setupMiddlewares(server)
	setupRoutes(server)

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

func setupRoutes(server *echo.Echo) {
	server.GET("/api/healthcheck", HandleHealthCheck)
	server.POST("/api/login", HandleLogin)
	server.DELETE("/api/login", HandleLogout)
	server.POST("/api/register", HandleRegister)
	server.PUT("/api/user/:id/edit", HandleEditUser)
}
