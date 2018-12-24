package providers

import (
	"github.com/ko/cms-db/model"
)

//ProductES struct
type ProductES struct {
	Product                 *model.Product
	SkuProductEntityMapping []*model.SkuProductEntityMapping
	Category                [][]uint
}

//NewProductES create a new product ES struct
func NewProductES(product *model.Product) (productES *ProductES) {

	productES.Product = product
    productES.
	a := make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
	}
}
