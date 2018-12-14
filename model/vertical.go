package model

import (
	"fmt"
	"time"

	"github.com/sanchitlohia2711/db-mysql/db"
	eError "github.com/sanchitlohia2711/go-extended-error/err"
)

//VERTICALSTATUS : order status enum
type VERTICALSTATUS int

const (
	//Init status
	Init VERTICALSTATUS = 0
)

//Vertical represents vertical
type Vertical struct {
	ID          int64     `xorm:"'id' pk autoincr"`
	Name        string    `xorm:"'name'"`
	Description string    `xorm:"'description'"`
	Active      int       `xorm:"'active'"`
	CreatedAt   time.Time `xorm:"'created_at' not null"`
	UpdatedAt   time.Time `xorm:"'updated_at' not null"`
}

//CreateVertical  create the vertical
func CreateVertical(vertical *Vertical) (err error) {
	now := time.Now()

	vertical.CreatedAt = now
	vertical.UpdatedAt = now

	err = db.InsertOne("vertical", vertical)
	e := eError.NewExtendedError("TEST_101", err.Error())
	fmt.Println(e)
	if err != nil {

	}
	return
}
