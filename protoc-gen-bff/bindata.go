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

var _bffTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x4d\x6f\xdb\x38\x13\x3e\x8b\xbf\x62\x5e\x1d\x0a\xa9\xb0\xe9\xf7\xb8\x48\xe1\xc3\x36\xed\xa6\x41\x37\x6d\x6a\xbb\xed\x21\x08\x16\xb4\x34\x92\xb8\x95\x49\x85\xa4\xec\xb8\x02\xff\xfb\x62\x28\xc9\x76\x92\x45\x82\x45\x6f\xe2\x7c\xcf\x33\x0f\x87\x9a\xcd\xe0\x5c\xe7\x08\x25\x2a\x34\xc2\x61\x0e\xeb\x3d\x34\x46\x3b\x9d\x4d\x4b\x54\xd3\x75\x51\x70\x78\xf7\x19\x3e\x7d\x5e\xc1\xfb\x77\x97\x2b\xce\x58\x23\xb2\x1f\xa2\x44\xe8\x3a\xbe\x12\xa6\x44\x77\xdd\x0b\xbc\x67\x4c\x6e\x1a\x6d\x1c\x24\x2c\x8a\x15\xba\x59\xe5\x5c\x13\xb3\x28\x96\x7a\x26\x75\xeb\x64\x4d\x07\x54\x99\xce\xa5\x2a\x67\x7f\x5b\xad\x48\x50\x6c\x5c\xcc\x58\x14\x77\x1d\xbf\xd0\x87\x60\xa4\x29\xa5\xab\xda\x35\xcf\xf4\x66\x96\x6b\xfd\x13\xcd\x34\xc7\x99\x41\xeb\x62\x16\x75\x9d\x11\xaa\x44\xe0\x97\x21\xa7\xf5\x9e\x45\x21\x46\x70\xed\xba\x29\xa0\xca\xbd\x67\x29\x63\x5b\x61\xa8\xa4\xb5\xb0\xf8\x75\x71\x09\x73\x20\xb3\xb7\xfd\x29\x58\xff\x05\x73\xe8\x0b\xe4\xef\xa4\xcd\x84\xc9\xc9\xcd\xed\x1b\x84\x2f\xcb\x6b\x61\xc4\x06\x1d\x1a\xb0\xce\xb4\x99\x83\x8e\x45\x1f\x71\x0f\x40\x67\xa9\x4a\x98\xcd\xe0\x7b\x25\xb3\x0a\x7e\xe0\x9e\x45\x7f\x48\xac\xf3\x13\x55\xa9\xd1\x82\xd3\xb0\x0b\x26\x45\xd0\x4a\x05\x02\x4a\x3d\x04\x64\xd1\x8a\x32\x9d\xf8\xd8\x4a\xb7\x75\x0e\x95\xd8\x22\xec\x2a\xe1\x80\x4a\x61\x9e\x31\x56\xb4\x2a\x83\xc4\xc2\xeb\x8b\x71\x60\x4b\x34\x5b\x99\x61\x0a\x17\xe8\x3e\x08\x95\xd7\x68\xec\x4a\x2f\xb0\x94\xd6\xa1\x49\x52\xb8\xb9\x25\xc4\xf8\x28\xa1\xf2\x0d\xba\xd6\xa8\xc7\x9a\x0e\x0e\x98\x5e\xa0\xbb\x12\x4d\x83\xf9\x15\xba\x4a\xe7\x3d\xb8\x1d\xf4\xa7\xb3\x80\xdf\xe2\xfd\x72\xd5\x9f\xbd\x8f\x27\x70\x2d\x5c\xd5\x2b\x3e\x08\xb3\xd1\x4a\xfe\xc4\x9c\x4c\x48\x1e\x0c\x86\xda\xce\xc0\xf2\xae\xe3\x9f\xc4\x06\xbd\x07\x3f\x39\x1d\x55\xe4\xa9\xc7\xe7\x5b\x1c\xc6\x96\xa4\x23\x5c\xc7\x76\x86\xf9\xbe\x10\x63\x89\xee\xbd\x31\xda\x0c\xf5\x24\x15\x04\x10\x4e\x65\x29\x20\x9d\x28\xb4\x2c\xa0\x82\xf9\x1c\x94\xac\xe9\x38\xa6\x2a\x36\x83\x47\x91\xc4\xa7\x9e\x90\x09\xa5\xb4\x83\x35\x92\x47\x9c\x52\x4b\x91\xe5\x78\x6a\x32\x87\xea\x50\xb2\x92\x35\x95\x1b\x98\xf6\xb8\xd4\x13\xba\x8d\x92\xae\xe3\xc3\x27\x31\x86\x10\x7b\x10\xf9\x49\x23\x07\x28\x3e\xe1\xee\x71\xf8\xc4\x3e\x09\x37\x81\xe7\xc3\xa5\x90\x3c\x01\x74\xf0\x49\x07\xb0\xec\xbf\x80\xa5\x64\x3d\x79\x80\xd8\xaa\x42\x28\xe5\x16\x15\x5c\x2c\xae\xcf\x61\xec\xee\x04\xbb\xb6\x1e\xc0\x3b\x20\xf5\xea\x71\x62\x4a\x30\x7c\x9e\x81\x9d\xb0\xe8\x01\x16\x67\x0f\x5a\x99\xb0\xc8\x4f\x46\xb0\x9f\x23\xf9\x33\xbc\x39\x90\x36\xd9\x01\xed\x35\xbe\x40\xdb\x68\x65\xf1\xbb\x91\x0e\xcd\x04\x0c\xbc\x1e\xe4\x77\x2d\x5a\x17\x10\xc9\xdc\x3d\x9c\xcd\xc1\xf0\x73\xad\x1c\xde\xbb\x24\xa5\x86\xee\x48\xd6\x75\xfc\x52\x35\xad\x23\xe8\xfb\xc0\x9d\xa7\xcb\x20\x0b\xc0\x3b\x88\x5f\xc7\x10\xae\xd8\x5b\x9d\xef\x61\x1a\xee\xdf\x5a\xe7\xfb\x00\x37\xb9\x0f\xeb\x6a\x81\x22\xff\xbd\xae\x13\xc3\xc9\x30\x65\x11\x0d\x81\x4c\xfe\x77\x1c\xc3\x23\x02\x26\xbb\x09\x98\x10\x87\xcc\x07\x7c\x59\x44\x4c\x8d\x72\x2c\x68\xf0\x21\x18\x3f\xaf\xb5\xc5\x24\x65\x3d\xb4\x30\x07\x5a\xd8\xfc\xab\xda\x08\x63\x2b\x51\x27\x7d\x3d\xaf\x0c\xde\xfd\x6a\xde\xe3\x0a\xa0\xef\x9d\x74\x55\xdf\xfc\x97\x16\xcd\x7e\x19\xee\x79\x40\x60\x2b\xea\x16\x6d\x8f\xe8\xd7\xc5\x9f\x3c\xe8\x09\xd3\xe3\x6b\x10\xec\x64\x01\xdb\x09\xe8\x1f\x64\xd9\xfb\xdc\xd0\x66\xfa\x88\x7b\xef\xe3\xdb\x37\xa4\x09\x05\xca\x02\xee\x47\xbb\xc0\xf6\xae\xbb\x40\x77\xae\xd5\x16\x8d\x43\x43\x63\x01\x7e\x85\x4e\xe4\xc2\x09\xef\x93\xed\xcd\xff\x6f\xd3\xa3\x7b\xd4\x75\x07\x2d\x1f\x19\x13\x0a\xa0\xf6\xee\x1e\x6a\x5d\xbf\x09\x61\x0e\xf7\x64\xe0\x07\xc4\xbb\x6e\xd8\x7d\xe3\xc7\x43\x04\xc8\xe9\x9b\x30\xfd\x02\x6e\xe8\x25\xb2\x87\x62\x43\x4c\x92\x24\x99\xbb\x7f\x0a\xc2\x7f\xee\xae\x0f\x7f\x43\x44\xf7\xfe\x96\x7f\x23\xe0\x5e\xee\x36\x3a\xb6\x3b\x85\x17\x1a\x7e\xd2\x6c\x64\xd0\x36\x07\x4e\xdb\x71\x1d\x1d\x5f\x08\x6a\xed\xc0\xb1\xa7\x14\x7b\x8e\x61\x23\xc1\x3c\x8b\xd6\x87\x14\x81\xc2\x57\x03\x81\x29\xf9\xaf\x84\xdd\xf1\x0f\x28\x72\x7a\x64\xf9\x12\x5d\x12\x87\x4b\xae\xdc\x94\xd6\x79\x3c\x81\x58\x34\x4d\x2d\x33\xe1\xa4\x56\xe1\x67\xe7\x0d\x64\x95\x30\x16\xdd\xbc\x75\xc5\xf4\xb7\x38\x2c\x03\xeb\xc8\x79\xe9\x84\x6b\x6d\xc8\xd4\x17\xb5\xe3\x61\xaf\x24\xeb\x94\x79\x36\xc2\xf5\x4f\x00\x00\x00\xff\xff\x57\x7b\x5f\x71\xae\x09\x00\x00")

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

	info := bindataFileInfo{name: "bff.tmpl", size: 2478, mode: os.FileMode(420), modTime: time.Unix(1518969704, 0)}
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

