package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//CreateBrand create the country
func CreateBrand(b *requestDTOV1.BrandParams) (brand *model.Brand, err error) {
	brand = &model.Brand{}
	brand.Name = b.Name
	brand.Description = b.Description
	brand.Active = 1
	err = brand.Create()
	return
}

//CreateBrandVericalMapping : associates a brand with a vertical
func CreateBrandVericalMapping(brandID uint, verticalID uint) (err error) {
	brandVerticalMapping := &model.BrandVerticalMapping{}
	brandVerticalMapping.BrandID = brandID
	brandVerticalMapping.VerticalID = verticalID
	return brandVerticalMapping.Create()
}
