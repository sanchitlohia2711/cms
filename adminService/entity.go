package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//ENTITYMODEL vertical model constant
	ENTITYMODEL = "entities"
)

func init() {
	gormDb = db.GormDB()
}

//Entity represents entity
type Entity struct {
	gorm.Model
	Name       string
	CityID     uint
	CountryID  uint
	Lat        float64
	Long       float64
	MerchantID uint
	EntityType string
	Tags       interface{}
}

//create entity
func (e *Entity) create() (err error) {
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
