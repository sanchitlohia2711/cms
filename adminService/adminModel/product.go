package adminmodel

import (
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
	productModel.Tags = productParams.Tags
	productModel.VerticalID = productParams.VerticalID
	productModel.Visibility = model.ProductVisible
	productModel.Status = model.ProductEnabled
	productModel.BrandID = productParams.BrandID
	if productParams.StartDate.IsZero() {
		productModel.StartDate = time.Now()
	} else {
		productModel.StartDate = time.Now()
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
	err = productCategoryMapping.Create()
	return
}
