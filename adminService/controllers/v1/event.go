package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//Event represents vertical
type Event struct {
	gorm.Model
	Name        string
	Description string
	Active      int
	EntityID    int
	StartTime   time.Time
	EndTime     time.Time
	Tags        interface{}
}

const (
	//EVENTMODEL event model constant
	EVENTMODEL = "events"
)

//create event
func (e *Event) create() (err error) {
	errs := gormDb.Create(e).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", VERTICALMODEL, errs[0].Error())
	}
	return
}

//get event
func (e *Event) get() (err error) {
	return
}
