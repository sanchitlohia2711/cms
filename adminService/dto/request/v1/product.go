package v1

import (
	"time"

	"github.com/ko/cms-db/adminService/dto"
)

//ProductParams create update product
type ProductParams struct {
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	VerticalID      uint        `json:"verticalID"`
	BrandID         uint        `json:"brandID"`
	StartDate       time.Time   `json:"startDate"`
	EndDate         time.Time   `json:"EndDate"`
	Attributes      interface{} `json:"attributes"`
	Tags            []string    `json:"tags"`
	ShareURL        string      `json:"share_url"`
	ThumbURL        string      `json:"thumb_url"`
	ImageURL        string      `json:"image_url"`
	HowToRedeem     dto.Meta    `json:"how_to_redeem"`
	TermsConditions dto.Meta    `json:"terms_conditions"`
	ReturnPolicy    dto.Meta    `json:"return_policy"`
	InputFields     string      `json:"input_fields"`
}
