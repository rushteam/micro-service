package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//AppConfig ..
type AppConfig struct {
	Apps map[string]struct {
		Secret   string `yaml:"secret"`
		Channels map[string]struct {
			Name    string `yaml:"name"`
			PvdName string `yaml:"pvd_name"`
			AppID   string `yaml:"app_id"`
			ApiKey  string `yaml:"api_key"`
			MchID   string `yaml:"mch_id"`
			//TradeType string `yaml:"trade_type"`
		} `yaml:"channels"`
	} `yaml:"apps"` /* app_id:secret*/
}

//App ..
var App AppConfig

//Load ..
func (c *AppConfig) Load(fileName string) error {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	return err
}
