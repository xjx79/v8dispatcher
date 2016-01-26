// Code generated by go-bindata.
// sources:
// js/console.js
// js/registry.js
// js/setup.js
// DO NOT EDIT!

package v8dispatcher

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

var _jsConsoleJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\xc1\x6e\xdb\x30\x0c\x3d\xcf\x5f\x41\x18\x03\x6a\x27\x85\xbd\x73\xb3\x9e\x86\x62\xd8\x30\x74\x87\xee\x56\xf4\xa0\xd8\x94\xcc\xc5\x96\x02\x89\xea\x56\x14\xf9\xf7\x91\x4e\x03\xbb\xd7\xe9\x60\xc2\xef\x89\xef\xf1\xd1\x6e\x37\x05\x6c\xe0\xd7\x40\x09\x2c\x8d\x08\x52\x53\xde\xff\xc6\x8e\x81\x03\xf0\x80\xc0\x18\xa7\x04\xc6\xf7\xd0\x05\xdf\x13\x53\xf0\x09\x7a\xb4\xe4\xb1\x07\xf2\xda\x3e\x77\x5e\xfd\xf8\xf6\xe5\xee\xfe\xe1\xae\xe1\xbf\x7c\x75\x0d\x7f\x06\xea\x06\x95\x3b\x9a\xc8\x10\xac\x68\xa9\x76\xc8\xb1\x43\x51\xea\x51\x88\xee\x60\x1c\x36\xa2\xa0\x22\x26\xf3\x10\xe2\x0d\xe0\x44\xdd\x61\x44\x12\xac\x2d\xda\x56\x5d\x53\x38\x4f\x96\x93\x58\xda\x10\xc1\x21\x33\x79\x07\x63\x70\x80\x9e\x23\x61\x92\x51\xc0\x28\xe0\x30\x42\xf0\xf3\xe8\x5f\x03\x24\xea\xc5\xa1\x6d\x8b\x8b\xcc\x2d\xbc\x9e\x76\x97\xb7\xe6\x18\xc9\xb3\x60\x36\xfb\x4e\x93\x55\x35\xbc\x16\x20\xe7\xd9\x44\x98\x92\x13\xaa\x2c\x77\x33\xa2\xbe\x95\xc2\x24\xe0\xa7\x9d\x94\xcf\x60\xa2\xcb\x93\x0c\x90\x9a\x11\xbd\xe3\x41\xd0\xed\xf6\x22\xa1\x47\x25\xb6\xb7\xcb\xbd\x47\x7a\x82\x2d\x94\x50\x95\x52\xf8\xe5\x88\xc1\x56\x6b\xb2\x56\xb6\x86\x37\xcb\xd3\xfc\xfc\x38\x0f\x59\x89\x54\x5d\x9c\x0a\x5d\x89\xc6\x66\x73\x90\xd0\x46\x07\x25\xb3\x97\x60\x3e\x4f\x7b\x8d\x6e\x17\xb7\x55\xee\x46\x7b\x56\x39\xa1\xdd\x2c\xf7\x64\xd3\xb0\x0e\x2e\x44\x92\xcb\x8f\x4f\xe7\x31\xc4\xd1\x8e\x86\x19\x65\xc3\xe3\xb8\xb4\x7d\xf8\xff\xbd\xa8\x43\x73\xcc\x69\x78\x9f\xfe\x5d\xec\x84\xbe\xaf\xbe\x3f\xfc\xbc\x6f\x92\x7c\x61\xef\xc8\xbe\x54\x8b\x42\x19\xb1\x43\x7a\xc6\x58\xde\x40\xf9\x16\xb2\xbc\x5e\xe8\x84\xa3\xfc\xc4\x61\xa6\x25\xfb\x9a\x52\x73\x81\xb5\x9c\xfd\x6a\x31\x3e\x15\xff\x02\x00\x00\xff\xff\x5e\x37\xa5\xd2\x0b\x03\x00\x00")

func jsConsoleJsBytes() ([]byte, error) {
	return bindataRead(
		_jsConsoleJs,
		"js/console.js",
	)
}

