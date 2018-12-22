package main

import (
	"fmt"

	"github.com/ko/cms-db/adminService/adminModel"
	"github.com/ko/cms-db/err"
)

func main() {
	err.Initialize()
	e := adminModel.TestModel()
	fmt.Println(e)
}
