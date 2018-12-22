package main

import (
	"fmt"

	"github.com/ko/cms-db/err"
	"github.com/ko/cms-db/model"
)

func main() {
	err.Initialize()
	_, err := model.NewVertical("somename", "somename")
	fmt.Println(err)
	fmt.Println(err)
}
