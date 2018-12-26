package model

import "time"

//ProductES struct
type ProductES struct {
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Tags            []string    `json:"tags"`
	VerticalID      uint        `json:"vertical_id"`
	BrandID         uint        `json:"brand_id"`
	Status          uint8       `json:"status"`
	Visibility      uint8       `json:"visibility"`
	StartDate       time.Time   `json:"start_date"`
	EndDate         time.Time   `json:"end_date"`
	ShareURL        string      `json:"share_url"`
	ThumbURL        string      `json:"thumb_url"`
	ImageURL        string      `json:"image_url"`
	InputFields     interface{} `json:"input_fields"`
	HowToRedeem     string      `json:"how_to_redeem"`
	ReturnPolicy    string      `json:"return_policy"`
	TermsConditions string      `json:"terms_conditions"`
	Attributes      interface{} `json:"attributes"`
	Category        [][]uint    `json:"category"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}
