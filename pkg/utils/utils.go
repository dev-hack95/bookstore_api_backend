package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Marshelling and Unmarshelling data
// Marshal takes any value ( any is a wrapper around interface{} ) and converts it into a byte slice.
// Unmarshal takes a byte slice, parses it, and stores the result to v

func ParseBody(r *http.Request, x interface{}) {
	// The code parse the body of an http request
	// If there is a body it reads it in and unmarshals the json into x
	// The code parses the body of an http request and stores it in a interface x
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
