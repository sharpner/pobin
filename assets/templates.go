package assets_test

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_sites_index_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8d\x41\x0a\xc2\x30\x10\x45\xf7\x9e\xe2\x53\x28\xe8\xc6\x5e\xa0\xb8\xf3\x1c\x32\x4d\x63\x09\x26\x99\xd0\x8c\x52\x09\xb9\xbb\x89\x41\x28\x2e\xe7\xff\x37\xef\xa7\x1e\x7a\x13\xed\xe7\x88\xce\xd2\x9b\x9f\x32\x4c\x14\xf5\x59\x82\xed\xd0\xe7\x43\xe9\x27\xcb\xea\x01\x31\x62\x75\x4d\x80\xeb\x46\x2e\x94\xc3\xf2\x62\x3c\x02\x2d\x1a\x95\x2b\x92\x86\xee\xdf\x14\xfb\x62\x97\xf6\x08\x94\xdc\xb8\xc0\xab\xa0\x73\xa4\x56\x8e\xc3\x9d\x57\x17\xdb\xdc\xd7\x77\xab\xc1\x0f\x1f\x67\xf3\xba\xa4\xb4\x6b\x8e\xa4\xc4\xb0\x3f\x21\xe7\x71\xa8\xed\xff\xf2\x27\x00\x00\xff\xff\x07\x3d\x65\xff\xd1\x00\x00\x00")

func templates_sites_index_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_sites_index_tpl,
		"templates/sites/index.tpl",
	)
}

func templates_sites_index_tpl() (*asset, error) {
	bytes, err := templates_sites_index_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/sites/index.tpl", size: 209, mode: os.FileMode(420), modTime: time.Unix(1446622326, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_sites_layout_base_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xc9\x28\xc9\xcd\xb1\xe3\xe2\xb4\xc9\x48\x4d\x4c\x01\xd2\x9c\x36\x25\x99\x25\x39\xa9\x76\xd5\xaa\x0a\x49\x39\xf9\xc9\xd9\x0a\x60\xae\x82\x6a\xad\x4b\x6a\x5a\x62\x69\x4e\x89\x42\x08\x88\x0f\x94\x4d\xcd\x4b\x81\x28\x50\xad\xb5\xd1\x87\xe8\x01\x9a\xa2\x0f\x35\xc6\x26\x29\x3f\xa5\x12\x61\x48\x72\x7e\x5e\x49\x6a\x5e\x09\x50\x2d\xba\x4e\xb0\x3a\x2e\xa0\x3e\xb0\x33\x00\x01\x00\x00\xff\xff\x3d\xbc\xce\x0d\x8e\x00\x00\x00")

func templates_sites_layout_base_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_sites_layout_base_tpl,
		"templates/sites/layout/base.tpl",
	)
}

func templates_sites_layout_base_tpl() (*asset, error) {
	bytes, err := templates_sites_layout_base_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/sites/layout/base.tpl", size: 142, mode: os.FileMode(420), modTime: time.Unix(1446621903, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_sites_macros_forms_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xcd\x4a\xc5\x30\x10\x85\xf7\xf7\x29\x86\x81\x0b\xba\xea\x0b\x24\x79\x01\x17\x0a\xea\x5a\x72\x9b\x51\x03\xcd\x0f\xc9\x14\x2b\xa5\xef\x6e\xda\xf4\xc7\x45\xb9\xb3\x9a\xe1\x9c\x33\x1f\x9c\xf1\x0a\x4e\xb7\x29\x40\x17\xbe\xac\xff\xf8\x0c\xc9\x3d\xe8\x96\x6d\xf0\x8f\x40\x43\x0c\x89\xe1\x3a\x5d\xa0\x8e\x98\x65\x70\xc4\xdf\xc1\x48\x7c\x79\x7e\x7d\x43\xa8\x66\x89\xe3\xb8\xae\x30\x4d\xa8\xb6\xc4\x36\xa2\xd3\x37\xea\xa0\xe4\x25\x2e\x24\x54\xef\x99\x92\xd7\x8e\x44\xb3\x68\x4a\x58\x1f\x7b\x06\x6b\x36\x07\xf0\x6f\x24\x89\x4c\x03\x23\xcc\xce\x5d\x68\xee\xfe\x8f\x3a\xe7\x9f\x90\x0c\xaa\x27\xf2\xbe\x6c\x7c\x82\xd8\x4d\x2b\xe5\xb8\x2b\xe9\xb8\x4f\x60\xf5\x4d\x0d\xe6\xfe\xe6\x2c\xff\xb7\x89\x66\xae\x49\x5d\x4a\xb5\xe4\x4d\x6d\xb7\x74\xf8\x17\x00\x00\xff\xff\x31\x54\x38\xca\x6a\x01\x00\x00")

func templates_sites_macros_forms_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_sites_macros_forms_tpl,
		"templates/sites/macros/forms.tpl",
	)
}

func templates_sites_macros_forms_tpl() (*asset, error) {
	bytes, err := templates_sites_macros_forms_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/sites/macros/forms.tpl", size: 362, mode: os.FileMode(420), modTime: time.Unix(1446622400, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/sites/index.tpl": templates_sites_index_tpl,
	"templates/sites/layout/base.tpl": templates_sites_layout_base_tpl,
	"templates/sites/macros/forms.tpl": templates_sites_macros_forms_tpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"sites": &_bintree_t{nil, map[string]*_bintree_t{
			"index.tpl": &_bintree_t{templates_sites_index_tpl, map[string]*_bintree_t{
			}},
			"layout": &_bintree_t{nil, map[string]*_bintree_t{
				"base.tpl": &_bintree_t{templates_sites_layout_base_tpl, map[string]*_bintree_t{
				}},
			}},
			"macros": &_bintree_t{nil, map[string]*_bintree_t{
				"forms.tpl": &_bintree_t{templates_sites_macros_forms_tpl, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

