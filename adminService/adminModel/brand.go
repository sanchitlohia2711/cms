package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Brand struct
type Brand struct {
	*model.Brand
}

type BrandFilters struct {
	*model.BrandFilters
}

//CreateBrand create the brand
func CreateBrand(brandParams *requestDTOV1.BrandParams) (*Brand, error) {
	brandModel := &model.Brand{}
	brandModel.Name = brandParams.Name
	brandModel.Description = brandParams.Description
	brandModel.Active = 1
	if err := brandModel.Create(); err != nil {
		return nil, err
	}
	brand := &Brand{brandModel}
	return brand, nil
}

//CreateBrandIfNotExists create the brand if not exits
func CreateBrandIfNotExists(brandParams *requestDTOV1.BrandParams) (*Brand, error) {

	filters := &BrandFilters{
		BrandFilters: &model.BrandFilters{},
	}
	filters.Name = brandParams.Name
	filters.Active = 1

	brands, err := GetBrand(filters)
	if err != nil {
		return nil, err
	}

	if len(brands) == 1 {
		return brands[0], nil
	}
	return CreateBrand(brandParams)
}

//GetBrand get brand
func GetBrand(brandFilters *BrandFilters) ([]*Brand, error) {
	results, err := model.GetBrand(brandFilters.BrandFilters)
	if err != nil {
		return nil, err
	}
	var brands []*Brand
	for _, result := range results {
		brand := &Brand{result}
		brands = append(brands, brand)
	}
	return brands, nil
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
