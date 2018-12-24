package model

import (
	"fmt"
	"time"

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

//SkuProductFilterParams represents sku product entity mapping
type SkuProductEntityMappingFilters struct {
	ProductID uint
	EntityID  uint
	SkuID     uint
	CreatedAt time.Time
	UpdatedAt time.Time
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

//GetSkuProductEntityMapping sku product entity mapping
func GetSkuProductEntityMapping(s *SkuProductEntityMappingFilters) (skuProductEntityMappings []*SkuProductEntityMapping, err error) {
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
