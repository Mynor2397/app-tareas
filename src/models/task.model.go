package models

// Task es el model de una tarea simple
type Task struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Property string `json:"property,omitempty"`
}
