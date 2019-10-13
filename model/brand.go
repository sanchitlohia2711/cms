package model

import (
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//BRANDMODEL brand model constant
	BRANDMODEL = "brand"
)

//Brand represents a brand
type Brand struct {
	Base
	Name        string
	Description string
}

//BrandFilters used for filtering brands
type BrandFilters struct {
	Brand
}

//Create brand
func (b *Brand) Create() error {
	var err error
	errs := gormDb.Create(b).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", BRANDMODEL, errs[0].Error())
	}
	return err
}

//GetBrand get the brand
func GetBrand(filters *BrandFilters) ([]*Brand, error) {
	var brands []*Brand
	if filters.Name != "" {
		gormDb = gormDb.Where("name = ?", filters.Name)
	}

	var err error
	errs := gormDb.Find(&brands).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_GET_ERROR", BRANDMODEL, errs[0].Error())
	}
	return brands, err
}
