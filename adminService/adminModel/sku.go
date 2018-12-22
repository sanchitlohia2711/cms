package adminmodel

import (
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Sku service struct
type Sku struct {
	*model.Sku
}

//CreateSku create the sku
func CreateSku(params *requestDTOV1.SkuParams) (sku *Sku, err error) {
	sModel := &model.Sku{}
	sModel.Name = params.Name
	sModel.Description = params.Description
	sModel.Tags = params.Tags
	sModel.Visibility = model.SkuVisible
	sModel.Status = model.SkuEnabled
	sModel.LongRichDesc = params.LongRichDesc
	sModel.PayTypeSupported = params.PayTypeSupported
	sModel.ProductID = params.ProductID
	sModel.Tags = params.Tags
	sModel.MerchantID = params.MerchantID
	sModel.EventID = params.EventID
	sModel.VerticalID = params.VerticalID

	if params.Price > 0 {
		sModel.Price = params.Price
	}
	if params.OfferPrice > 0 {
		sModel.OfferPrice = params.OfferPrice
	}
	err = sModel.Create()
	if err != nil {
		return
	}
	sku = &Sku{sModel}
	return
}
