package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Country service struct
type Country struct {
	*model.Country
}

//CreateCountry create the country
func CreateCountry(params *requestDTOV1.CountryParams) (country *Country, err error) {
	c := &model.Country{}
	c.Name = params.Name
	c.Description = params.Description
	c.Active = 1
	err = c.Create()
	if err != nil {
		return
	}
	country = &Country{c}
	return
}
