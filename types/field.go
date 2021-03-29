package types

import (
	docs "github.com/kumahq/protoc-gen-kumadoc/proto/generated"

	pgs "github.com/lyft/protoc-gen-star"
)

type Field struct {
	Name        string
	Description string
	Typ         string
	Embed       *Message
	IsEmbed     bool
	IsRequired  bool
	IsEnum      bool
	Enum        []string
}

func ParseField(f pgs.Field) *Field {
	var required bool

	if _, err := f.Extension(docs.E_Required, &required); err != nil {
		panic(err)
	}

	description := TrimComments(f.SourceCodeInfo().LeadingComments())
	typ := f.Type()

	field := &Field{
		Name:        f.Name().String(),
		Description: description,
		Typ:         typ.ProtoType().String(),
		IsRequired:  required,
	}

	if typ.IsEmbed() {
		field.IsEmbed = true
		field.Embed = ParseMessage(typ.Embed())
	}

	if typ.IsEnum() {
		field.IsEnum = true

		for _, value := range typ.Enum().Values() {
			field.Enum = append(field.Enum, value.Name().String())
		}
	}

	return field
}
