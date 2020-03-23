package model

//OAuthChannel ..
type OAuthChannel struct {
	Name     string `yaml:"name"`
	Provider string `yaml:"provider"`
	AppID    string `yaml:"app_id"`
	Secret   string `yaml:"secret"`
}

//OAuthChanels ..
var OAuthChanels map[string]*OAuthChannel

func init() {
	OAuthChanels = make(map[string]*OAuthChannel, 10)
	OAuthChanels["wx"] = &OAuthChannel{
		Name:     "微信公众号",
		Provider: "wx",
		AppID:    "wx5e596d33cb663cd1",
		Secret:   "",
	}
}
