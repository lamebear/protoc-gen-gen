// Code generated by go-bindata.
// sources:
// service.template
// DO NOT EDIT!

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _serviceTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x4d\x6f\xe3\x36\x13\x3e\x8b\xbf\x62\x5e\xe3\x5d\x40\x5a\x38\xcc\x3d\x0b\x1f\x5a\x6f\xb6\xc8\x61\x93\x60\x93\x9e\x0b\x85\x1a\xa9\x82\x15\x52\x1d\xd2\x69\x02\x81\xff\xbd\xe0\x87\x3e\xac\xd8\x86\x13\x64\xdb\x5c\x22\x93\xf3\xf1\x70\xe6\x99\x67\xba\xee\x0c\xfe\xff\x90\x6b\xbc\x43\x7a\xaa\x05\xc2\xc5\x0a\xf8\x75\xfe\x88\x60\x2d\xf3\x97\x3a\x5c\xdc\xe6\x62\x93\x57\xe1\xfe\x96\x94\x51\xfd\xc1\xd4\x0e\xc9\xbb\x5e\xac\xa0\xa5\x5a\x9a\x12\x16\x9f\xf4\x9d\x3f\x5f\xec\x8f\xba\xdf\xbc\x16\xd8\xdb\x9f\x59\xcb\xce\xcf\xd7\xaa\x40\xa8\x50\x22\xe5\x06\x0b\x78\x78\x81\xd6\x41\x10\x67\x15\xca\xb3\x90\x98\xc3\xd7\x1b\xb8\xbe\xb9\x87\xcb\xaf\x57\xf7\x9c\xb5\x11\x5d\xd7\xf1\x08\xd4\x5a\xc6\xea\xc7\x56\x91\x81\x94\x01\x00\x2c\x84\x92\x06\x9f\xcd\x22\xfc\xaa\xd5\x82\x75\x1d\x50\x2e\x2b\x04\x7e\xe5\x2d\xb5\xc3\xeb\x6e\xbb\x8e\x47\xe4\x28\x0b\x17\xc9\xbb\x54\x4a\x55\x0d\xf2\x4a\x35\xb9\xac\xb8\xa2\xea\xbc\xa2\x56\x9c\x0b\x55\xa0\x5e\xb0\xe4\xd0\xbd\x36\xb9\xd9\xea\x05\xcb\x98\x8f\x58\x97\xc0\xbf\xd5\xa4\x8d\x8b\x6b\x5e\x5a\x04\x67\x75\x49\xa4\x08\x6a\x69\x90\xca\x5c\x20\x74\x3e\xa5\x2b\x44\x9a\x81\xcf\xc0\xdd\x0f\x66\x19\x2b\xb7\x52\x00\xa1\xd9\x92\xfc\xed\xc7\xed\xda\x7b\xa6\x48\x04\xe8\xbe\xb2\xf0\x0f\x3a\x96\xd4\x25\x54\x48\xb4\x04\xb5\x71\x45\x47\x22\x9e\x0e\xb9\xb2\x2f\xee\xb8\x63\x49\x12\x42\x41\x80\xc9\x43\x38\xe7\xc7\x43\xf6\xa5\x77\x0c\xc7\x59\xc6\x12\xcb\xd8\x5e\x97\x00\xf2\x77\xb9\x91\xea\x6f\x39\x73\xda\xa9\xa5\x7f\x73\xd7\x4d\x18\x64\x2d\x68\x43\x5b\x61\xe2\xb3\xf5\x93\xe8\x0d\x22\x69\xac\x65\xa3\xe7\x2e\x9d\xac\xdd\x29\x9b\x4b\x14\x9b\xfa\x1d\xcd\x9f\xaa\xd0\x3d\x09\xeb\x12\x72\x59\x00\x0f\x0c\xbd\x33\x84\xf9\x63\x2d\x2b\xe0\xeb\xa6\x46\x69\xc6\x83\x91\x04\x21\x75\xaa\xfd\xd5\x04\xd2\xc0\x31\xde\x75\xd3\x71\xb2\xf6\x8f\xc1\x2b\xa4\x89\xed\x08\xef\x6f\x34\xfa\xf6\xcf\x11\xbc\x4a\x28\xcc\x33\x44\xba\xf2\x75\xf8\xbf\x04\xc2\xbf\xe0\xb3\x8b\x33\x4e\x0f\xff\xa4\x17\xc0\xaf\x64\xbb\x35\xfc\x76\x53\xf5\x9f\xfd\x28\x65\x90\xee\xda\xaf\xb7\xa4\x95\x1b\xce\x9b\xad\x99\x1a\x2e\x23\x77\x76\x61\x1e\xa8\xcb\xaf\x58\xd5\xf2\x28\x56\x9f\x37\xde\xd3\xeb\xd8\x3f\xf1\xb9\x7b\x1d\xe2\x63\xbd\xc7\xd1\x87\x7b\x7e\xc6\x2f\xf0\x94\x1b\x15\xe2\xfe\xa5\xc5\xef\x68\xf2\x22\x37\xb9\x97\xa9\x7e\x96\xa7\xcd\x44\xef\xd6\xd3\xb4\x97\xc0\x50\xf5\x57\xc3\x7d\x8d\xcf\x26\xcd\xe0\x41\xa9\x26\x0c\xfb\x96\x08\xa5\x3b\xfa\xec\x9c\x1d\x5c\x4f\xb0\x51\x4a\x9d\xd5\x25\x51\xda\x93\x6a\x07\xec\xe4\xf3\xe0\x10\xbc\x69\x06\xbc\xce\xa4\x7a\x3e\xaa\xd9\x87\x0e\x46\xac\x85\x93\xaf\x8b\x15\x68\xae\x9f\x04\x9f\xc7\xcf\xbc\x49\x5d\x7a\xab\xff\xad\x40\xd6\x4d\x74\x73\x7f\x51\x8b\xf6\x08\x62\xf0\x8b\xea\x1d\xcd\x64\xdd\xf4\x65\x3b\x42\xf3\xd8\xc1\x9e\xc2\xfb\x5b\xf7\x56\x72\x86\x9e\x05\x59\x7f\x6c\x1b\x34\x4e\xda\xdf\xcf\x58\x1f\x28\x97\x02\x9b\x34\x1b\xd6\xc2\xbf\xd4\xae\x30\xd4\x63\xc7\x66\x92\x10\xf2\xf4\xa3\xec\xb7\xc6\xbc\x7b\xc3\xda\xd9\xdb\xb7\xc4\xb2\xa4\xc0\x12\x09\x88\x0f\x6f\x64\x49\x19\xf6\x5a\x22\xc6\xf4\x21\xd3\x0f\x14\x4f\x69\xc6\x92\x3e\xcd\x6a\x05\xb5\xe2\x97\x37\xdf\xbc\x79\xf2\x40\x98\x6f\x58\x92\xd8\xa1\xe9\x33\x2c\xc7\xc1\xf8\x95\x97\xf8\xb8\x40\xdc\xf7\x5e\x4c\x92\xbd\x25\x4e\xdc\x9e\xba\x1d\x5e\xe0\xd6\x6c\xcf\x86\x77\x94\x89\x45\x5c\xb1\x10\x77\x28\x8b\x5f\x64\xb1\x6e\x94\xc6\xd4\xe5\x79\x5f\xcc\xc3\xc3\x32\x13\x8e\xd3\x64\xe2\xed\x52\xbe\x84\x0f\xe3\xaa\x58\x1e\x57\x97\x91\xa6\x7e\xe7\x7c\x88\xd8\x38\xa2\x0a\x1e\x05\x7e\xf4\x56\x0e\x85\xe0\x83\xd0\x0f\x17\xbb\x6c\x76\x4d\x4c\xd5\x78\x7b\x08\xcc\x29\x80\x02\xa8\x09\xb4\x18\xcc\x03\xf1\xbb\xe4\xcb\x4f\x90\xd5\xd3\x68\xf1\xdf\xaf\xfd\x5e\xce\xa6\x03\x39\x67\x89\x30\xcf\xa7\x13\x43\xd6\xcd\xf2\xd4\x9a\x85\xac\x93\xca\xcd\xf7\xf8\x3f\x01\x00\x00\xff\xff\x37\x64\x8f\x70\xad\x0d\x00\x00")

func serviceTemplateBytes() ([]byte, error) {
	return bindataRead(
		_serviceTemplate,
		"service.template",
	)
}

func serviceTemplate() (*asset, error) {
	bytes, err := serviceTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service.template", size: 3501, mode: os.FileMode(420), modTime: time.Unix(1531772547, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"service.template": serviceTemplate,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"service.template": &bintree{serviceTemplate, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

