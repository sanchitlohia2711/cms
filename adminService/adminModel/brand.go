package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

//Brand struct
type Brand struct {
	*model.Brand
}

//CreateBrand create the brand
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

//CreateBrandIfNotExists create the brand if not exits
func CreateBrandIfNotExists(brandParams *requestDTOV1.BrandParams) (*model.Brand, error) {

	filters := &model.BrandFilters{model.Brand{Name: brandParams.Name}}
	filters.Active = 1

	brands, err := model.GetBrand(filters)
	if err != nil {
		return nil, err
	}
	if len(brands) != 1 {
		return nil, exerr.NewExtendedError("SQL_GET_ERROR")
	}
	if len(brands) == 1 {
		return brands[0], nil
	}
	brandModel := &model.Brand{}
	brandModel.Name = brandParams.Name
	brandModel.Description = brandParams.Description
	brandModel.Active = 1
	if err = brandModel.Create(); err != nil {
		return nil, err
	}

	return
}

//GetBrand get brand
func GetBrand(brandFilters *BrandFilters) ([]*Brand, error) {
	brands, err = brandModel.Get()
	if err != nil {
		return nil, err
	}
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
