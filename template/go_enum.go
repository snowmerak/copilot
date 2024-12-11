package template

const GoEnum = `package {{.EnumName}}

type {{.EnumName}} {{.EnumType}}

const (
{{- range .Values }}
	{{.Name}} {{.EnumName}} = {{.Value}}
{{- end }}
)

func (e {{.EnumName}}) String() string {
	switch e {
{{- range .Values }}
	case {{.Name}}:
		return "{{.Name}}"
{{- end }}
	default:
		return "Unknown"
	}
}

func FromString(s string) {{.EnumName}} {
	switch s {
{{- range .Values }}
	case "{{.Name}}":
		return {{.Name}}
{{- end }}
	default:
		return 0
	}
}

func FromIndex(i int) {{.EnumName}} {
	switch i {
{{- range .Values }}
	case {{.Value}}:
		return {{.Name}}
{{- end }}
	default:
		return 0
	}
}
`
