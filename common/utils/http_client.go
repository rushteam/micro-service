package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//HttpPost ..
func HttpPost(url string, params []byte) (int, []byte, error) {
	paramsReader := bytes.NewReader(params)
	resp, err := http.Post(url, "application/xml", paramsReader)
	if err != nil {
		return 0, []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, body, err
}
