package model

import (
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
