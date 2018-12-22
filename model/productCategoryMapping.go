package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//ProductCategoryMapping represents a product cagegory mapping
type ProductCategoryMapping struct {
	gorm.Model
	ProductID  uint
	CategoryID uint
}

//create vertical
func (p *ProductCategoryMapping) create() (err error) {
	errs := gormDb.Create(p).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", VERTICALMODEL, errs[0].Error())
	}
	return
}

//get vertical
func (p *ProductCategoryMapping) get() (err error) {
	return
}
