package model

import (
	"time"

	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//PRODUCTSTATUS : product status enum
type PRODUCTSTATUS int

//PRODUCTVISIBILITY : product visibility enum
type PRODUCTVISIBILITY int

const (
	//ProductDisalbed status
	ProductDisalbed PRODUCTSTATUS = 0
	//ProductEnabled status
	ProductEnabled PRODUCTSTATUS = 1

	//ProductInvisible status
	ProductInvisible PRODUCTVISIBILITY = 0
	//ProductVisible status
	ProductVisible PRODUCTVISIBILITY = 1

	//PRODUCTMODEL product model constant
	PRODUCTMODEL = "product"
)

func init() {
	gormDb = db.GormDB()
}

//Product represents product
type Product struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	VerticalID  uint
	BrandID     uint
	StartDate   time.Time
	Status      PRODUCTSTATUS
	Visibility  PRODUCTVISIBILITY
	Tags        interface{}
}

//Create product
func (p *Product) Create() (err error) {
	errs := gormDb.Create(p).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", PRODUCTMODEL, errs[0])
	}
	return
}

//Get product
func (p *Product) get() (err error) {
	return
}
