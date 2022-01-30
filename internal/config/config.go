package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {

}

// LoadConfig not used at the moment
func LoadConfig() (c *Config, err error) {
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, c)
	return c, err
}
