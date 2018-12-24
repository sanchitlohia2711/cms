package model

import (
	"context"
	"strings"

	"github.com/ko/cms-db/cms-es/es"
	esModel "github.com/ko/cms-db/cms-es/model"
	dbModel "github.com/ko/cms-db/model"
)

//Product ES struct
type Product struct {
	*esModel.ProductES
}

var (
	esClient = es.GetClient()
)

//CreateProductES ES document
func CreateProductES(product *dbModel.Product) (err error) {
	p := &esModel.ProductES{}
	p.Name = product.Name
	p.Description = product.Description
	p.EndDate = product.EndDate
	p.HowToRedeem = product.HowToRedeem
	p.ImageURL = product.ImageURL
	p.ShareURL = product.ShareURL
	p.ThumbURL = product.ThumbURL
	p.InputFields = product.InputFields
	p.StartDate = product.StartDate
	p.TermsConditions = product.TermsConditions
	p.Status = uint8(product.Status)
	p.Visibility = uint8(product.Visibility)
	p.Tags = strings.Split(product.Tags, ",")
	p.Attributes = product.Attributes
	p.BrandID = product.BrandID
	p.VerticalID = product.VerticalID
	p.ReturnPolicy = product.ReturnPolicy
	_, err = esClient.Index().
		Index("product").
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
