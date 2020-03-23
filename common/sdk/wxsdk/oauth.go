package wxsdk

import "fmt"

//OAuth is a interface
type OAuth interface {
	String() string
	GetAuthorizeURL(redirect, state string) string
	GetUserinfoAuthorizeURL(redirect, state string) string
	GetAuthorizeQrURL(redirect, state string) string
	GetAccessToken(code string) (*OAuthAccessToken, error)
	RefreshToken(refreshToken string) (*OAuthAccessToken, error)
}

//ErrInfo ..
type ErrInfo struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *ErrInfo) Error() string {
	return fmt.Sprintf("wxsdk: [%d] %s", e.ErrCode, e.ErrMsg)
}

//OAuthAccessToken ...
type OAuthAccessToken struct {
	*ErrInfo
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}
