package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//PERSONMODEL person model constant
	PERSONMODEL = "persons"
)

func init() {
	gormDb = db.GormDB()
}

//Person represents person
type Person struct {
	gorm.Model
	Name        string
	Description string
	EntityType  string
	Tags        interface{}
}

//create person
func (p *Person) create() (err error) {
	errs := gormDb.Create(p).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", ENTITYMODEL, errs[0].Error())
	}
	return
}

//get person
func (p *Person) get() (err error) {
	return
}
