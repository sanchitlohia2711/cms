package model

import (
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//BRANDMODEL brand model constant
	BRANDMODEL = "brand"
)

//Brand represents a brand
type Brand struct {
	Base
	Name        string
	Description string
}

//Create brand
func (b *Brand) Create() (err error) {
	errs := gormDb.Create(b).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", BRANDMODEL, errs[0].Error())
	}
	return
}

//Get brand
func (b *Brand) Get() (err error) {
	return
}
