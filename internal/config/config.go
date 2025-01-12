// Package config
package config

import (
	"os"

	"github.com/pashifika/util/files"
	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*Config, error) {
	if !files.Exists(path) {
		return nil, ErrConfigNotFound
	}
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config *Config
	err = yaml.Unmarshal(buf, &config)

	return config, err
}
