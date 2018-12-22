package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//CreateCity create the country
func CreateCity(cityParams *requestDTOV1.CityParams) (err error) {
	c := &model.City{}
	c.Name = cityParams.Name
	c.Description = cityParams.Description
	c.CountryID = cityParams.CountryID
	c.Active = 1
	return c.Create()
}
