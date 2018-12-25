package adminmodel

import (
	"encoding/json"
	"strings"
	"time"

	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

//Product struct
type Product struct {
	*model.Product
}

//CreateProduct create the country
func CreateProduct(productParams *requestDTOV1.ProductParams) (product *Product, err error) {
	productModel := &model.Product{}
	productModel.Name = productParams.Name
	productModel.Description = productParams.Description
	productModel.Tags = strings.Join(productParams.Tags, ",")
	productModel.VerticalID = productParams.VerticalID
	productModel.Visibility = model.ProductVisible
	productModel.Status = model.ProductEnabled
	productModel.BrandID = productParams.BrandID
	productModel.ImageURL = productParams.ImageURL
	productModel.ShareURL = productParams.ShareURL
	productModel.ThumbURL = productParams.ThumbURL
	productModel.InputFields = "{}"
	productModel.Attributes = "{}"
	howToRedeem, err := json.Marshal(productParams.HowToRedeem)
	if err != nil {
		return
	}
	productModel.HowToRedeem = string(howToRedeem)
	termsConditions, err := json.Marshal(productParams.TermsConditions)
	if err != nil {
		return
	}
	productModel.TermsConditions = string(termsConditions)
	returnPolicy, err := json.Marshal(productParams.ReturnPolicy)
	if err != nil {
		return
	}
	productModel.ReturnPolicy = string(returnPolicy)
	if productParams.StartDate.IsZero() {
		productModel.StartDate = time.Now()
	} else {
		productModel.StartDate = time.Now()
	}
	if productParams.EndDate.IsZero() {
		productModel.EndDate = time.Now()
	} else {
		productModel.EndDate = time.Now()
	}

	err = productModel.Create()
	if err != nil {
		return
	}
	product = &Product{productModel}
	return
}

//AssociateCategories with the product
func (product *Product) AssociateCategories(categories []*Category) (err error) {
	productCategoryMapping := &model.ProductCategoryMapping{}
	productCategoryMapping.ProductID = product.ID
	productCategoryMapping.CategoryID = categories[0].ID
	productCategoryMapping.Active = 1
	err = productCategoryMapping.Create()
	return
}
