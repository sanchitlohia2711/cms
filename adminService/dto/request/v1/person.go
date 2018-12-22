package v1

//PersonParams create update person
type PersonParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	EntityID    uint        `json:"entityId"`
	Active      uint8       `json:"active"`
	Tags        interface{} `json:"tags"`
}
