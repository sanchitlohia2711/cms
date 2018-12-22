package main

import (
	"fmt"

	"github.com/ko/cms-db/adminService/adminmodel"
	"github.com/ko/cms-db/err"
)

func main() {
	err.Initialize()
	e := adminmodel.TestModel()
	fmt.Println(e)
	s := "b"
	fmt.Println(s)
}
