package wxsdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//GetURL ..
func GetURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

//PostURL ...
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
