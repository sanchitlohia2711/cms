package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//CreateCountry create the country
func CreateCountry(countryParams *requestDTOV1.CountryParams) (country *model.Country, err error) {
	country = &model.Country{}
	country.Name = countryParams.Name
	country.Description = countryParams.Description
	country.Active = 1
	err = country.Create()
	return
}
