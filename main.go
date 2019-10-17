package main

import (
	"fmt"

	adminmodel "github.com/ko/cms-db/adminService/adminModel" //indexCreate "github.com/ko/cms-db/cms-es/resources/scripts/mapping/create"
	"github.com/ko/cms-db/err"
)

func main() {
	err.Initialize(	e := adminmodel.TestModel()
	//e := esModel.TestESModel()
	//e := indexCreate.ProductMapping()
	fmt.Println(e)
	s := "b"
	fmt.Println(s)
}
