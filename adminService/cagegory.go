package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//Category represents category
type Category struct {
	gorm.Model
	Name        string
	Description string
	Active      uint8
	VerticalID  uint
	ParentID    uint
}

const (
	//CATEGORYMODEL vertical model constant
	CATEGORYMODEL = "categories"
)

func init() {
	gormDb = db.GormDB()
}

//create vertical
func (c *Category) create() (err error) {
	errs := gormDb.Create(c).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", VERTICALMODEL, errs[0].Error())
	}
	return
}

//get vertical
func (c *Category) get() (err error) {
	return
}
