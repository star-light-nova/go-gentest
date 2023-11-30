package templates

import (
	"text/template"
)

const SIMPLE_TEMPLATE = `package {{ .PackageName }}

import (
    "testing"
    "fmt"
)
{{ range .TV }}
func Test{{ .FuncTestName }}(t *testing.T) {
    got := {{ .FuncName }}()
    want := ""

    if got != want {
        t.Fatal("Test didn't pass")
    }
}
{{ end }}
`

func SimpleTemplate() (*template.Template, error) {
	templ, err := template.New("SimpleTemplate").Parse(SIMPLE_TEMPLATE)

	if err != nil {
		panic(err)
	}

	return templ, err
}
