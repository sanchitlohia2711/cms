package main

import (
	"fmt" //esModel "github.com/ko/cms-db/cms-esingestor/model"

	"github.com/ko/cms-db/adminService/adminmodel"
	"github.com/ko/cms-db/err"
)

func main() {
	err.Initialize()
	e := adminmodel.TestModel()
	//e := esModel.TestESModel()
	//e := indexCreate.ProductMapping()
	fmt.Println(e)
	s := "b"
	fmt.Println(s)
}
