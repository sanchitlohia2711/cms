package model

import (
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//CITYMODEL vertical model constant
	CITYMODEL = "city"
)

//City represents entity
type City struct {
	Base
	Name        string
	Description string
	CountryID   uint
}

//Create city
func (c *City) Create() (err error) {
	errs := gormDb.Create(c).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", CITYMODEL, errs[0].Error())
	}
	return
}

//get vertical
func (c *City) get() (err error) {
	return
}
