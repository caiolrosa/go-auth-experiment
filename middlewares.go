package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// SetupMiddlewares combines all used middlewares and returns a handler
func SetupMiddlewares(handler http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:8080"}),
			handlers.AllowedMethods([]string{
				http.MethodGet,
				http.MethodPost,
				http.MethodDelete,
				http.MethodPut,
			}),
		)(handler))
}
