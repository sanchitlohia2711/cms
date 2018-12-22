package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//SKUMODEL product model constant
	SKUMODEL = "skus"
)

//Sku represents sku
type Sku struct {
	gorm.Model
	Name        string
	Description string
	Tags        interface{}
}

//create sku
func (s *Sku) create() (err error) {
	errs := gormDb.Create(s).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", PRODUCTMODEL, errs[0].Error())
	}
	return
}

//get sku
func (s *Sku) get() (err error) {
	return
}
