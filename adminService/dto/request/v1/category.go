package v1

//CategoryParams brand params
type CategoryParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      uint8  `json:"active"`
	VerticalID  uint   `json:"verticalId"`
	ParentID    uint   `json:"parentId"`
}
