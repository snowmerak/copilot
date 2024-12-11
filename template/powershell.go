package template

const PowerShell = `{{- range .Commands }}
{{.}}
{{- end }}
`
