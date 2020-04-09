package services

import (
	"context"
	"errors"
	"strings"

	mod "github.com/Mynor2397/api-simple/src/models"
	"github.com/Mynor2397/api-simple/src/mysql"
)

var db = mysql.Connect()

// CreateTask es una función servicio para crear tareas
func CreateTask(ctx context.Context, task mod.Task) error {
	task.Name = strings.TrimSpace(task.Name)
	task.Property = strings.TrimSpace(task.Property)

	query := "INSERT INTO task VALUES (name, property) VALUES (?, ?)"
	_, err := db.QueryContext(ctx, query, task.Name, task.Property)
	if err != nil {
		return err
	}

	return errors.New("Agregado satisfactoriamente")
}
