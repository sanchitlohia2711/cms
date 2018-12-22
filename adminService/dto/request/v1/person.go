package v1

//PersonParams create update person
type PersonParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Active      uint8       `json:"active"`
	Tags        interface{} `json:"tags"`
}
