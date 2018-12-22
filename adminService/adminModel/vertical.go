package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Vertical service
type Vertical struct {
	*model.Vertical
}

//CreateVertical create the vertical
func CreateVertical(b *requestDTOV1.VerticalParams) (vertical *Vertical, err error) {
	verticalModel := &model.Vertical{}
	verticalModel.Name = b.Name
	verticalModel.Description = b.Description
	verticalModel.Active = 1
	if err = verticalModel.Create(); err != nil {
		return
	}
	vertical = &Vertical{verticalModel}
	return
}
