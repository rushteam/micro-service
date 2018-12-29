package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func PostURL(url string, params []byte) ([]byte, error) {
	paramsReader := bytes.NewReader(params)
	resp, err := http.Post(url, "application/xml", paramsReader)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
