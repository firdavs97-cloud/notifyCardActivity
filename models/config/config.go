package config

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Port      int    `yaml:"port"`
	DBAddress string `yaml:"db_address"`
}

func InitConfig() (*Config, error) {
	cfg := &Config{}
	file := "./config/default.yaml"
	yamlFile, err := ioutil.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(yamlFile)
	dec := yaml.NewDecoder(r)
	if err := dec.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, err
}
