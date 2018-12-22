package model

import "github.com/sanchitlohia2711/go-extended-error/exerr"

const (
	//BRANDVERTICALMAPPING brand vertical mapping model constant
	BRANDVERTICALMAPPING = "brand_vertical_mapping"
)

//BrandVerticalMapping represents a brand
type BrandVerticalMapping struct {
	Base
	BrandID    uint
	VerticalID uint
}

//Create brand
func (b *BrandVerticalMapping) Create() (err error) {
	errs := gormDb.Create(b).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", BRANDVERTICALMAPPING, errs[0].Error())
	}
	return
}

//get brand
func (b *BrandVerticalMapping) get() (err error) {
	return
}
