package wxsdk

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const openURL = "https://open.weixin.qq.com"
const apiURL = "https://api.weixin.qq.com"

//NewOAuth 创建微信授权
func NewOAuth(appID, secret string) OAuth {
	return &wxOAuth{appID: appID, secret: secret}
}

type wxOAuth struct {
	appID  string
	secret string
}

//String 获取当前方式名
func (o *wxOAuth) String() string {
	return "wx_open"
}

//scope: snsapi_login,snsapi_base,snsapi_userinfo
func (o *wxOAuth) getAuthorizeURL(redirect, state, scope string) string {
	url := fmt.Sprintf("%s/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=%s&scope=%s&state=%s#wechat_redirect",
		openURL, o.appID, url.QueryEscape(redirect), "code", scope, state)
	return url
}

//GetUserinfoAuthorizeURL 获取授权url (用户信息) [微信公众号]
func (o *wxOAuth) GetUserinfoAuthorizeURL(redirect, state string) string {
	return o.getAuthorizeURL(redirect, state, "snsapi_userinfo")
}

//GetAuthorizeURL 获取授权url (基本信息) [微信公众号]
func (o *wxOAuth) GetAuthorizeURL(redirect, state string) string {
	return o.getAuthorizeURL(redirect, state, "snsapi_base")
}

//GetAuthorizeQrURL 获取授权二维码url [微信PC扫码]
//https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
func (o *wxOAuth) GetAuthorizeQrURL(redirect, state string) string {
	var scope = "snsapi_login" //snsapi_login,snsapi_base,snsapi_userinfo
	url := fmt.Sprintf("%s/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=%s&scope=%s&state=%s#wechat_redirect",
		openURL, o.appID, url.QueryEscape(redirect), "code", scope, state)
	return url
}

//GetAccessToken 获取授权AccessToken
func (o *wxOAuth) GetAccessToken(code string) (*OAuthAccessToken, error) {
	url := fmt.Sprintf(
		"%s/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=%s",
		apiURL, o.appID, o.secret, code, "authorization_code")
	body, err := getURL(url)
	if err != nil {
		return nil, err
	}
	accessToken := &OAuthAccessToken{}
	err = json.Unmarshal(body, accessToken)
	if accessToken.ErrInfo.ErrCode != 0 {
		return nil, accessToken.ErrInfo
	}
	return accessToken, err
}

//RefreshToken 刷新授权AccessToken
func (o *wxOAuth) RefreshToken(refreshToken string) (*OAuthAccessToken, error) {
	url := fmt.Sprintf(
		"%s/sns/oauth2/refresh_token?appid=%s&refresh_token=%s&grant_type=%s",
		apiURL, o.appID, refreshToken, "refresh_token")
	body, err := getURL(url)
	if err != nil {
		return nil, err
	}
	accessToken := &OAuthAccessToken{}
	err = json.Unmarshal(body, accessToken)
	if accessToken.ErrInfo.ErrCode != 0 {
		return nil, accessToken.ErrInfo
	}
	return accessToken, err
}
