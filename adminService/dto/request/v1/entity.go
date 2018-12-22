package v1

//EntityParams create update params
type EntityParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CityID      uint        `json:"cityId"`
	CountryID   uint        `json:"countryId"`
	Active      uint8       `json:"active"`
	Lat         float64     `json:"lat"`
	Lon         float64     `json:"long"`
	MerchantID  uint        `json:"merchantId"`
	Type        string      `json:"entityType"`
	Tags        interface{} `json:"tags"`
}
