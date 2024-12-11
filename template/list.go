package template

var list = []string{
	"go enum",
	"go union",
	"cpp enum",
	"dockerfile",
	"shell script",
}

func List() []string {
	return list
}

var set = map[string]string{
	"go enum":      GoEnum,
	"go union":     GoUnion,
	"cpp enum":     CppEnum,
	"dockerfile":   Dockerfile,
	"shell script": ShellScript,
}

func Get(name string) string {
	return set[name]
}
