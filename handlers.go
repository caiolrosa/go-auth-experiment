package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"guardian-api/user"
)

// HandleLogin POST /api/login
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	jsonEncoder := json.NewEncoder(w)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var reqUser user.User
	if err = json.Unmarshal(data, &reqUser); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		jsonEncoder.Encode(map[string]string{"error": "Unable to parse message body"})
		return
	}

	authUser, err := reqUser.FindByEmail()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonEncoder.Encode(map[string]string{"error": "User not found"})
		return
	}

	if err = authUser.Authenticate(reqUser.Password); err != nil {
		w.WriteHeader(http.StatusForbidden)
		jsonEncoder.Encode(map[string]string{"error": "Incorrect password"})
		return
	}

	jsonEncoder.Encode(authUser)
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
	if err = json.Unmarshal(data, &user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		jsonEncoder.Encode(map[string]string{"error": "Unable to parse message body"})
		return
	}

	if !user.Valid() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		jsonEncoder.Encode(map[string]string{"error": "Invalid user"})
		return
	}

	if err = user.EncryptPassword(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonEncoder.Encode(map[string]string{"error": "Unable to create user"})
		return
	}

	if err = user.Save(); err != nil {
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
