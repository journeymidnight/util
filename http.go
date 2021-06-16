package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// read from ReadCloser and unmarshal to out;
// `out` should be of POINTER type
func ReadJsonBody(body io.ReadCloser, out interface{}) (err error) {
	defer func() {
		_ = body.Close()
	}()
	jsonBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	// fmt.Println("json body:", string(jsonBytes))
	err = json.Unmarshal(jsonBytes, out)
	if err != nil {
		return err
	}
	return nil
}
