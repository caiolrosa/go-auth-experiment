package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"guardian-api/user"

	"golang.org/x/crypto/bcrypt"
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
	jsonEncoder := json.NewEncoder(w)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonEncoder.Encode(map[string]string{"error": "Unable to process message body"})
		return
	}

	var user user.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		jsonEncoder.Encode(map[string]string{"error": "Unable to parse message body"})
		return
	}

	if !user.Valid() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		jsonEncoder.Encode(map[string]string{"error": "Invalid user"})
		return
	}

	encrypted, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonEncoder.Encode(map[string]string{"error": "Unable to create user"})
		return
	}

	user.Password = string(encrypted)
	err = user.Save()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonEncoder.Encode(map[string]string{"error": err.Error()})
		return
	}
	jsonEncoder.Encode(user)
}

// HandleLogout DELETE /api/logout
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout")
}
