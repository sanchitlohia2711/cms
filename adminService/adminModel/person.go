package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Person servcie struct
type Person struct {
	*model.Person
}

//CreatePerson create the entity
func CreatePerson(params *requestDTOV1.PersonParams) (err error) {
	pModel := &model.Person{}
	pModel.Name = params.Name
	pModel.Description = params.Description
	err = pModel.Create()
	if err != nil {
		return
	}
	return
}
