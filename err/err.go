package err

import (
	"fmt"
	"path/filepath"

	"github.com/sanchitlohia2711/go-extended-error/err"
)

const (
	//ErrFailedToLoadErrorJSON message
	ErrFailedToLoadErrorJSON = "Failed to load error json %v"
)

//Initialize : initialize the db
func Initialize() {
	errorFilePath := "./err/error.json"
	absPath, _ := filepath.Abs(errorFilePath)
	err := err.Initialize(absPath)
	if err != nil {
		err = fmt.Errorf(ErrFailedToLoadErrorJSON, err.Error())
		panic(err.Error())
	}
}
