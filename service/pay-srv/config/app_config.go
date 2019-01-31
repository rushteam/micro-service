package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//DbConfigs ...
type DbConfigs map[string]DbConfig

//DbConfig ..
type DbConfig struct {
	DbType string   `yaml:"db_type"`
	Nodes  []string `yaml:"nodes"`
}

//Default ..
func (dbs DbConfigs) Default() (*DbConfig, error) {
	return dbs.Get("default")
}

//Get ..
func (dbs DbConfigs) Get(name string) (*DbConfig, error) {
	conf, ok := dbs[name]
	if !ok {
		return nil, fmt.Errorf("%s not found", name)
	}
	return &conf, nil
}

//AppConfig ..
type AppConfig struct {
	Apps map[string]struct {
		Secret      string   `yaml:"secret"`
		PayChannels []string `yaml:"pay_channels"`
	} `yaml:"apps"`
	PayChannels map[string]struct {
		Name      string `yaml:"name"`
		Provider  string `yaml:"provider"`
		AppID     string `yaml:"app_id"`
		ApiKey    string `yaml:"api_key"`
		MchID     string `yaml:"mch_id"`
		NotifyURL string `yaml:"notify_url"`
	} `yaml:"pay_channels"`
	Db DbConfigs `yaml:"db_configs"`
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
