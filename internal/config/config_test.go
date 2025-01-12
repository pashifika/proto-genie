// Package config
package config_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/hashicorp/go-version"

	"github.com/pashifika/proto-genie/internal/config"
)

func TestLoadConfig(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		t.Errorf("Getwd() error = %v", err)
	}

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				path: "../../examples/generate.yml",
			},
			want: &config.Config{
				Version: *version.Must(version.NewVersion("1.0.0")),
				Settings: config.Setting{
					TemplateRoot: "./xx",
				},
				Generators: []config.Generator{
					{
						Name:     "msg_usecase",
						Type:     "message",
						Template: "xx.tmpl",
						Output:   "./xx/{{snake_name}}.go",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				path: "../../examples/generate2.yml",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.path = filepath.Join(root, tt.args.path)
			got, err := config.LoadConfig(tt.args.path)
			if err != nil {
				if tt.wantErr {
					t.Logf("LoadConfig() error = %v", err)
				} else {
					t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
