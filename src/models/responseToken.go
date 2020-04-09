package models

//ResponseToken es el modelo de respuesta
type ResponseToken struct {
	Token string `json:"token,omitempty"`
	User  `json:"user,omitempty"`
}
