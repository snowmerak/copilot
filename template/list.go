package template

var list = []string{
	"go enum",
	"cpp enum",
}

func List() []string {
	return list
}

var set = map[string]string{
	"go enum":  GoEnum,
	"cpp enum": CppEnum,
}

func Get(name string) string {
	return set[name]
}
