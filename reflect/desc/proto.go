// Package desc
package desc

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	protoDesc "github.com/pashifika/proto-genie/internal/proto"
)

type Proto struct {
	fd []FDSetFile
}

func NewProto() *Proto {
	return &Proto{}
}

// Load is parses a proto file and cleans up resources.
func (p *Proto) Load(input string) error {
	d := protoDesc.NewDescriptor()
	if err := d.LoadProto(input); err != nil {
		return err
	}
	var fdSet descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(d.Buf(), &fdSet); err != nil {
		return err
	}
	p.phase(&fdSet)

	return d.Close()
}

func (p *Proto) phase(fdSet *descriptorpb.FileDescriptorSet) {
	files := fdSet.GetFile()
	p.fd = make([]FDSetFile, len(files))
	// Phase files
	for i, dp := range files {
		services := dp.GetService()
		messages := dp.GetMessageType()
		p.fd[i] = FDSetFile{
			Name:     dp.GetName(),
			Messages: make([]FDSetMessage, len(messages)),
			Services: make([]FDSetService, len(services)),
		}

		// Phase messages
		for j, msg := range messages {
			fields := msg.GetField()
			p.fd[i].Messages[j] = FDSetMessage{
				Name:   msg.GetName(),
				Fields: make([]FDSetField, len(fields)),
			}
			for k, field := range fields {
				p.fd[i].Messages[j].Fields[k] = FDSetField{
					Name:  field.GetName(),
					Type:  field.GetType().String(),
					Label: field.GetLabel().String(),
				}
			}
		}

		// Phase services
		for j, svc := range services {
			methods := svc.GetMethod()
			p.fd[i].Services[j] = FDSetService{
				Name:    svc.GetName(),
				Methods: make([]FDSetMethod, len(methods)),
			}
			for k, method := range methods {
				p.fd[i].Services[j].Methods[k] = FDSetMethod{
					Name:   method.GetName(),
					Input:  method.GetInputType(),
					Output: method.GetOutputType(),
				}
			}
		}
	}
}
