package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//COUNTRYMODEL city model constant
	COUNTRYMODEL = "cities"
)

//Country represents country
type Country struct {
	gorm.Model
	Name        string
	Description string
	Tags        interface{}
}

//create city
func (c *Country) create() (err error) {
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
