package template

var list = []string{
	"go enum",
}

func List() []string {
	return list
}

var set = map[string]string{
	"go enum": GoEnum,
}

func Get(name string) string {
	return set[name]
}
