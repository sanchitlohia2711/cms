package model

import "time"

//ProductES struct
type ProductES struct {
	Name         string
	Description  string
	Tags         []string
	VerticalID   uint
	BrandID      uint
	Status       uint8
	Visibility   uint8
	StartDate    time.Time
	EndDate      time.Time
	ShareURL     string
	ThumbURL     string
	ImageURL     string
	InputFields  string
	HowToRedeem  string
	ReturnPolicy string
	TNC          string
	Attributes   interface{}
	Category     [][]uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Location struct {
	EntityID uint
	Lat      float64
	Lon      float64
}

type Stores struct {
	ID    uint
	Name  uint
	Image uint
}