func jsConsoleJs() (*asset, error) {
	bytes, err := jsConsoleJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/console.js", size: 779, mode: os.FileMode(420), modTime: time.Unix(1453531099, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRegistryJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\x4d\x6f\xda\x40\x10\xbd\xfb\x57\x8c\x72\x28\x06\x81\x0d\x4a\x1a\x45\x20\xd4\x43\x9b\x43\xa5\xb6\x97\x54\xbd\x44\x51\xb5\xac\xc7\xf6\x06\x7b\xd7\xdd\x1d\x13\x10\xf0\xdf\x3b\xbb\xb1\xa9\xa5\x34\x7b\x18\xc3\x7c\xbc\x99\xf7\x66\xd2\x49\x04\x13\xf8\x59\x2a\x07\xb9\xaa\x10\xf8\xeb\xda\xcd\x33\x4a\x02\x32\x40\x25\x02\xa1\xad\x1d\x08\x9d\x81\x34\x3a\x53\xa4\x8c\x76\x90\x61\xae\x34\x66\xa0\xb4\x2f\x0f\x95\xa3\x6f\x5f\x3f\xdf\xff\x78\xb8\x4f\x68\x4f\xa3\x29\xbc\x94\x4a\x96\x1e\xae\x11\x96\xc0\xe4\x8c\xe5\xb1\x4d\x6b\x25\x32\x52\x86\x1c\x90\x5b\x51\x60\xc2\x08\x1e\x44\xb4\x54\x1a\xbb\x04\xac\x95\xdc\x56\xa8\xd8\x97\x46\xd1\x4e\x58\xf8\x75\xf7\x05\xd6\xc1\x9e\x4e\x70\xbc\x2a\x2a\xb3\x11\x95\xbb\x5a\x1e\xcf\xe7\x55\x14\xa5\x29\x94\x44\xcd\x32\x4d\x1d\x31\xa2\xd9\xa1\xcd\x2b\xf3\x92\x48\x53\xa7\x7f\x5a\x74\x61\xe0\x74\x31\xff\x38\xbf\xbe\x49\xa5\x45\x41\x38\x2b\x5a\x95\xcd\x5a\x6f\x94\x9e\x3d\x8b\x9d\x70\xd2\xaa\x86\x22\xee\x91\x78\x37\xb7\xcb\x5b\x2d\x7d\x69\x3c\x86\x63\x04\xfc\x2c\x52\x6b\x35\x8c\xf6\xdd\x9b\x05\x73\xe3\xcd\xa1\xff\xdb\xbf\x51\x62\xb1\xa9\x84\xc4\x38\x7d\xdc\x1f\x9e\xd2\x62\xfa\x0f\x4f\xf6\x80\xfe\x79\x7a\x96\xbb\x7d\x17\x54\x26\x96\x45\x36\x35\x37\x9c\xc0\xe2\x16\x4e\x30\x9f\x5e\xf2\x42\x2e\xe7\x49\x58\xaf\x79\x84\x11\x7c\xe2\xb2\x25\xc4\x16\x3e\xc0\x7c\x7f\xed\x93\xf7\x77\xe3\xd5\x25\xbf\x1b\x76\x97\x90\x79\x20\xab\x74\x11\x2f\x6e\xbb\xf0\x99\xbf\xe7\x20\x5b\x3f\xd2\x6f\x8b\x85\x72\x64\x0f\xb0\x45\x6c\x1c\xa8\x0c\x35\xa9\xfc\x20\x36\xbc\xd6\x78\x73\x80\x02\x35\x5a\xd6\x8d\xf7\x9d\x8d\x2f\x65\x8e\x31\x82\x62\x6f\x71\xd6\x70\xe4\xd5\xfc\x37\x96\x68\xa3\x91\x13\x5a\xdd\xdd\xd0\x7b\x79\x4d\x4b\xc3\x35\xf8\x1f\xbd\x72\x41\x35\xcc\x5f\x8f\x22\x2c\x2c\xee\xc8\xf9\x23\x7b\xe4\xd0\x53\x57\xba\x1a\xae\x8e\xfd\x3d\x75\x12\x5b\xec\xdc\x2e\x5c\x79\xdf\x07\x98\xae\x22\xe7\x73\xd1\xa2\xe6\x5b\xf5\x97\x6f\xb1\xe6\xbb\x62\x65\x08\x72\x6b\xea\x50\x71\x19\xf4\x5d\x19\x92\xd0\x65\xc0\x81\x41\x87\x14\xbc\x9f\xc3\x97\x99\xdf\x52\x18\x88\x34\xe0\xf1\x4a\xec\x1c\xfd\x0d\x00\x00\xff\xff\x4d\xeb\xf4\xde\xbc\x03\x00\x00")

func jsRegistryJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRegistryJs,
		"js/registry.js",
	)
}

func jsRegistryJs() (*asset, error) {
	bytes, err := jsRegistryJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/registry.js", size: 956, mode: os.FileMode(420), modTime: time.Unix(1453842872, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsSetupJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x56\xdf\x6f\xdb\x36\x10\x7e\x9e\xff\x8a\x9b\x50\xa0\x72\x62\xc8\xfb\xf1\x32\x24\x0b\xf6\xd0\x64\x43\x8b\x2d\x01\x9a\xa0\x2f\x86\x31\xd0\xd2\x49\x66\x23\x93\x02\x49\x25\x31\x86\xfc\xef\xbb\xa3\x44\x89\x76\x9d\xa4\xc5\xba\x01\xf3\x43\x62\x98\x77\xdf\x7d\xf7\xdd\x0f\x72\x7e\x34\x81\x23\xb8\x59\x4b\x0b\xa5\xac\x11\xe8\xbf\x6d\x57\x1f\x31\x77\xe0\x34\xb8\x35\x82\x43\xb3\xb1\x20\x54\x01\xb9\x56\x85\x74\x52\x2b\x0b\x05\x96\x52\x61\x01\x52\xb1\xbb\xf7\x7c\xfd\xfb\xdb\x37\x17\x97\xd7\x17\x99\x7b\x70\xaf\x67\x70\xbf\x96\xf9\x9a\xe1\x1a\x61\x1c\xe8\x92\xb0\x18\x5b\xb7\x26\x47\x42\x2a\x90\x0e\xf2\x5b\x51\x61\x46\x08\x0c\x22\x5a\xb7\xd6\xe6\x04\x70\x23\xf3\xdb\x1a\x25\xfd\x36\x9f\x4c\x3e\xfc\x74\x9e\x19\xcc\x51\xde\xe1\x1b\x51\xd7\x2b\xf2\x81\x33\x28\x5b\x95\x33\x91\x74\x63\xab\x29\xfc\x35\x01\xfa\xdc\x09\x03\x7a\xf5\x91\x4e\xdf\x5d\x5f\x5d\x66\x14\xd6\xa2\x3f\x3f\x1d\x8e\x29\x01\x87\x0f\x8e\x4c\x98\x4c\xf7\xbb\x2c\x21\x25\xb7\x10\xc4\xc0\xb7\x67\x90\xf0\x71\x12\x80\x83\xb7\x12\x1b\xb4\x44\x1a\x2d\x01\xc4\x2e\x99\x6d\x6a\xe9\xd2\x24\x4b\xfa\x58\xfc\x29\xb5\x81\x94\xdd\x24\x59\x7f\x77\x4a\xff\x7e\x8e\x10\xb2\x1a\x55\xe5\xd6\xf4\xf3\xf1\x71\x1c\x87\x3f\x23\xcb\xfe\xdb\x62\xf4\x5b\xc8\xe5\x72\x8c\xf1\x38\x19\xff\x72\x28\x96\x25\x72\x63\x8e\x16\x6b\xaa\xa5\x36\xcb\x31\x5b\x6f\x45\x59\xaa\xb6\xae\xe3\xd8\x06\x5d\x6b\x54\xa7\x9e\x75\x46\xaa\x4a\x96\x5b\x6f\x9d\x89\xa6\xa9\xb7\x69\x8f\x3b\xf3\xc9\x0b\x53\xd9\x69\x9f\xef\x23\x60\x6d\x31\x82\x9a\xcf\xc1\x99\x2d\x21\x36\xda\x38\x02\xf2\x8d\x84\xc6\x68\x33\x98\x30\x13\x02\xb4\x9a\x5a\xe7\x00\x99\x5e\x08\x3e\xce\x6a\x5d\xa5\xc9\xe2\xdd\xf5\x12\x5a\x25\x56\x64\x4f\x9d\xd9\xa0\x21\x85\x37\xc9\x0c\xc6\x12\x8f\x62\xf4\x1c\x6e\xae\xce\xaf\x42\x5a\x3e\xfa\x2f\xfb\xc9\x26\x1c\x37\xe9\xb3\x98\x3c\x4e\x26\xec\xc5\x8d\x9a\x87\x66\xe3\xa6\x45\xe7\xcb\xb9\xa6\x21\xa8\x39\x9b\xd0\x7e\xde\x8a\x06\xc7\xe8\x0d\xfc\x46\xd3\x62\x84\xb2\x25\x05\xa2\xc1\x10\xb6\x13\x92\x01\xdf\x3a\xe8\x11\xe9\x20\x18\xb7\x96\x91\x92\x7b\x6d\x6e\xa9\x83\xae\x51\x15\x69\x96\x65\xd3\x24\xeb\x28\x18\x7d\x4f\x43\x07\xd7\x5b\xe5\xc4\xc3\x05\x73\x07\x7c\xc8\xb1\xf1\x61\x65\xe9\x05\xed\x8a\xe4\xd5\xe0\x66\xe7\x20\x4a\x3b\xea\x84\x5a\x16\x43\xf4\xc9\x2b\x6a\xd3\xbb\xf4\xc0\x1c\x91\x6a\x5f\x3b\x5d\xa2\x25\x1c\x11\x6d\x78\x7f\x88\x20\x32\x11\x6a\xf1\x4b\x94\xa0\xac\xf3\x7f\x47\x0d\x78\xef\x29\x59\xef\xe1\x19\x53\x8f\x1a\xb4\x48\xc0\x1e\x4c\x77\x60\x31\xf3\xf0\xdb\x27\x7a\x8c\xf2\x7a\xc2\x4f\x4b\xcc\x09\x9f\x4b\x9a\x60\xd7\xad\xc4\xd6\x46\xe9\x13\x61\x3e\xa7\xf4\x86\x22\x0c\x82\x7b\x39\xef\x49\x5c\x83\x95\xb4\xb4\x89\xb1\xf0\x41\x39\xd4\x0e\x68\xb4\x12\xc3\x97\xf7\x58\xc2\xfc\x68\x06\x34\xa8\xed\x86\xf2\xb3\xb4\x4f\x21\x5e\x96\x3c\xc1\xe4\xb8\x58\x66\xb6\x96\x39\x7a\xc0\x74\xb0\x9e\xfa\xad\x96\x63\xfa\x7d\xbc\x40\xc7\x15\xcc\x14\x42\xa8\x3f\x3b\x7a\x66\x9b\x39\x71\x8b\x31\x83\xe9\xb0\x75\x0e\xdb\x2b\xad\x10\xce\xce\x06\xe0\x78\x05\xbc\x6a\xa8\xa0\xae\x9f\x7c\xa5\x47\x55\xb8\x39\x0d\x52\xe7\xa1\xca\xf1\x24\x81\x63\x88\x23\x9e\xee\xcd\xf8\x69\xb4\x24\x43\x98\x7e\xa1\xf1\x9a\xf7\x02\x59\xf2\xea\x86\xff\x0f\xb4\x96\x2e\x25\xee\x42\xae\x94\xf0\x4b\xc8\x99\x96\x77\xe8\x20\x7d\x6c\x34\x2a\x1f\xfb\xa6\xe1\x6e\x98\x41\xd8\xc0\xb4\x36\x55\xd7\x7d\x21\x49\x0e\x9f\x15\xc2\x09\x02\x19\xd3\x4e\x82\x6b\x72\x02\x03\xca\x78\x1a\xe0\xe8\x74\x40\x1e\x4f\x43\x82\x74\x1a\xa2\x45\xa7\x9c\x29\x9d\xbc\x54\xf1\x1f\xbb\xaa\x3d\x1e\xd2\xc4\xe9\x7e\x6a\x9e\x1d\xa3\x43\x4a\x65\x8d\xd1\x4e\xbb\x6d\x83\x43\xba\x1d\xd8\xc9\xd8\xbb\x2f\x5f\x47\x83\x66\xd3\xdd\xa5\xcd\xa9\x74\xf9\x86\xab\x81\x8b\xb7\x53\x4d\xc5\xd3\xc6\xef\x98\x98\x7c\x37\xe1\x7e\x16\xfd\xb0\x51\x12\x6d\xed\xe2\x19\xeb\x51\xa3\x09\xfb\xb4\xb6\xcf\x0f\x1a\xdd\x50\x5f\xaf\xc2\x9f\x59\xc3\x1f\x86\x1a\xee\x6b\xd9\x3d\x8c\x5e\xd9\xb0\x67\xf7\x04\xe6\xeb\x74\x3a\x8c\x83\xdf\x4c\x2f\x0a\x5a\x68\x7a\x13\x5d\x5e\xdd\x84\x38\x62\x5c\xf9\x83\x8c\xff\x6f\x01\xbd\x5c\x07\xa5\x8a\x95\xba\x59\xe3\x0b\xed\xd7\xbd\x8a\x73\xd1\xdd\xa5\xbe\x03\xc3\x9c\xee\x5c\x2b\x7c\x53\x7a\x55\xf9\x0a\xeb\x55\x1d\xdb\xb5\xbf\x90\xfa\x48\x3b\x32\x7b\x0a\xcf\x4a\x3d\xee\xa1\x5f\xc3\xe2\xfa\xef\xc4\x8f\xf6\xd3\xe1\x0b\xa1\x69\x5d\xba\xcf\x6f\xfa\x4f\x56\xd8\xe7\x55\x8f\xdf\x3c\xa2\x28\xec\x9c\xf6\x58\xed\x9f\xf8\xd1\x72\xa0\xf6\xe4\x8e\x36\xd2\x3f\x3e\x65\x57\x89\xaa\xd6\x2b\xc1\xaf\x21\x89\x75\xc1\x15\xe1\x7c\x42\x25\x18\x2e\x2a\x42\xf0\xbd\xa4\x77\xfc\x4c\x3a\xfb\x81\x71\x59\xe3\x6f\xd8\xb8\x47\x5a\xc4\x56\x4b\x72\x0f\x86\x81\x63\x85\xee\xd0\xea\xfa\x72\x76\xd5\xd3\xec\x3c\xa9\xbe\xdf\x9e\xe4\x46\x84\xfe\x0e\x00\x00\xff\xff\x0a\x09\xab\x8e\x37\x0e\x00\x00")

