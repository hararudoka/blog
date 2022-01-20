package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port string `yaml:"port"`
}

// ConfigInit not used at the moment
func (c *Config) ConfigInit() error {
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, c)
	return err
}
