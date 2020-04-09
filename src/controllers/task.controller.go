package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	mod "github.com/Mynor2397/api-simple/src/models"
	ser "github.com/Mynor2397/api-simple/src/services"
)

// CreateTask  nos permite crear tareas
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task mod.Task
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ser.CreateTask(r.Context(), task)

	if err == errors.New("Agregado satisfactoriamente") {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

}
