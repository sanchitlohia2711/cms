package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Category struct
type Category struct {
	*model.Category
}

//CreateCategory create the country
func CreateCategory(categoryParams *requestDTOV1.CategoryParams) (category *Category, err error) {
	categoryModel := &model.Category{}
	categoryModel.Name = categoryParams.Name
	categoryModel.Description = categoryParams.Description
	categoryModel.Active = 1
	categoryModel.VerticalID = categoryParams.VerticalID
	if categoryParams.ParentID > 0 {
		categoryModel.ParentID = categoryParams.ParentID
	}

	if err = categoryModel.Create(); err != nil {
		return
	}
	category = &Category{categoryModel}
	return
}
