package model

import (
	"time"

	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//Event represents vertical
type Event struct {
	Base
	Name        string
	Description string
	EntityID    uint
	StartTime   time.Time
	EndTime     time.Time
	StartDate   time.Time
	Tags        interface{}
}

const (
	//EVENTMODEL event model constant
	EVENTMODEL = "event"
)

//Create event
func (e *Event) Create() (err error) {
	errs := gormDb.Create(e).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", EVENTMODEL, errs[0].Error())
	}
	return
}

//Get event
func (e *Event) get() (err error) {
	return
}
