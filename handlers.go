package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HandleLogin POST /api/login
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	fmt.Fprintf(w, "Login")
}

// HandleRegister POST /api/register
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
	fmt.Fprintf(w, "Register")
}

// HandleLogout DELETE /api/logout
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout")
}
