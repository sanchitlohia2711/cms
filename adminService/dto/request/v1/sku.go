package v1

//SkuParams create update sku
type SkuParams struct {
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	ProductID        uint        `json:"productId"`
	Price            float64     `json:"price"`
	OfferPrice       float64     `json:"offerPrice"`
	EventID          uint        `json:"eventId"`
	PayTypeSupported string      `json:"payTypeSupported"`
	LongRichDesc     string      `json:"longRichDesc"`
	Tags             interface{} `json:"tags"`
	MerchantID       string      `json:"merchantId"`
	VerticalID       uint        `json:"verticalId"`
}
