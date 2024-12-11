package template

const ShellScript = `#!{{.Shell}}
{{- range .Commands }}
{{.}}
{{- end }}
`
