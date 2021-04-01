package types

import (
	docs "github.com/kumahq/protoc-gen-kumadoc/proto/generated"

	pgs "github.com/lyft/protoc-gen-star"
)

type Field struct {
	Name             string
	Description      string
	HideDescription  bool
	ShortDescription string
	ProtoType        string
	Embed            *Message
	IsEmbed          bool
	IsRequired       bool
	IsEnum           bool
	Enum             []string
	IsRepeated       bool
	Package          string
	PolicyPackage    string
}

func ParseField(policyPackage string, f pgs.Field) *Field {
	var required bool

	if _, err := f.Extension(docs.E_Required, &required); err != nil {
		panic(err)
	}

	description := TrimComments(f.SourceCodeInfo().LeadingComments())

	typ := f.Type()

	field := &Field{
		Name:          f.Name().String(),
		Description:   description,
		ProtoType:     typ.ProtoType().String(),
		IsRequired:    required,
		IsRepeated:    typ.IsRepeated(),
		Package:       f.Package().ProtoName().String(),
		PolicyPackage: policyPackage,
	}

	if typ.IsEmbed() {
		field.IsEmbed = true
		field.Embed = ParseMessage(policyPackage, typ.Embed())
		field.Package = typ.Embed().Package().ProtoName().String()
	}

	if typ.IsEnum() {
		field.IsEnum = true

		for _, value := range typ.Enum().Values() {
			field.Enum = append(field.Enum, value.Name().String())
		}
	}

	if typ.IsRepeated() && typ.Element().IsEnum() {
		field.IsEnum = true

		for _, value := range typ.Element().Enum().Values() {
			field.Enum = append(field.Enum, value.Name().String())
		}
	}

	return field
}
