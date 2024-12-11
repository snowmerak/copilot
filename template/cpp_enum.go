package template

const CppEnum = `#pragma once

enum class {{.EnumName}} : {{.EnumType}} {
{{- range .Values }}
	{{.Name}} = {{.Value}},
{{- end }}
};

inline std::string to_string({{.EnumName}} e) {
	switch (e) {
{{- range .Values }}
	case {{.EnumName}}::{{.Name}}:
		return "{{.Name}}";
{{- end }}
	default:
		return "Unknown";
	}
}

inline {{.EnumName}} from_string(const std::string& s) {
	if (s == "Unknown") {
		return {{.EnumName}}::Unknown;
	}
{{- range .Values }}
	if (s == "{{.Name}}") {
		return {{.EnumName}}::{{.Name}};
	}
{{- end }}
	return {{.EnumName}}::Unknown;
}

inline {{.EnumName}} from_index(int i) {
	switch (i) {
{{- range .Values }}
	case {{.Value}}:
		return {{.EnumName}}::{{.Name}};
{{- end }}
	default:
		return {{.EnumName}}::Unknown;
	}
}

inline int to_index({{.EnumName}} e) {
	switch (e) {
{{- range .Values }}
	case {{.EnumName}}::{{.Name}}:
		return {{.Value}};
{{- end }}
	default:
		return 0;
	}
}
`
