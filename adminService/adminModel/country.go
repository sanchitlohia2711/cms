package adminModel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//CreateCountry create the country
func CreateCountry(countryParams *requestDTOV1.CountryParams) (err error) {
	c := &model.Country{}
	c.Name = countryParams.Name
	c.Description = countryParams.Description
	c.Active = 1
	return c.Create()
}
