package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//GetRequestBody get request body
func GetRequestBody(r *http.Request, req interface{}) (err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return
	}
	return
}
