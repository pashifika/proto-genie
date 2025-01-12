// Package proto
package proto

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pashifika/util/mem"
)

type Descriptor struct {
	buf  *mem.FakeIO
	path string
}

func NewDescriptor() *Descriptor {
	return &Descriptor{}
}

func (d *Descriptor) create(input string) error {
	d.path = strings.TrimSuffix(input, filepath.Ext(input)) + ".pb"
	cmd := exec.Command("protoc",
		"--descriptor_set_out="+filepath.Base(d.path),
		"--include_imports",
		filepath.Base(input),
	)
	cmd.Dir = filepath.Dir(input)

	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// LoadProto parses a given proto file, generates a descriptor_set's content into memory.
func (d *Descriptor) LoadProto(input string) error {
	if err := d.create(input); err != nil {
		return err
	}
	buf, err := os.ReadFile(d.path)
	if err != nil {
		return err
	}
	d.buf = mem.NewFakeIO(buf)

	if err := os.Remove(d.path); err != nil {
		return err
	}

	d.path = input
	return nil
}

func (d *Descriptor) Buf() []byte {
	return d.buf.Bytes()
}

func (d *Descriptor) Close() error {
	if d.buf == nil {
		return nil
	}
	return d.buf.Close()
}
