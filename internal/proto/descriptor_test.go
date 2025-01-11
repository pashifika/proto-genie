// Package protodesc
package proto_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pashifika/proto-genie/internal/proto"
)

func TestDescriptor(t *testing.T) {
	type fields struct {
		path string
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				input: "../../examples/service.proto",
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				input: "../../examples/service2.proto",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := proto.NewDescriptor()
			root, err := os.Getwd()
			if err != nil {
				t.Errorf("Getwd() error = %v", err)
			}
			tt.args.input = filepath.Join(root, tt.args.input)
			if err := d.LoadProto(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("LoadProto() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && len(d.Buf()) == 0 {
				t.Errorf("LoadProto() error = buf is empty")
			}
			_ = d.Close()
		})
	}
}
