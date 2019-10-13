package model

import "time"

//EntityES struct
type EntityES struct {
	ESID           string   `json:"ESID"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Tags           []string `json:"tags"`
	Meta           string   `json:"meta"`
	Attributes     string   `json:"attributes"`
	EntityID       string   `json:"person_id"`
	EntityTypeID   uint     `json:"person_type_id"`
	EntityTypeName string   `json:"person_type_name"`
	Location       Location `json:"location"`
	CityID         uint     `json:"city_id"`
	Address        string   `json:"address"`
	PinCode        string
	Active         uint8          `json:"acitve"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	EmbeddedEntity EmbeddedEntity `json:"embedded_entity"`
}

//EmbeddedEntity represents embedded entity
type EmbeddedEntity struct {
	EntityID   uint
	Lat        float64
	Lon        float64
	EntityType string
}

//Location location struct
type Location struct {
	Lat float64
	Lon float64
}
