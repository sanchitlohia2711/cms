package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//ProductCategoryMappingModel mapping of product and category
	ProductCategoryMappingModel = "product_category_mapping"
)

//ProductCategoryMapping represents a product cagegory mapping
type ProductCategoryMapping struct {
	Base
	ProductID  uint
	CategoryID uint
}

//ProductCategoryMappingFilters represents a product cagegory mapping
type ProductCategoryMappingFilters struct {
	Active     uint8
	ProductID  uint
	CategoryID uint
}

//Create product category mapping
func (p *ProductCategoryMapping) Create() (err error) {
	errs := gormDb.Create(p).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", ProductCategoryMappingModel, errs[0].Error())
	}
	return
}

//Get product category mapping
func (p *ProductCategoryMapping) get() (err error) {
	return
}

//GetProductCateogryMapping sku product entity mapping
func GetProductCateogryMapping(filters *ProductCategoryMappingFilters) (productCategoryMappings []*ProductCategoryMapping, err error) {
	var db *gorm.DB
	if filters.ProductID > 0 {
		db = gormDb.Where("product_id = ?", filters.ProductID)
	}

	db = db.Where("active = ?", true)

	if filters.CategoryID > 0 {
		db = gormDb.Where("category_id = ?", filters.CategoryID)
	}
	errs := db.Find(&productCategoryMappings).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs[0])
		err = exerr.NewExtendedError("SQL_GET_ERROR", SKUPRODUCTENTITYMODEL, errs[0].Error())
	}
	return
}
