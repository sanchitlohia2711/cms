package v1

//CountryParams create update params
type CountryParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Active      uint8       `json:"active"`
	Tags        interface{} `json:"tags"`
}
