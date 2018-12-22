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
	Name        string `xorm:"'name'"`
	Description string `xorm:"'description'"`
}

//Create brand
func (b *Brand) Create() (err error) {
	errs := gormDb.Create(b).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", BRANDMODEL, errs[0].Error())
	}
	return
}

//get brand
func (b *Brand) get() (err error) {
	return
}
