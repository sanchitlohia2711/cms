package v1

//SkuParams create update sku
type SkuParams struct {
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	ProductID        uint        `json:"productId"`
	Price            float64     `json:"price"`
	OfferPrice       float64     `json:"offerPrice"`
	EventID          uint        `json:"eventId"`
	PayTypeSupported interface{} `json:"payTypeSupported"`
	LongRichDesc     string      `json:"longRichDesc"`
	Active           uint8       `json:"active"`
	Tags             interface{} `json:"tags"`
}
