package model

import (
	"time"

	"github.com/sanchitlohia2711/db-mysql/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//BrandVerticalMapping represents a brand
type BrandVerticalMapping struct {
	ID         int       `xorm:"'id' pk autoincr"`
	BrandID    int       `xorm:"'brand_id'"`
	VerticalID int       `xorm:"'vertical_id'"`
	Active     int       `xorm:"'active'"`
	CreatedAt  time.Time `xorm:"'created_at' not null"`
	UpdatedAt  time.Time `xorm:"'updated_at' not null"`
}

//CreateBrandVerticalMapping  create the brand
func CreateBrandVerticalMapping(brandID, verticalID int) (err error) {
	now := time.Now()

	brandVerticalMapping := &BrandVerticalMapping{}
	brandVerticalMapping.CreatedAt = now
	brandVerticalMapping.UpdatedAt = now

	err = db.InsertOne("brand_vertical_mapping", brandVerticalMapping)
	if err != nil {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", "BrandVerticalMapping", err.Error())
	}
	return
}

//DeactivateBrandVerticalMapping deactivates the mapping
func DeactivateBrandVerticalMapping(brandID, verticalID int) (err error) {
	// sess := db.NewSqlSession()
	// defer sess.Close()

	// sess = sess.Table("eligible_transaction")

	// _, err = sess.ID(e.ID).Update(e)
	// return
	return
}
