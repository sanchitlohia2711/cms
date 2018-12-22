package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//CreateVertical create the vertical
func CreateVertical(b *requestDTOV1.VerticalParams) (vertical *model.Vertical, err error) {
	vertical = &model.Vertical{}
	vertical.Name = b.Name
	vertical.Description = b.Description
	vertical.Active = 1
	if err = vertical.Create(); err != nil {
		return
	}
	return
}
