package v1

//BrandParams brand params
type BrandParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      uint8  `json:"active"`
}
