package err

import (
	"fmt"
	"path/filepath"

	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

const (
	//ErrFailedToLoadErrorJSON message
	ErrFailedToLoadErrorJSON = "Failed to load error json %v"
)

//Initialize : initialize the db
func Initialize() {
	errorFilePath := "./err/error.json"
	absPath, _ := filepath.Abs(errorFilePath)
	err := exerr.Initialize(absPath)
	if err != nil {
		err = fmt.Errorf(ErrFailedToLoadErrorJSON, err.Error())
		panic(err.Error())
	}
}
