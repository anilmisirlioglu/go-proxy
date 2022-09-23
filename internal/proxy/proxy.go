package proxy

import (
	"fmt"
	"go-proxy/internal/module"
	"io"
	"net/http"
	"unicode/utf8"
	_ "unsafe"
)

type Proxy struct {
	Registry module.Registry
	Tpl      *module.Template
}

func New() *Proxy {
	return &Proxy{
		Registry: module.NewInMemoryRegistry(),
		Tpl:      module.NewTemplate(),
	}
}

func (p *Proxy) ListenAndServe() error {
	http.HandleFunc("/", p.rootHandler)
	return http.ListenAndServe(":8080", nil)
}

func (p *Proxy) rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := clean(r.URL.Path)
	switch path {
	case "favicon.ico":
		return
	case "":
		p.index(w, r)
	default:
		p.search(w, r, path)
	}
}

func (p *Proxy) index(w http.ResponseWriter, _ *http.Request) {
	str, err := p.Tpl.Index(p.Registry.Index())
	if err != nil {
		svrerror(err, w)
		return
	}

	_, _ = io.WriteString(w, str)
}

func (p *Proxy) search(w http.ResponseWriter, _ *http.Request, path string) {
	m := p.Registry.Search(path)
	if m == nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, fmt.Sprintf("module '%s' not found in module registry", path))
		return
	}

	str, err := p.Tpl.Search(m)
	if err != nil {
		svrerror(err, w)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_, _ = io.WriteString(w, str)
}

func svrerror(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = io.WriteString(w, err.Error())
}

func clean(p string) string {
	_, i := utf8.DecodeRuneInString(p)
	return p[i:]
}
