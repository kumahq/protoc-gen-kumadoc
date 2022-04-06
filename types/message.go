package types

import (
	pgs "github.com/lyft/protoc-gen-star"

	doc "github.com/kumahq/protoc-gen-kumadoc/proto"
)

type Message struct {
	Name             string
	Description      string
	Fields           []*Field
	IsHidden         bool
	ComponentPackage string
}

func ParseMessage(componentPackage string, message pgs.Message) *Message {
	var fields []*Field

	for _, field := range message.Fields() {
		fields = append(fields, ParseField(componentPackage, field))
	}

	description := TrimComments(message.SourceCodeInfo().LeadingComments())

	var isHidden bool
	if _, err := message.Extension(doc.E_Hide, &isHidden); err != nil {
		panic(err)
	}

	return &Message{
		Name:             message.Name().String(),
		Description:      description,
		Fields:           fields,
		IsHidden:         isHidden,
		ComponentPackage: componentPackage,
	}
}
