package model

import "github.com/ko/cms-db/model"

//Product struct service
type Product struct {
	*model.Product
}

//GetAllEntities of a product
func GetAllEntities(product *Product) (entities []*model.SkuProductEntityMapping, err error) {
	filters := &model.SkuProductEntityMappingFilters{ProductID: product.ID}
	entities, err = model.GetSkuProductEntityMapping(filters)
	return
}

//GetProduct
