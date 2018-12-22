package model

import (
	"time"

	"github.com/sanchitlohia2711/db-mysql/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//Brand represents a brand
type Brand struct {
	ID          int64     `xorm:"'id' pk autoincr"`
	Name        string    `xorm:"'name'"`
	Description string    `xorm:"'description'"`
	Active      int       `xorm:"'active'"`
	CreatedAt   time.Time `xorm:"'created_at' not null"`
	UpdatedAt   time.Time `xorm:"'updated_at' not null"`
}

//NewBrand creates a new brand
func NewBrand(name, description string) (brand *Brand, err error) {
	b := &Brand{}
	b.Name = name
	b.Description = description
	err = CreateBrand(b)
	return
}

//CreateBrand  create the brand
func CreateBrand(brand *Brand) (err error) {
	now := time.Now()

	brand.CreatedAt = now
	brand.UpdatedAt = now

	err = db.InsertOne("brand", brand)
	if err != nil {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", err.Error())
	}
	return
}
