package wxsdk

import (
	"encoding/json"
	"fmt"
)

var baseURL = "https://api.weixin.qq.com"

//AccessToken ..
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//GetToken ..
func GetToken(appID string, secret string) (AccessToken, error) {
	var accessToken AccessToken
	url := fmt.Sprintf(
		"%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		baseURL, appID, secret)
	body, err := GetURL(url)
	if err != nil {
		return accessToken, err
	}
	err = json.Unmarshal(body, &accessToken)
	return accessToken, err
}

//AuthAccessToken ...
type AuthAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

//GetAuthAccessToken ..
func GetAuthAccessToken(appID, secret, code string) (AuthAccessToken, error) {
	var accessToken AuthAccessToken
	url := fmt.Sprintf(
		"%s/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=%s",
		baseURL, appID, secret, code, "authorization_code")
	body, err := GetURL(url)
	if err != nil {
		return accessToken, err
	}
	err = json.Unmarshal(body, &accessToken)
	return accessToken, err
}

//AuthRefreshToken ...
func AuthRefreshToken(appID, refreshToken string) (AuthAccessToken, error) {
	var accessToken AuthAccessToken
	url := fmt.Sprintf(
		"%s/sns/oauth2/refresh_token?appid=%s&refresh_token=%s&grant_type=%s",
		baseURL, appID, refreshToken, "refresh_token")
	body, err := GetURL(url)
	if err != nil {
		return accessToken, err
	}
	err = json.Unmarshal(body, &accessToken)
	return accessToken, err
}

//Userinfo ..
type Userinfo struct {
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        string   `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

//GetUserinfo ..
func GetUserinfo(accessToken, openID string) (Userinfo, error) {
	var userinfo Userinfo
	url := fmt.Sprintf(
		"%s/sns/userinfo?access_token=%s&openid=%s&lang=%s",
		baseURL, accessToken, openID, "zh_CN")
	body, err := GetURL(url)
	if err != nil {
		return userinfo, err
	}
	err = json.Unmarshal(body, &userinfo)
	return userinfo, err
}