func jsSetupJsBytes() ([]byte, error) {
	return bindataRead(
		_jsSetupJs,
		"js/setup.js",
	)
}

func jsSetupJs() (*asset, error) {
	bytes, err := jsSetupJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/setup.js", size: 3639, mode: os.FileMode(420), modTime: time.Unix(1453843509, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func getAsset(name string) ([]byte, error) {
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
func mustAsset(name string) []byte {
	a, err := getAsset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func assetInfo(name string) (os.FileInfo, error) {
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
func assetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"js/console.js":  jsConsoleJs,
	"js/registry.js": jsRegistryJs,
	"js/setup.js":    jsSetupJs,
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
func assetDir(name string) ([]string, error) {
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
	"js": &bintree{nil, map[string]*bintree{
		"console.js":  &bintree{jsConsoleJs, map[string]*bintree{}},
		"registry.js": &bintree{jsRegistryJs, map[string]*bintree{}},
		"setup.js":    &bintree{jsSetupJs, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func restoreAsset(dir, name string) error {
	data, err := getAsset(name)
	if err != nil {
		return err
	}
	info, err := assetInfo(name)
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
func restoreAssets(dir, name string) error {
	children, err := assetDir(name)
	// File
	if err != nil {
		return restoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = restoreAssets(dir, filepath.Join(name, child))
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
