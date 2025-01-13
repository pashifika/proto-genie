// Package desc
package desc_test

import (
	"testing"

	"github.com/pashifika/proto-genie/reflect/desc"
)

func TestProto_Load(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := desc.NewProto()
			if err := p.Load(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
