package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//SkuProductEntityMapping represents sku product entity mapping
type SkuProductEntityMapping struct {
	gorm.Model
	Name        string
	Description string
	Active      int
}

const (
	//SKUPRODUCTENTITYMODEL sku_product_entity_mappings  constant
	SKUPRODUCTENTITYMODEL = "sku_product_entity_mappings"
)

func init() {
	gormDb = db.GormDB()
}

//create sku production entity mapping
func (s *SkuProductEntityMapping) create() (err error) {
	errs := gormDb.Create(s).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", SKUPRODUCTENTITYMODEL, errs[0].Error())
	}
	return
}

//get vertical
func (s *SkuProductEntityMapping) get() (err error) {
	return
}
