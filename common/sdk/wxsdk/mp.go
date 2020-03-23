package wxsdk

import (
	"encoding/json"
	"fmt"
)

//AccessToken 微信公众号access_token
type AccessToken struct {
	*ErrInfo
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//Request 请求 微信公众号access_token
func (at *AccessToken) Request(appID, secret string) error {
	url := fmt.Sprintf(
		"%s/cgi-bin/token?grant_type=%s&appid=%s&secret=%s",
		apiURL, "client_credential", appID, secret)
	body, err := getURL(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, at)
	if at.ErrInfo.ErrCode != 0 {
		return at.ErrInfo
	}
	return nil
}
