package template

const GoUnion = `package {{.UnionName}}

type {{.UnionName}} any

{{- range .Values }}
type {{.Name}} struct {
	Value {{.Type}}
}

func New{{.Name}}(value {{.Type}}) {{.UnionName}} {
	return {{.UnionName}}({{.Name}}{Value: value})
}

func Is{{.Name}}(u {{.UnionName}}) bool {
	_, ok := u.({{.Name}})
	return ok
}

func As{{.Name}}(u {{.UnionName}}) ({{.Name}}, bool) {
	v, ok := u.({{.Name}})
	return v, ok
}
{{- end }}
`
