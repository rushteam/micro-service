package wxsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//getURL ..
func getURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

//postURL ...
func postURL(url string, params []byte) ([]byte, error) {
	paramsReader := bytes.NewReader(params)
	resp, err := http.Post(url, "application/xml", paramsReader)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

//Userinfo ..
type Userinfo struct {
	*ErrInfo
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        string   `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
	Language   string   `json:"language"`
}

//Request 请求用户信息
func (u *Userinfo) Request(accessToken, openID string) error {
	url := fmt.Sprintf(
		"%s/sns/userinfo?access_token=%s&openid=%s&lang=%s",
		apiURL, accessToken, openID, "zh_CN")
	body, err := getURL(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, u)
	if u.ErrInfo.ErrCode != 0 {
		return u.ErrInfo
	}
	return u
}

//GetCode2Session
