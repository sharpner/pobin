package pobin

import (
	"bytes"
	"io"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
)

type memoryTemplateLoader struct {
	loaderFunc    func(path string) ([]byte, error)
	allowFallback bool
}

//NewMemoryTemplateLoader loads a bindata template loader
func NewMemoryTemplateLoader(loaderFunc func(path string) ([]byte, error)) pongo2.TemplateLoader {
	return &memoryTemplateLoader{loaderFunc: loaderFunc, allowFallback: false}
}

func (m memoryTemplateLoader) Abs(base, name string) string {
	if !filepath.IsAbs(name) {
		if !strings.HasPrefix(name, "templates") {
			//TODO don't hardcode this prefix
			return "templates/sites/" + name
		}

		return name
	}

	return base
}

func (m memoryTemplateLoader) Get(path string) (io.Reader, error) {
	data, err := m.loaderFunc(path)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
