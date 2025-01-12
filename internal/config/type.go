// Package config
package config

import "github.com/hashicorp/go-version"

type Config struct {
	Version    version.Version `yaml:"version"`
	Settings   Setting         `yaml:"settings"`
	Generators []Generator     `yaml:"generators"`
}

type Setting struct {
	TemplateRoot string `yaml:"template_root"`
}

type Generator struct {
	Name     string        `yaml:"name"`
	Type     EnumProtoType `yaml:"type"`
	Template string        `yaml:"template"`
	Output   string        `yaml:"output"`
}
