package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Entity servcie struct
type Entity struct {
	*model.Entity
}

//CreateEntity create the entity
func CreateEntity(entityParams *requestDTOV1.EntityParams) (entity *Entity, err error) {
	entityModel := &model.Entity{}
	entityModel.CityID = entityParams.CityID
	entityModel.CountryID = entityParams.CountryID
	entityModel.Description = entityParams.Description
	entityModel.Type = entityParams.Type
	entityModel.Lat = entityParams.Lat
	entityModel.Lon = entityParams.Lon
	entityModel.Name = entityParams.Name
	entityModel.MerchantID = entityParams.MerchantID
	err = entityModel.Create()
	if err != nil {
		return
	}
	entity = &Entity{entityModel}
	return
}
