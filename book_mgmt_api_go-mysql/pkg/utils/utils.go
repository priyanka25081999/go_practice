package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// unmarshell the json received from the user to be used by our controller code
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// if no err, then start unmarshalling
		if err := json.Unmarshal([]byte(body), x); err != nil {
			fmt.Println("Error occured during JSON unmarshelling, err:", err)
			return
		}
	}
}
