package types

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	doc "github.com/kumahq/protoc-gen-kumadoc/proto"
)

type Component struct {
	Name     string
	Header   string
	Messages []*Message
	FileName string
	Package  string
}

func ParseComponent(ctx pgsgo.Context, ext *doc.Config, f pgs.File) *Component {
	var name string
	if name = ext.GetName(); name == "" {
		name = ctx.Name(f).UpperCamelCase().String()
	}

	var fileName string
	if fileName = ext.GetFileName(); fileName == "" {
		fileName = ctx.Name(f).String()
	}

	info := f.SourceCodeInfo()
	componentPackage := f.Package().ProtoName().String()

	var messages []*Message
	for _, message := range f.Messages() {
		messages = append(messages, ParseMessage(componentPackage, message))
	}

	return &Component{
		Name:     name,
		Header:   TrimComments(info.LeadingComments()),
		Messages: messages,
		FileName: fileName + ".md",
		Package:  componentPackage,
	}
}
