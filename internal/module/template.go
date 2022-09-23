package module

import (
	"bytes"
	_ "embed"
	"fmt"
	"go-proxy/internal"
	"html/template"
	"net/url"
	"strings"
)

var (
	//go:embed tpl/module.tpl
	tplModuleIn string

	//go:embed tpl/index.tpl
	tplIndexIn string

	tplFuncs = map[string]any{
		"join":   strings.Join,
		"clean":  clean,
		"module": module,
	}
)

type Template struct {
	search *template.Template
	index  *template.Template
}

func NewTemplate() *Template {
	return &Template{
		search: template.Must(template.New("search").Funcs(tplFuncs).Parse(tplModuleIn)),
		index:  template.Must(template.New("index").Funcs(tplFuncs).Parse(tplIndexIn)),
	}
}

func (t *Template) Search(m *Module) (string, error) {
	return t.execute(t.search, map[string]interface{}{
		"Domain": internal.Domain,
		"Module": m,
	})
}

func (t *Template) Index(modules []*Module) (string, error) {
	return t.execute(t.index, modules)
}

func (t *Template) execute(tpl *template.Template, data any) (string, error) {
	var buf bytes.Buffer
	err := tpl.Execute(&buf, data)
	return buf.String(), err
}

func clean(s string) string {
	u, _ := url.Parse(s)
	return u.Host + u.Path
}

func module(m *Module) string {
	return fmt.Sprintf("%s/%s", internal.Domain, m.Name)
}