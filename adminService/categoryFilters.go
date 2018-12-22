package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//CategoryFilter represents category
type CategoryFilter struct {
	gorm.Model
	Name        string
	Description string
	CategoryID  uint
}

const (
	//CATEGORYFILTERMODEL vertical model constant
	CATEGORYFILTERMODEL = "category_filters"
)

func init() {
	gormDb = db.GormDB()
}

//create vertical
func (c *CategoryFilter) create() (err error) {
	errs := gormDb.Create(c).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", VERTICALMODEL, errs[0].Error())
	}
	return
}

//get category filters
func (c *CategoryFilter) get() (err error) {
	return
}
