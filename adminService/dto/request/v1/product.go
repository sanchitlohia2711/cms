package v1

//ProductParams create update product
type ProductParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Active      uint8       `json:"active"`
	Tags        interface{} `json:"tags"`
}
