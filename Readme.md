# Pobin a memory template loader

[![Build Status](https://travis-ci.org/sharpner/pobin.svg?branch=master)](https://travis-ci.org/sharpner/pobin)

Pobin is a Template Loader for [pongo2](https://github.com/flosch/pongo2) that is used to
load assets from [go-bindata](https://github.com/jteeuwen/go-bindata).

# Usage

It's easy to use pobin in combination with pongo2 and go-bindata:

1. First compile your assets
```go-bindata -o assets/templates.go -pkg=assets templates/...```
2. Initialize a pongo2 template set. Because of the way go-bindata works, you have to define
your own loading function as a parameter. 

```go
assetSet := pongo2.NewSet("assetfs", pobin.NewMemoryTemplateLoader(assets.Asset))
```

