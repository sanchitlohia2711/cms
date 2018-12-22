package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Brand struct
type Brand struct {
	*model.Brand
}

//CreateBrand create the country
func CreateBrand(brandParams *requestDTOV1.BrandParams) (brand *Brand, err error) {
	brandModel := &model.Brand{}
	brandModel.Name = brandParams.Name
	brandModel.Description = brandParams.Description
	brandModel.Active = 1
	if err = brandModel.Create(); err != nil {
		return
	}
	brand = &Brand{brandModel}
	return
}

//CreateBrandVericalMapping : associates a brand with a vertical
func CreateBrandVericalMapping(brandID uint, verticalID uint) (err error) {
	brandVerticalMapping := &model.BrandVerticalMapping{}
	brandVerticalMapping.BrandID = brandID
	brandVerticalMapping.VerticalID = verticalID
	return brandVerticalMapping.Create()
}

//AssociateVertical : associates a brand with a vertical
func (brand *Brand) AssociateVertical(verticalID uint) (err error) {
	brandVerticalMapping := &model.BrandVerticalMapping{}
	brandVerticalMapping.BrandID = brand.ID
	brandVerticalMapping.VerticalID = verticalID
	return brandVerticalMapping.Create()
}
