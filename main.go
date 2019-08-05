package main

import (
	"net/http"

	"guardian-api/db/migrations"

	"github.com/gorilla/mux"
)

func main() {
	migrations.Migrate()
	router := mux.NewRouter()
	router.HandleFunc("/api/login", HandleLogin).Methods(http.MethodPost)
	router.HandleFunc("/api/register", HandleRegister).Methods(http.MethodPost)
	http.ListenAndServe(":3000", SetupMiddlewares(router))
}
