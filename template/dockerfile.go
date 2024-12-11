package template

const Dockerfile = `FROM {{.BaseImage}}

WORKDIR /app

{{- if .Arg }}
ARG {{.Arg}}
{{- end }}

{{- if .Label }}
LABEL {{.Label}}
{{- end }}

{{- if .Maintainer }}
MAINTAINER {{.Maintainer}}
{{- end }}

COPY {{.Source}} /app

{{- if .Port }}
EXPOSE {{.Port}}
{{- end }}

{{- if .Volume }}
VOLUME {{.Volume}}
{{- end }}

{{- if .Environment }}
ENV {{.Environment}}
{{- end }}

{{- range .ShellCommands }}
RUN {{.}}
{{- end }}

CMD ["./{{.Executable}}", "{{.Arguments}}"]
`
