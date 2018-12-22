package model

import (
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//PERSONMODEL person model constant
	PERSONMODEL = "person"
)

func init() {
	gormDb = db.GormDB()
}

//Person represents person
type Person struct {
	Base
	Name        string
	Description string
	Tags        interface{}
}

//Create person
func (p *Person) Create() (err error) {
	errs := gormDb.Create(p).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", PERSONMODEL, errs[0].Error())
	}
	return
}

//Get person
func (p *Person) Get() (err error) {
	return
}
