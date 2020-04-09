package models

// Rol es la estructura de roles de nuestra db
type Rol struct {
	IDRol   int64  `json:"id_rol,omitempty"`
	NameRol string `json:"name_rol,omitempty"`
}
