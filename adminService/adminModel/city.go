package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//City service struct
type City struct {
	*model.City
}

//CreateCity create the country
func CreateCity(cityParams *requestDTOV1.CityParams) (city *City, err error) {
	c := &model.City{}
	c.Name = cityParams.Name
	c.Description = cityParams.Description
	c.CountryID = cityParams.CountryID
	c.Active = 1
	err = c.Create()
	if err != nil {
		return
	}
	city = &City{c}
	return
}
