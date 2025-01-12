// Package config
package config

import "gopkg.in/yaml.v3"

type EnumProtoType string

const (
	ProtoTypeMessage EnumProtoType = "message"
	ProtoTypeService EnumProtoType = "service"
)

func (e *EnumProtoType) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode || value.IsZero() {
		return ErrProtoType
	}

	// check user input
	*e = EnumProtoType(value.Value)
	switch *e {
	case ProtoTypeMessage, ProtoTypeService:
		break
	default:
		return ErrProtoType
	}

	return nil
}
