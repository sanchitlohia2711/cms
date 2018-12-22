package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//CITYMODEL vertical model constant
	CITYMODEL = "cities"
)

//City represents entity
type City struct {
	gorm.Model
	Name        string
	Description string
	CountryID   string
	Tags        interface{}
}

//create city
func (c *City) create() (err error) {
	errs := gormDb.Create(c).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", ENTITYMODEL, errs[0].Error())
	}
	return
}

//get vertical
func (c *City) get() (err error) {
	return
}
