# {{ .Name }}

{{ define "field" }}
- `{{ .Name | camelcase | untitle }}`
{{- if .IsRequired }} (required)
{{- else }} (optional){{ end }}
{{- if .Description }}

{{ .Description | indent 2 }}
{{- end -}}
{{- if .Embed }}
  {{ range .Embed.Fields }}
    {{- include "field" . | indent 2 -}}
  {{- end -}}
{{- end -}}
{{- if .IsEnum }}

  Supported values:
  {{- range .Enum }}

  - `{{ . }}`
  {{- end }}
{{- end -}}
{{ end -}}

{{ define "message" }}
## {{ .Name -}}
{{ if .Description }}

{{ .Description }}
{{ end -}}

{{ range .Fields -}}
  {{ include "field" . }}
{{ end }}
{{ end -}}

{{- if .Description -}}
    {{- .Description -}}
{{- end }}

{{- range .Messages }}
  {{ template "message" . }}
{{- end -}}
