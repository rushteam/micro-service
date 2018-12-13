package wxpay

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"gitee.com/rushteam/micro-service/common/utils"
	"io/ioutil"
	"net/http"
)

//MakeSign ..
func MakeSign(APIKey string,v interface{}) string {
	var params map[string]string
	params = utils.Struct2Map(v, "json")
	return utils.Sign(params, "", fmt.Sprintf("&key=%s", APIKey), md5.New())
}
//CheckSign ..
func CheckSign(APIKey string,v interface{}) bool {
	var params map[string]string
	params = utils.Struct2Map(v, "json")
	var sign string
	if params["sign_type"] == "MD5" {
		sign = utils.Sign(params, "", fmt.Sprintf("&key=%s", APIKey), md5.New())
	}
	if params["sign"] == sign {
		return  true
	}
	return false
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
