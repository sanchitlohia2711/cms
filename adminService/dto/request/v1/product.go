package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//PRODUCTMODEL product model constant
	PRODUCTMODEL = "products"
)

func init() {
	gormDb = db.GormDB()
}

//Product represents product
type Product struct {
	gorm.Model
	Name        string
	Description string
	Tags        interface{}
}

//create product
func (p *Product) create() (err error) {
	errs := gormDb.Create(p).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", PRODUCTMODEL, errs[0].Error())
	}
	return
}

//get product
func (p *Product) get() (err error) {
	return
}
