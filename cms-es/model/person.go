package model

import "time"

//PersonES struct
type PersonES struct {
	ESID           string         `json:"ESID"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Tags           []string       `json:"tags"`
	Meta           string         `json:"meta"`
	Attributes     string         `json:"attributes"`
	PersonID       string         `json:"person_id"`
	PersonTypeID   uint           `json:"person_type_id"`
	PersonTypeName string         `json:"person_type_name"`
	Active         uint8          `json:"acitve"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	EmbeddedEntity EmbeddedEntity `json:"embedded_entity"`
}
