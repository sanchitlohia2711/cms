package model

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/ko/cms-db/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//Category represents category
type Category struct {
	Base
	Name        string
	Description string
	VerticalID  uint
	ParentID    uint
}

//CategoryFilters represents category filters
type CategoryFilters struct {
	Base
	Name        string
	Description string
	VerticalID  uint
	ParentID    uint
}

const (
	//CATEGORYMODEL category model constant
	CATEGORYMODEL = "category"
)

func init() {
	gormDb = db.GormDB()
}

//Create category
func (c *Category) Create() (err error) {
	errs := gormDb.Create(c).GetErrors()
	if len(errs) > 0 {
		err = exerr.NewExtendedError("SQL_INSERT_ERROR", CATEGORYMODEL, errs[0].Error())
	}
	return
}

//Get category
func (c *Category) Get() (err error) {
	return
}

//GetCategoryByIDs get category by Id
func GetCategoryByIDs(categoryIDs []uint) (categoryArray []*Category, err error) {
	var db *gorm.DB
	if len(categoryIDs) > 0 {
		db = gormDb.Where("id in (?)", strings.Trim(strings.Replace(fmt.Sprint(categoryIDs), " ", ",", -1), "[]"))
	}
	errs := db.Find(&categoryArray).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs[0])
		err = exerr.NewExtendedError("SQL_GET_ERROR", SKUPRODUCTENTITYMODEL, errs[0].Error())
	}
	return
}

//GetCategoryByID get category by Id
func GetCategoryByID(categoryID uint) (category *Category, err error) {
	var db *gorm.DB
	if categoryID > 0 {
		db = gormDb.Where("category_id = ?", categoryID)
	}
	errs := db.Find(&category).GetErrors()
	if len(errs) > 0 {
		fmt.Println(errs[0])
		err = exerr.NewExtendedError("SQL_GET_ERROR", SKUPRODUCTENTITYMODEL, errs[0].Error())
	}
	return
}

//GetParentCategoryIds get parent category ids
func (c *Category) GetParentCategoryIds(inclusiveCurrent bool) (parentCategoryIds []uint, err error) {
	if inclusiveCurrent {
		parentCategoryIds = append(parentCategoryIds, c.ID)
	}
	if c.ParentID == 0 {
		return
	}

	parentCategory, err := GetCategoryByID(c.ParentID)
	if err != nil {
		return
	}
	parentIds, err := parentCategory.GetParentCategoryIds(true)
	if err != nil {
		return
	}
	for _, parentCategoryID := range parentIds {
		parentCategoryIds = append(parentCategoryIds, parentCategoryID)
	}
	return
}
