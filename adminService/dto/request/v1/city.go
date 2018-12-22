package v1

//CityParams city create update params
type CityParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Active      uint8       `json:"active"`
	CountryID   string      `json:"countryId"`
	Tags        interface{} `json:"tags"`
}
