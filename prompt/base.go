package prompt

import "strings"

const Base = `Generate a code snippet from template and data.
The template is a single file, package or class that is used to generate code.
The data is a pseudo code or a inspiration that is used to generate code.
You only need to provide the template and data, the code will be generated automatically.
You only reply to the prompt with the generated code.
Do not explain the code and use code block, just provide the code.`

func Make(template string, query string) string {
	builder := strings.Builder{}
	builder.WriteString("# System\n\n")
	builder.WriteString(Base)
	builder.WriteString("\n\n# Template\n\n")
	builder.WriteString(template)
	builder.WriteString("\n\n# Data\n\n")
	builder.WriteString(query)
	return builder.String()
}
