// Code generated by go-bindata.
// sources:
// bff.tmpl
// DO NOT EDIT!

package main

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

var _bffTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x4d\x6f\xdb\x38\x13\x3e\x8b\xbf\x62\x5e\x1d\x0a\xa9\xb0\xe9\xf7\xb8\x48\xe1\xc3\x36\xed\xa6\x41\x37\x6d\x6a\xa7\xed\x21\x08\x16\xb4\x34\x92\xb8\x95\x49\x85\xa4\xec\xb8\x02\xff\xfb\x62\x28\xc9\x76\xe2\x45\x82\x45\x6f\xe2\x7c\xcf\x33\x0f\x87\x9a\xcd\xe0\x5c\xe7\x08\x25\x2a\x34\xc2\x61\x0e\xab\x1d\x34\x46\x3b\x9d\x4d\x4b\x54\xd3\x55\x51\x70\x78\xf7\x19\x3e\x7d\xbe\x81\xf7\xef\x2e\x6f\x38\x63\x8d\xc8\x7e\x88\x12\xa1\xeb\xf8\x8d\x30\x25\xba\xeb\x5e\xe0\x3d\x63\x72\xdd\x68\xe3\x20\x61\x51\xac\xd0\xcd\x2a\xe7\x9a\x98\x45\xb1\xd4\x33\xa9\x5b\x27\x6b\x3a\xa0\xca\x74\x2e\x55\x39\xfb\xdb\x6a\x45\x82\x62\xed\x62\xc6\xa2\xb8\xeb\xf8\x85\xde\x07\x23\x4d\x29\x5d\xd5\xae\x78\xa6\xd7\xb3\x5c\xeb\x9f\x68\xa6\x39\xce\x0c\x5a\x17\xb3\xa8\xeb\x8c\x50\x25\x02\xbf\x0c\x39\xad\xf7\x2c\x0a\x31\x82\x6b\xd7\x4d\x01\x55\xee\x3d\x4b\x19\xdb\x08\x43\x25\xad\x84\xc5\xaf\x8b\x4b\x98\x03\x99\xbd\xed\x4f\xc1\xfa\x2f\x98\x43\x5f\x20\x7f\x27\x6d\x26\x4c\x4e\x6e\x6e\xd7\x20\x7c\x59\x5e\x0b\x23\xd6\xe8\xd0\x80\x75\xa6\xcd\x1c\x74\x2c\xfa\x88\x3b\x00\x3a\x4b\x55\xc2\x6c\x06\xdf\x2b\x99\x55\xf0\x03\x77\x2c\xfa\x43\x62\x9d\x1f\xa9\x4a\x8d\x16\x9c\x86\x6d\x30\x29\x82\x56\x2a\x10\x50\xea\x21\x20\x8b\x6e\x28\xd3\x91\x8f\xad\x74\x5b\xe7\x50\x89\x0d\xc2\xb6\x12\x0e\xa8\x14\xe6\x19\x63\x45\xab\x32\x48\x2c\xbc\xbe\x18\x07\xb6\x44\xb3\x91\x19\xa6\x70\x81\xee\x83\x50\x79\x8d\xc6\xde\xe8\x05\x96\xd2\x3a\x34\x49\x0a\xb7\x77\x84\x18\x1f\x25\x54\xbe\x41\xd7\x1a\xf5\x54\xd3\xc1\x1e\xd3\x0b\x74\x57\xa2\x69\x30\xbf\x42\x57\xe9\xbc\x07\xb7\x83\xfe\x74\x16\xf0\x5b\xbc\x5f\xde\xf4\x67\xef\xe3\x09\x5c\x0b\x57\xf5\x8a\x0f\xc2\xac\xb5\x92\x3f\x31\x27\x13\x92\x07\x83\xa1\xb6\x33\xb0\xbc\xeb\xf8\x27\xb1\x46\xef\xc1\x4f\x8e\x47\x15\x79\xea\xf1\xf9\x16\x87\xb1\x25\xe9\x08\xd7\xa1\x9d\x61\xbe\x2f\xc4\x58\xa2\x7b\x6f\x8c\x36\x43\x3d\x49\x05\x01\x84\x63\x59\x0a\x48\x27\x0a\x2d\x0b\xa8\x60\x3e\x07\x25\x6b\x3a\x8e\xa9\x8a\xf5\xe0\x51\x24\xf1\xb1\x27\x64\x42\x29\xed\x60\x85\xe4\x11\xa7\xd4\x52\x64\x39\x1e\x9b\xcc\xa1\xda\x97\xac\x64\x4d\xe5\x06\xa6\x3d\x2d\xf5\x88\x6e\xa3\xa4\xeb\xf8\xf0\x49\x8c\x21\xc4\x1e\x45\x3e\x69\x64\x0f\xc5\x27\xdc\x3e\x0d\x9f\xd8\x93\x70\x13\x78\x3e\x5c\x0a\xc9\x09\xa0\x83\x4f\x3a\x80\x65\xff\x05\x2c\x25\xeb\xc9\x23\xc4\x5c\x85\x50\xca\x0d\x2a\xb8\x58\x5c\x9f\xc3\xd8\xdd\x11\x76\x6d\x3d\x80\xb7\x47\xea\xd5\xd3\xc4\x94\x60\xf8\x3c\x03\x3b\x61\xd1\x23\x2c\xce\x1e\xb5\x32\x61\x91\x9f\x8c\x60\x3f\x47\xf2\x67\x78\xb3\x27\x6d\xb2\x05\xda\x6b\x7c\x81\xb6\xd1\xca\xe2\x77\x23\x1d\x9a\x09\x18\x78\x3d\xc8\xef\x5b\xb4\x2e\x20\x92\xb9\x07\x38\x9b\x83\xe1\xe7\x5a\x39\x7c\x70\x49\x4a\x0d\xdd\x93\xac\xeb\xf8\xa5\x6a\x5a\x47\xd0\xf7\x81\x3b\x4f\x97\x41\x16\x80\xf7\x10\xbf\x8e\x21\x5c\xb1\xb7\x3a\xdf\xc1\x34\xdc\xbf\x95\xce\x77\x01\x6e\x72\x1f\xd6\xd5\x02\x45\xfe\x7b\x5d\x27\x86\x93\x61\xca\x22\x1a\x02\x99\xfc\xef\x30\x86\x27\x04\x4c\xb6\x13\x30\x21\x0e\x99\x0f\xf8\xb2\x88\x98\x1a\xe5\x58\xd0\xe0\x43\x30\x7e\x5e\x6b\x8b\x49\xca\x7a\x68\x61\x0e\xb4\xb0\xf9\x57\xb5\x16\xc6\x56\xa2\x4e\xfa\x7a\x5e\x19\xbc\xff\xd5\xbc\x87\x15\x40\xdf\x5b\xe9\xaa\xbe\xf9\x2f\x2d\x9a\xdd\x32\xdc\xf3\x80\xc0\x46\xd4\x2d\xda\x1e\xd1\xaf\x8b\x3f\x79\xd0\x13\xa6\x87\xd7\x20\xd8\xc9\x02\x36\x13\xd0\x3f\xc8\xb2\xf7\xb9\xa5\xcd\xf4\x11\x77\xde\xc7\x77\x6f\x48\x13\x0a\x94\x05\x3c\x8c\x76\x81\xed\x5d\x77\x81\xee\x5c\xab\x0d\x1a\x87\x86\xc6\x02\xfc\x0a\x9d\xc8\x85\x13\xde\x27\x9b\xdb\xff\xdf\xa5\x07\xf7\xa8\xeb\xf6\x5a\x3e\x32\x26\x14\x40\xed\xdd\x3f\xd6\xba\x7e\x13\xc2\x1c\x1e\xc8\xc0\x0f\x88\x77\xdd\xb0\xfb\xc6\x8f\xc7\x08\x90\xd3\x37\x61\xfa\x05\xdc\xd0\x4b\x64\xf7\xc5\x86\x98\x24\x49\x32\xf7\x70\x0a\xc2\x7f\xee\xae\x0f\x7f\x4b\x44\xf7\xfe\x8e\x7f\x23\xe0\x5e\xee\x36\x3a\xb4\x3b\x85\x17\x1a\x3e\x69\x36\x32\x68\x9b\x3d\xa7\xed\xb8\x8e\x0e\x2f\x04\xb5\xb6\xe7\xd8\x29\xc5\x9e\x63\xd8\x48\x30\xcf\xa2\xd5\x3e\x45\xa0\xf0\xd5\x40\x60\x4a\xfe\x2b\x61\xb7\xfc\x03\x8a\x9c\x1e\x59\xbe\x44\x97\xc4\xe1\x92\x2b\x37\xa5\x75\x1e\x4f\x20\x16\x4d\x53\xcb\x4c\x38\xa9\x55\xf8\xd9\x79\x03\x59\x25\x8c\x45\x37\x6f\x5d\x31\xfd\x2d\x0e\xcb\xc0\x3a\x72\x5e\x3a\xe1\x5a\x1b\x32\xf5\x45\x6d\x79\xd8\x2b\xc9\x2a\x65\x9e\x8d\x70\xfd\x13\x00\x00\xff\xff\xb0\x76\xd8\xc7\xae\x09\x00\x00")

func bffTmplBytes() ([]byte, error) {
	return bindataRead(
		_bffTmpl,
		"bff.tmpl",
	)
}

func bffTmpl() (*asset, error) {
	bytes, err := bffTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bff.tmpl", size: 2478, mode: os.FileMode(420), modTime: time.Unix(1561111682, 0)}
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
	"bff.tmpl": bffTmpl,
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
	"bff.tmpl": &bintree{bffTmpl, map[string]*bintree{}},
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
