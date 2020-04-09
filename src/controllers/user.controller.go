package controllers

import (
	"encoding/json"
	"net/http"

	mod "github.com/Mynor2397/api-simple/src/models"
	ser "github.com/Mynor2397/api-simple/src/services"
	uti "github.com/Mynor2397/api-simple/src/utils"
)

type response struct {
	Message string `json:"message,omitempty"`
}

// CreateUser permite ingresar usuarios
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user mod.User
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ser.CreateUser(r.Context(), user)
	if err == uti.ErrUserOk {
		respond(w, response{Message: "Usuario ingresado correctamente"}, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Login es para hacer login
func Login(w http.ResponseWriter, r *http.Request) {
	var user mod.User
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := ser.Login(r.Context(), user)
	if err == uti.ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	respond(w, out, http.StatusOK)
}
