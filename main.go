package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"guardian-api/db/migrations"
)

func main() {
	migrations.Migrate()
	router := mux.NewRouter()
	router.HandleFunc("/api/login", HandleLogin).Methods(http.MethodPost)
	http.ListenAndServe(":3000", SetupMiddlewares(router))
}
