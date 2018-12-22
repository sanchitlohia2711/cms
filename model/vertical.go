package model

import (
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//Vertical represents vertical
type Vertical struct {
	Base
	Name        string
	Description string
}

const (
	//VERTICALMODEL vertical model constant
	VERTICALMODEL = "vertical"
)

func init() {
	gormDb = db.GormDB()
}

//create vertical
func (v *Vertical) Create() (err error) {
	errs := gormDb.Create(v).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", VERTICALMODEL, errs[0].Error())
	}
	return
}

//get vertical
func (v *Vertical) get() (err error) {
	return
}
