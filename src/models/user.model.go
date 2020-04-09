package models

// User es el modelo para el usuario
type User struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Rol      `json:"rol,omitempty"`
}
