package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
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
	ID              uint `gorm:"primary_key"`
	Name            string
	Price           uint
	Description     string
	Tags            string
	VerticalID      uint
	BrandID         uint
	Status          PRODUCTSTATUS
	Visibility      PRODUCTVISIBILITY
	StartDate       time.Time
	EndDate         time.Time
	ShareURL        string
	ThumbURL        string
	ImageURL        string
	InputFields     interface{}
	HowToRedeem     string
	ReturnPolicy    string
	TermsConditions string
	Attributes      interface{}
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//ProductFilterParams product filter params
type ProductFilterParams struct {
	ID uint
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

//GetProduct get all product
func GetProduct(filters *ProductFilterParams) (products []*Product, err error) {
	var db *gorm.DB
	if filters.ID > 0 {
		db = gormDb.Where("id = ?", filters.ID)
	}
	errs := db.Find(&products).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs[0])
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", SKUPRODUCTENTITYMODEL, errs[0].Error())
	}
	return
}
