package pobin

import (
	"bytes"
	"io"
	"path/filepath"

	"github.com/flosch/pongo2"
)

type memoryTemplateLoader struct {
	loaderFunc     func(path string) ([]byte, error)
	fallbackLoader *pongo2.TemplateLoader
}

//NewMemoryTemplateLoader loads a bindata template loader
func NewMemoryTemplateLoader(loaderFunc func(path string) ([]byte, error)) pongo2.TemplateLoader {
	return &memoryTemplateLoader{loaderFunc: loaderFunc}
}

//NewMemoryTemplateLoaderWithFallback loads a bindata template loader, but if the file is not found
//it will try to find it in the fallback loader with the same path
func NewMemoryTemplateLoaderWithFallback(loaderFunc func(path string) ([]byte, error), fallbackLoader pongo2.TemplateLoader) pongo2.TemplateLoader {
	return &memoryTemplateLoader{loaderFunc: loaderFunc, fallbackLoader: &fallbackLoader}
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
		if m.fallbackLoader != nil {
			return (*m.fallbackLoader).Get(path)
		}

		return nil, err
	}

	return bytes.NewReader(data), nil
}
