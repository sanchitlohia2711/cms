package model

import (
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//ENTITYMODEL entity model constant
	ENTITYMODEL = "entity"
)

func init() {
	gormDb = db.GormDB()
}

//Entity represents entity
type Entity struct {
	Base
	Name        string
	Description string
	CityID      uint
	CountryID   uint
	Lat         float64
	Lon         float64
	MerchantID  uint
	Type        string
	Tags        interface{}
}

//Create entity
func (e *Entity) Create() (err error) {
	errs := gormDb.Create(e).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", ENTITYMODEL, errs[0].Error())
	}
	return
}

//get entity
func (e *Entity) get() (err error) {
	return
}
