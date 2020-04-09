package utils

import "errors"

var (
	// ErrUserOk este es un estado satisfactorio
	ErrUserOk = errors.New("Usario ingresado correctamente")

	// ErrUserNotFound error de usuario no encontrado en la db
	ErrUserNotFound = errors.New("Usuario no encontrado")
)
