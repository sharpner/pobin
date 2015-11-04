package pobin

import (
	"bytes"
	"io"
	"path/filepath"

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
	if filepath.IsAbs(name) || base == "" {
		return name
	}

	if name == "" {
		return base
	}

	return filepath.Dir(base) + string(filepath.Separator) + name
}

func (m memoryTemplateLoader) Get(path string) (io.Reader, error) {
	data, err := m.loaderFunc(path)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
