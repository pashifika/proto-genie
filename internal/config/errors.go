// Package config
package config

import "errors"

var (
	ErrProtoType      = errors.New("not supported this proto type")
	ErrConfigNotFound = errors.New("config not found")
)
