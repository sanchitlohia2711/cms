package model

import "time"

//CategoryES category ES doc
type CategoryES struct {
	ESID         uint      `json:"es_id,omitempty"`
	CategoryDBID uint      `json:"category_db_id"`
	Name         uint      `json:"name"`
	Description  string    `json:"description"`
	VerticalID   string    `json:"vertical_id"`
	ParentID     string    `json:"parent_id"`
	ParentIDs    string    `json:"parent_ids"`
	Active       bool      `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdagtedAt   time.Time `json:"updated_at"`
}
