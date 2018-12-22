package v1

import (
	"time"
)

//EventParams for create update event
type EventParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Active      uint8       `json:"active"`
	EntityID    uint        `json:"entityId"`
	StartTime   time.Time   `json:"startTime"`
	EndTime     time.Time   `json:"endTime"`
	StartDate   time.Time   `json:"startDate"`
	Tags        interface{} `json:"tags"`
}
