package model

import (
	"time"

	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//SkuStatus : product status enum
type SkuStatus int

//SkuVisibility : product visibility enum
type SkuVisibility int

const (
	//SkuDisalbed status
	SkuDisalbed SkuStatus = 0
	//SkuEnabled status
	SkuEnabled SkuStatus = 1

	//SkuInvisible status
	SkuInvisible SkuVisibility = 0
	//SkuVisible status
	SkuVisible SkuVisibility = 1

	//SKUMODEL product model constant
	SKUMODEL = "sku"
)

//Sku represents sku
type Sku struct {
	ID               uint `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Name             string
	Description      string
	Price            float64
	OfferPrice       float64
	Visibility       SkuVisibility
	Status           SkuStatus
	EventID          uint
	ProductID        uint
	VerticalID       uint
	PayTypeSupported string
	LongRichDesc     string
	Tags             interface{}
	MerchantID       string
}

//Create sku
func (s *Sku) Create() (err error) {
	errs := gormDb.Create(s).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", SKUMODEL, errs[0].Error())
	}
	return
}

//Get sku
func (s *Sku) Get() (err error) {
	return
}
