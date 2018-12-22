package model

import (
	"time"

	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//COUNTRYMODEL city model constant
	COUNTRYMODEL = "cities"
)

type t struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    uint8
}

//Country represents country
type Country struct {
	Base
	Name        string
	Description string
}

//Create city
func (c *Country) Create() (err error) {
	errs := gormDb.Create(c).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", ENTITYMODEL, errs[0].Error())
	}
	return
}

//get vertical
func (c *Country) get() (err error) {
	return
}
