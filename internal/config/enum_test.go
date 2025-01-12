// Package config
package config_test

import (
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/pashifika/proto-genie/internal/config"
)

func TestEnumProtoType_UnmarshalYAML(t *testing.T) {
	type args struct {
		rawYAML string
	}
	tests := []struct {
		name    string
		args    args
		want    config.EnumProtoType
		wantErr bool
	}{
		{
			name: "normal_message",
			args: args{
				rawYAML: "type: message",
			},
			want:    config.ProtoTypeMessage,
			wantErr: false,
		},
		{
			name: "normal_service",
			args: args{
				rawYAML: "type: service",
			},
			want:    config.ProtoTypeService,
			wantErr: false,
		},
		{
			name: "type error 1",
			args: args{
				rawYAML: "type: x",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "type error 2",
			args: args{
				rawYAML: "type: 123",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "empty error",
			args: args{
				rawYAML: `type: ""`,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type testType struct {
				Type config.EnumProtoType `yaml:"type"`
			}
			var tst testType
			if err := yaml.Unmarshal([]byte(tt.args.rawYAML), &tst); err != nil {
				if !tt.wantErr {
					t.Errorf("UnmarshalYAML() error = %v", err)
				}
				t.Logf("UnmarshalYAML() error = %v", err)
			}
			if !tt.wantErr && tst.Type != tt.want {
				t.Errorf("UnmarshalYAML() error = type[%s] error, wantErr %v", tst.Type, tt.wantErr)
			}
		})
	}
}
