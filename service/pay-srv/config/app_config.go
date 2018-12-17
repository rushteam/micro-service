package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//AppConfig ..
type AppConfig struct {
	Apps map[string]struct {
		Secret      string   `yaml:"secret"`
		PayChannels []string `yaml:"channels"`
	} `yaml:"apps"`
	PayChannels map[string]struct {
		Name      string `yaml:"name"`
		Provider  string `yaml:"provider"`
		AppID     string `yaml:"app_id"`
		ApiKey    string `yaml:"api_key"`
		MchID     string `yaml:"mch_id"`
		NotifyURL string `yaml:"notify_url"`
	} `yaml:"pay_channels"`
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
