package v1

//CategoryFilterParams represents category
type CategoryFilterParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      uint8  `json:"active"`
	CategoryID  uint   `json:"categoryId"`
}
