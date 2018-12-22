package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Event servcie struct
type Event struct {
	*model.Event
}

//CreateEvent create the entity
func CreateEvent(params *requestDTOV1.EventParams) (event *Event, err error) {
	e := &model.Event{}
	e.Name = params.Name
	e.Description = params.Description
	e.EntityID = params.EntityID
	e.StartTime = params.StartTime
	e.EndTime = params.EndTime
	e.StartDate = params.StartDate
	e.Tags = params.Tags
	err = e.Create()
	if err != nil {
		return
	}
	event = &Event{e}
	return
}
