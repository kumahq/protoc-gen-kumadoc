package types

import (
	doc "github.com/kumahq/protoc-gen-kumadoc/proto"
	pgs "github.com/lyft/protoc-gen-star"
)

type Message struct {
	Name          string
	Description   string
	Fields        []*Field
	IsHidden      bool
	PolicyPackage string
}

func ParseMessage(policyPackage string, message pgs.Message) *Message {
	var fields []*Field

	for _, field := range message.Fields() {
		fields = append(fields, ParseField(policyPackage, field))
	}

	description := TrimComments(message.SourceCodeInfo().LeadingComments())

	var isHidden bool
	if _, err := message.Extension(doc.E_Hide, &isHidden); err != nil {
		panic(err)
	}

	return &Message{
		Name:          message.Name().String(),
		Description:   description,
		Fields:        fields,
		IsHidden:      isHidden,
		PolicyPackage: policyPackage,
	}
}
