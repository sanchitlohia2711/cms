package model

import (
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//PersonEntityMappingModel brand model constant
	PersonEntityMappingModel = "person_entity_mapping"
)

//PersonEntityMapping represents a brand
type PersonEntityMapping struct {
	Base
	PersonID uint
	EntityID uint
}

//Create brand
func (p *PersonEntityMapping) Create() (err error) {
	errs := gormDb.Create(p).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", PersonEntityMappingModel, errs[0].Error())
	}
	return
}

//Get brand
func (p *PersonEntityMapping) Get() (err error) {
	return
}
