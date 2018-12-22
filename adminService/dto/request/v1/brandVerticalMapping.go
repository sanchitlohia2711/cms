package v1

//BrandVerticalMapping brand vertical mapping params
type BrandParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      uint8  `json:"active"`
}
