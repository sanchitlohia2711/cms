package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

func TestModel() (err error) {
	country, err := TestCountry()
	if err != nil {
		return
	}
	err = TestCity(country)
	if err != nil {
		return
	}

	brand, err := TestBrand()
	if err != nil {
		return
	}

	vertical, err := TestVertical()
	if err != nil {
		return
	}

	if err = CreateBrandVericalMapping(brand.ID, vertical.ID); err != nil {
		return
	}
	return
	// TestBrandVerticalMapping()
	// TestCategory()
	// TestCategoryFilters()

}

func TestCountry() (country *model.Country, err error) {
	countryParams := &requestDTOV1.CountryParams{}
	countryParams.Name = "country"
	countryParams.Description = "countryDescription"
	country, err = CreateCountry(countryParams)
	return
}

func TestCity(country *model.Country) (err error) {
	cityParams := &requestDTOV1.CityParams{}
	cityParams.Name = "city"
	cityParams.Description = "cityDescription"
	cityParams.CountryID = country.ID
	err = CreateCity(cityParams)
	return
}

func TestBrand() (brand *model.Brand, err error) {
	brandParams := &requestDTOV1.BrandParams{}
	brandParams.Name = "brandName"
	brandParams.Description = "brandDescription"
	return CreateBrand(brandParams)
}

func TestVertical() (vertical *model.Vertical, err error) {
	verticalParams := &requestDTOV1.VerticalParams{}
	verticalParams.Name = "verticalName"
	verticalParams.Description = "verticalDescription"
	return CreateVertical(verticalParams)
}
