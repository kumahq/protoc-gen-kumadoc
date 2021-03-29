package types

import pgs "github.com/lyft/protoc-gen-star"

type Message struct {
	Name        string
	Description string
	Fields      []*Field
}

func ParseMessage(message pgs.Message) *Message {
	var fields []*Field

	for _, field := range message.Fields() {
		fields = append(fields, ParseField(field))
	}

	description := TrimComments(message.SourceCodeInfo().LeadingComments())

	return &Message{
		Name:        message.Name().String(),
		Description: description,
		Fields:      fields,
	}
}
