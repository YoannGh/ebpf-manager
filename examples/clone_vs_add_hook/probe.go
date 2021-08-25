// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ebpf/bin/probe.o

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataProbeO = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x3d\x68\x14\x41\x18\x7d\x9b\xcb\xe5\xce\x24\xf8\x17\x90\xf3\xb8\x62\x40\x02\x09\xc2\xc6\x88\x04\x09\x08\x69\xd4\xc2\x14\x22\x16\x16\xc2\xb9\xb9\x6c\xcc\x71\xb7\xbb\xe7\xee\x1a\x8d\x1b\x50\x11\xfb\x28\xd8\x68\x25\x0a\xb6\xb1\x8a\xdd\x5d\x99\x32\x65\x4a\x4b\xb1\x8a\x36\xa6\x72\xe5\x9b\x9d\xdd\x9d\xcc\xed\x6a\xd2\x29\xf8\xc1\x65\xe6\xbd\xf9\xf9\x66\xde\xf7\x6e\x2e\x8f\x2f\xcf\x5f\x19\xd0\x34\xc4\xa1\xe1\x07\x52\x94\xc6\x5e\x29\xed\xcf\x89\xbf\xc7\xa0\xa1\x7b\x2a\xe2\x2a\x85\xfd\xf3\x83\xd9\xbd\x90\xda\x4d\x0d\x58\x1c\x06\x5a\xd5\x2f\x1c\x57\x34\xc0\x37\x6c\x9f\xfa\xec\x12\x1b\x0f\xaa\x9f\x13\x9e\x59\xab\x75\xea\x37\x1c\xdb\x0b\xaa\x3b\x09\xdf\x71\x9a\x36\xf5\xfd\x49\xb6\x16\x54\xb7\x13\x7e\xc9\x63\xcb\xd4\x77\x9c\x16\x0b\xaa\x5b\x09\x6f\xb5\x16\x9b\xd4\x77\xd9\xc4\x4a\x50\xed\x25\xe7\xa0\xf0\xaa\x5f\x39\xee\xbe\x8d\x70\x49\x03\x7a\x61\x18\x6e\x0e\x00\x67\x01\x3c\x07\x30\x24\xda\x32\x80\xbb\x74\x4e\x00\x0f\x44\xdb\x7d\x2f\xd6\x0d\x02\x7b\x61\x18\x76\x0d\x71\xff\x81\xfd\xf7\xef\x0a\x3d\x36\x8b\xd1\x3e\xb4\xdf\x69\xc2\x62\xfc\x55\x3c\xef\x2f\xd4\xcf\xb5\xfe\x1d\xfd\xe8\x5c\x93\x5c\x9f\xdd\xe4\xfc\x4e\x8b\x75\xf8\xbd\x9a\xb6\x1f\x48\xba\x79\x0d\xa3\x4d\xfd\x36\x5b\x76\x64\xdd\xfc\xfb\x2e\xd7\x87\x4d\x78\xab\xb2\x6e\xa9\x8f\x5c\x33\xd6\x2d\xd5\xe1\x7b\x9f\x0e\xdb\x42\x87\x33\x92\x0e\xea\x79\x07\xf1\x3f\x7e\x17\xa4\x39\xf9\x8e\xbc\x49\xbe\x24\xef\x92\x6f\xc9\xdb\xe4\x6b\xf2\x3e\xf9\x9e\xbe\x13\x14\xe4\xd5\xc3\xcc\x8f\xf6\x77\x4d\xaa\x39\xd5\x9b\x3c\xd1\x96\x97\x0e\xe3\xea\xf5\x79\xe0\x67\x18\xf2\xfa\x1e\x17\xe7\xd2\x1e\xdd\x40\x79\x6d\x44\x1b\x25\x6f\x88\x4f\x1c\x73\xd2\x83\x59\x03\x70\x21\x67\x2c\x1e\xbf\x29\xe1\x8d\x03\xea\xf2\x91\x7b\xe7\x5b\x98\x35\x56\x40\x21\x73\x4d\x01\xc5\x1c\xbe\xd4\xc7\xbd\x06\x70\x02\xc3\x09\x8e\x7d\xfa\x82\xf3\x23\x7d\xfc\x35\x00\x27\xa5\xbc\xf1\x3d\xc7\x39\x5f\xec\xe3\xef\x71\x3e\xcd\x1b\xdf\xbb\x24\x69\x2c\x47\x85\xe7\x2d\x27\xb8\x26\xf6\x89\x19\x82\x74\xda\xb2\x96\xe2\x91\x03\x8c\xd7\xa4\x71\x7a\x72\x2e\x4a\x98\x6e\x73\x4b\xc2\xfc\xae\xba\x6f\x3e\xf4\x21\xb9\x09\x96\xd1\xf1\xa6\x24\xc2\x43\xab\xe3\x3a\x0b\x66\x7d\x65\xc9\xab\x73\x3f\x42\x77\xcd\x76\x44\x4e\xa5\xa4\x34\x8b\xbb\x50\x9d\x15\x91\x2d\xd7\xf4\xa3\x89\x0a\x9e\x8a\x70\x7d\xc5\x74\xbd\xa6\x63\xa3\xde\x6e\x36\x4c\xdb\x33\xf9\x36\xba\xb9\x5c\x5f\x72\x0d\xcb\x84\x65\x34\x6d\xbd\x01\xdd\xf3\x5d\xdf\x58\x80\xee\xad\x5a\xbc\x75\x9d\x45\xc3\x37\x88\x9e\xd6\xa7\x67\x32\x14\x3f\x5c\x7c\x12\x7a\xa9\xc1\x44\x89\x77\x14\x5e\xfd\x9f\x42\x13\x9f\x21\x85\x9f\xcb\xc9\xa7\xbe\x9b\x77\xfe\xb0\x5e\xfd\xde\x95\x95\x79\xb7\x01\x1c\xc9\xc8\xb3\x21\x36\x64\x02\x1f\x15\xf7\x8c\xd7\xc7\x5e\x9d\xcd\xc9\xff\x44\x3b\x58\xfe\x99\x9c\xfc\xdb\x19\xf9\x8b\x19\xf9\x9f\xe6\xe4\xef\x89\xdf\x51\xf5\x5d\x51\xf3\x8f\x49\x9e\x97\xa3\x23\x88\x9a\x72\x7e\x55\xff\x2d\xb1\xfe\xbc\xc2\xaf\x8b\x89\xcf\x14\x3e\x79\x47\x45\xfb\x26\x27\xff\x58\x31\x3b\x9f\xea\x9f\x97\x39\xeb\x2b\x39\xeb\x55\xfc\x41\x7a\x03\xe4\x60\x62\x7d\x47\xe1\x55\xfd\xde\xe5\xd4\x6f\x57\x14\xe4\x9c\xc0\x54\xbf\xd1\x8c\xfa\xf5\x32\x72\x53\xac\x8b\xfc\x4c\x7a\xb7\xe4\xfa\xc7\xbf\x3b\xbf\x02\x00\x00\xff\xff\x80\xbf\xb3\x17\xb8\x0b\x00\x00")

func bindataProbeOBytes() ([]byte, error) {
	return bindataRead(
		_bindataProbeO,
		"/probe.o",
	)
}



func bindataProbeO() (*asset, error) {
	bytes, err := bindataProbeOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "/probe.o",
		size: 3000,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1629887216, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"/probe.o": bindataProbeO,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"": {Func: nil, Children: map[string]*bintree{
		"probe.o": {Func: bindataProbeO, Children: map[string]*bintree{}},
	}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
