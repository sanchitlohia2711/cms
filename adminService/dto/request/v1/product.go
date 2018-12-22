package v1

import "time"

//ProductParams create update product
type ProductParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	VerticalID  uint        `json:"verticalID"`
	BrandID     uint        `json:"brandID"`
	StartTime   time.Time   `json:"startTime"`
	Tags        interface{} `json:"tags"`
}
