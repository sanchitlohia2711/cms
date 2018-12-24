package model

import (
	dbModel "github.com/ko/cms-db/model"
)

func TestESModel() (err error) {
	productFilters := &dbModel.ProductFilterParams{}
	productFilters.ID = 5
	products, err := dbModel.GetProduct(productFilters)
	if err != nil {
		return
	}
	err = TestProduct(products[0])
	if err != nil {
		return
	}
	return
}

func TestProduct(product *dbModel.Product) (err error) {
	return CreateProductES(product)
}
