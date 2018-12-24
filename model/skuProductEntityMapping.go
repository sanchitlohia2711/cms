package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//SkuProductEntityMapping represents sku product entity mapping
type SkuProductEntityMapping struct {
	Base
	ProductID uint
	EntityID  uint
	SkuID     uint
}

const (
	//SKUPRODUCTENTITYMODEL sku_product_entity_mappings  constant
	SKUPRODUCTENTITYMODEL = "sku_product_entity_mappings"
)

func init() {
	gormDb = db.GormDB()
}

//Create sku product entity mapping
func (s *SkuProductEntityMapping) Create() (err error) {
	errs := gormDb.Create(s).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", SKUPRODUCTENTITYMODEL, errs[0].Error())
	}
	return
}

//Get sku product entity mapping
func (s *SkuProductEntityMapping) Get() (skuProductEntityMappings []*SkuProductEntityMapping, err error) {
	var db *gorm.DB
	if s.ProductID > 0 {
		db = gormDb.Where("product_id = ?", s.ProductID)
	}

	if s.EntityID > 0 {
		db = db.Where("entity_id = ?", s.EntityID)
	}

	if s.SkuID > 0 {
		db = db.Where("sku_id = ?", s.SkuID)
	}
	errs := db.Find(&skuProductEntityMappings).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs[0])
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", SKUPRODUCTENTITYMODEL, errs[0].Error())
	}
	return
}
