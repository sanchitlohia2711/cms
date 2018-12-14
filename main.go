package main

import (
	"fmt"

	"github.com/ko/cms-db/db"
	"github.com/ko/cms-db/err"
	"github.com/ko/cms-db/model"
)

func main() {
	err.Initialize()
	db.Initialize()
	v := &model.Vertical{}
	v.Name = "somename"
	err := model.CreateVertical(v)
	fmt.Println(err)
}
