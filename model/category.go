package model

import (
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//Category represents category
type Category struct {
	Base
	Name        string
	Description string
	VerticalID  uint
	ParentID    uint
}

const (
	//CATEGORYMODEL category model constant
	CATEGORYMODEL = "category"
)

func init() {
	gormDb = db.GormDB()
}

//Create category
func (c *Category) Create() (err error) {
	errs := gormDb.Create(c).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", CATEGORYMODEL, errs[0].Error())
	}
	return
}

//Get category
func (c *Category) Get() (err error) {
	return
}
