package main

import (
	"bytes"
	"embed"
	"fmt"
	"strings"
	"text/template"

	doc "github.com/kumahq/protoc-gen-kumadoc/proto"
	"github.com/kumahq/protoc-gen-kumadoc/types"

	"github.com/Masterminds/sprig/v3"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type Module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl map[doc.Config_Type]*template.Template
}

func New() pgs.Module {
	return &Module{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (m *Module) Name() string {
	return "KumaDoc"
}

func include(t *template.Template) func(name string, data interface{}) (string, error) {
	return func(name string, data interface{}) (string, error) {
		buf := bytes.NewBuffer(nil)
		if err := t.ExecuteTemplate(buf, name, data); err != nil {
			return "", err
		}
		return buf.String(), nil
	}
}

//go:embed templates
var fs embed.FS

func (m *Module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())

	t := template.New("templates")

	var funcMap template.FuncMap = map[string]interface{}{
		"include": include(t),
	}

	t = template.Must(t.Funcs(sprig.TxtFuncMap()).Funcs(funcMap).ParseFS(fs, "templates/*.tpl"))

	m.tpl = map[doc.Config_Type]*template.Template{
		doc.Config_Policy: t.Lookup("generic.md.tpl"),
		doc.Config_Proxy:  t.Lookup("generic.md.tpl"),
		doc.Config_Other:  t.Lookup("generic.md.tpl"),
	}
}

func (m *Module) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		var docExt doc.Config
		ok, err := f.Extension(doc.E_Config, &docExt)
		m.CheckErr(err, "unable to read parse extension from file")

		if ok {
			if tpl, ok := m.tpl[docExt.Type]; ok {
				component := types.ParseComponent(m.ctx, &docExt, f)
				filename := fmt.Sprintf(
					"%s_%s",
					strings.ToLower(docExt.Type.String()),
					component.FileName,
				)

				m.AddGeneratorTemplateFile(filename, tpl, component)
			}
		}
	}

	return m.Artifacts()
}
