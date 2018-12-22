package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

func HandleError(w http.ResponseWriter, err error) {
	e, ok := err.(*exerr.ExtendedError)
	if !ok {
		// log.Printf("Casting to ExtendedError failed: %+v\n", err)
		return
	}

	resp := make(map[string]interface{}, 0)
	resp["code"] = e.HTTPStatus
	resp["message"] = e.Message
	resp["customer_message"] = e.CustomerMessage

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		//log.Printf("Failed to encode error information %+v as JSON: %+v\n", resp, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPStatus)
	w.Write(jsonResp)
}
