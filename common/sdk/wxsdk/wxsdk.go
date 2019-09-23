package wxsdk

import (
	"context"
	"encoding/json"
	"fmt"
)

const baseURL = "https://api.weixin.qq.com"

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

//ErrInfo ..
type ErrInfo struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

//AuthAccessToken ...
type AuthAccessToken struct {
	*ErrInfo
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

//GetAuthAccessToken ..
func GetAuthAccessToken(ctx context.Context, appID, secret, code string) (AuthAccessToken, error) {
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
func AuthRefreshToken(ctx context.Context, appID, refreshToken string) (AuthAccessToken, error) {
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

//GetUserinfo ..
func GetUserinfo(ctx context.Context, accessToken, openID string) (Userinfo, error) {
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
