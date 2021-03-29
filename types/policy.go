package types

import (
	doc "github.com/kumahq/protoc-gen-kumadoc/proto/generated"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type Policy struct {
	Name        string
	Description string
	Messages    []*Message
}

func ParsePolicy(ctx pgsgo.Context, ext *doc.Config, f pgs.File) *Policy {
	var name string
	if name = ext.GetName(); name == "" {
		name = ctx.Name(f).UpperCamelCase().String()
	}

	info := f.SourceCodeInfo()

	var messages []*Message
	for _, message := range f.Messages() {
		messages = append(messages, ParseMessage(message))
	}

	return &Policy{
		Name:        name,
		Description: TrimComments(info.LeadingComments()),
		Messages:    messages,
	}
}
