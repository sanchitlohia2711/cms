package model

import (
	"context"

	"github.com/ko/cms-db/cms-es/es"
	esModel "github.com/ko/cms-db/cms-es/model"
	dbModel "github.com/ko/cms-db/model"
)

//Product ES struct
type Product struct {
	*esModel.Product
}

var (
	esClient = es.GetClient()
)

//CreateProductES ES document
func CreateProductES(product *dbModel.Product) (err error) {
	p := &esModel.Product{}
	p.Name = product.Name
	_, err = esClient.Index().
		Index("pro").
		Type("default").
		Id("1").
		BodyJson(p).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		return
	}
	return
}
