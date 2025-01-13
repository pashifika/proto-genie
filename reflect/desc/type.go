// Package desc
package desc

type FDSetFile struct {
	Name     string
	Messages []FDSetMessage
	Services []FDSetService
}

type FDSetMessage struct {
	Name   string
	Fields []FDSetField
}

type FDSetField struct {
	Name  string
	Type  string
	Label string
}

type FDSetService struct {
	Name    string
	Methods []FDSetMethod
}

type FDSetMethod struct {
	Name   string
	Input  string
	Output string
}
