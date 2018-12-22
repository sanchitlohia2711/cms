package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Base struct
type Base struct {
	Active    int       `xorm:"'active'"`
	CreatedAt time.Time `xorm:"'created_at' not null"`
	UpdatedAt time.Time `xorm:"'updated_at' not null"`
}

var (
	gormDb *gorm.DB
)

//NewBase return new base
func NewBase() (b Base) {
	b = Base{}
	return
}
