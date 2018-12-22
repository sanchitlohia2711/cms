package adminModel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
)

func TestModel() (err error) {
	err = TestCountry()
	if err != nil {
		return
	}
	// err = TestCity()
	// if err != nil {

	// }
	return

	// TestBrand()
	// TestBrandVerticalMapping()
	// TestCategory()
	// TestCategoryFilters()

}

func TestCountry() (err error) {
	countryParams := &requestDTOV1.CountryParams{}
	countryParams.Name = "country"
	countryParams.Description = "countryDescription"
	err = CreateCountry(countryParams)
	if err != nil {
		return
	}
	return
}

// func TestCity() {
// 	cityParams := requestDTOV1.CityParams{}
// 	cityParams.Name = "somename"
// }
