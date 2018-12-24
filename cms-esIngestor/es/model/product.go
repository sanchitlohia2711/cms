package model

import (
	"github.com/ko/cms-db/cms-esingestor/es/model"
	dbModel "github.com/ko/cms-db/model"
)

//Product ES struct
type Product struct {
	*model.Product
}

//Create ES document
func Create(product *dbModel.Product) {

}
