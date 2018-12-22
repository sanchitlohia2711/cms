package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Base struct
type Base struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    uint8
}

var (
	gormDb *gorm.DB
)
