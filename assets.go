// Code generated by go-bindata.
// sources:
// assets/index.tmpl
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

var _assetsIndexTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x3b\x79\x53\xdb\xba\xbe\x7f\x27\x9f\x42\x75\xef\x6d\x9c\x17\xb2\xa7\x14\x08\xc9\x0c\xd0\x8d\x6e\x70\x0b\xed\x59\xda\xbe\x8e\x63\x2b\xb1\x8a\x6d\xb9\x92\x4c\x08\xe7\xf0\xdd\xdf\x68\xb1\x2d\x2f\x09\x09\xd0\xf3\xe6\xa6\xd3\xc4\x96\x7e\xfb\xee\x85\xfd\x47\xcf\x4f\x8e\xce\xff\x38\x7d\x01\x5c\xe6\x7b\xe3\xea\xbe\xfc\x01\x00\x80\x7d\x17\x5a\x8e\x3c\x14\xa7\x3e\x64\x16\xb0\x5d\x8b\x50\xc8\x46\xc6\xa7\xf3\x97\xcd\x1d\x43\xdb\x66\x88\x79\x70\x7c\xe6\xe2\x88\x01\xca\xa2\xe9\xf4\xd1\x7e\x5b\xae\xa5\x30\x1e\x0a\x2e\x92\x33\xfe\x21\xd0\x1b\x19\x94\x2d\x3c\x48\x5d\x08\x99\x91\xd9\x74\x09\x9c\x8e\x0c\x97\xb1\x90\xee\xb5\xdb\xbe\x75\x65\x3b\x41\x6b\x82\x31\xa3\x8c\x58\x21\x3f\xb1\xb1\xdf\x4e\x16\xda\xfd\x56\xbf\xf5\xac\x6d\x53\x9a\xae\xb5\x7c\x14\xb4\x6c\x4a\xb3\x84\x51\xc0\xe0\x8c\x20\xb6\x18\x19\xd4\xb5\xfa\x3b\x83\xe6\xe1\xe7\x3f\x10\x3a\x3b\x7e\x09\xdf\x76\x9d\x57\xfe\x9b\x8f\x07\x17\x0b\x3b\x7a\x7d\xf0\xfa\xe3\xac\xdf\x3b\xf1\x3f\xd9\xf3\xf9\x33\x1c\xf4\x3f\xfe\xe1\xcc\x06\x9f\xad\xc6\xa9\x7f\x76\x4e\xaf\xdb\x6f\xb7\x77\x2e\x27\xce\x8b\x1f\xee\x20\xca\xd2\xb7\x09\xa6\x14\x13\x34\x43\xc1\xc8\xb0\x02\x1c\x2c\x7c\x1c\x51\x63\x5c\x4d\x2d\x21\x74\x1e\x57\xab\xd5\xea\xff\x80\xbf\xc4\x72\x68\x39\x0e\x0a\x66\x7b\xa0\x33\x14\xe7\xbe\x45\x66\x28\x10\xa7\x37\xd5\xea\x04\x3b\x8b\xdb\x01\x85\xd5\x20\x9a\xb9\x6c\x0f\x74\x3b\x9d\x4b\x57\x2e\xcd\x91\xc3\x5c\xb1\xf2\xef\xfc\xc2\xe5\x5c\xae\x4c\x2c\xfb\x62\x46\x70\x14\x38\x4d\x1b\x7b\x98\xec\x81\xc7\xb6\xed\x38\xd3\xa9\xe0\xde\x72\x91\xe3\xc0\x40\x09\x70\x89\x28\x9a\x20\x0f\xb1\xc5\x1e\x90\x1b\x12\xc8\xc6\x01\xb3\x50\x00\xc9\xaf\x12\x94\x33\x99\x62\xe2\x37\x29\x9a\x05\x28\x16\xc7\xb7\xae\x9a\x0a\xae\xdf\xef\x84\x57\xc3\x2c\xf3\xee\xd3\x78\x29\xe1\x0f\xac\x88\xe1\x22\x39\xfd\xa4\xc9\x83\x1f\x05\xb3\xad\x1c\x88\xed\x42\xfb\x62\x82\xaf\x12\xde\x9c\x64\x73\x82\x19\xc3\x3e\x97\x93\xf3\x2a\xd0\xcd\x21\x4d\x71\xc0\x9a\x73\xa5\x7d\x80\x89\x6f\x79\xcb\x84\xe1\x26\x25\xd8\x8b\x0d\x8a\x29\x62\x08\x07\x7b\x3c\x6f\x2c\x86\x2e\x61\xd6\x92\x52\x2d\xbe\xd2\x9c\xc3\xc9\x05\x62\xcd\x09\xbe\x6a\x52\x74\x2d\x0c\x31\xc1\xc4\x81\x84\x2f\x0d\xe3\x38\x6c\xfa\xf8\xfa\x36\x18\xf1\x59\x05\x93\x5a\x3a\x31\xbe\xd0\x90\xa2\x6b\xb8\x07\xba\xdb\xa5\x26\xd1\xb5\xdb\x9b\x62\x3b\xa2\x4a\xc7\xeb\x26\x0a\x1c\x78\xb5\x07\x7a\x45\x2c\x14\x84\x11\xfb\xc2\x16\x21\x1c\x19\x0c\x5e\x31\xe3\x5b\xb9\x1b\x9a\xdd\x58\x90\x44\x56\xbe\xd1\x24\xdc\x4c\x4d\x62\x39\x28\xa2\x49\x18\x66\x41\x3c\x38\xcd\x40\xac\x12\x21\xb4\x28\x9d\x63\xe2\x2c\x13\x23\xb5\x87\xe2\xc1\x70\x58\x60\x90\xdb\x2e\x88\x78\x53\xad\x3e\xf6\x21\xa5\x30\x98\x25\x89\x95\xda\xa8\xb3\x34\x7b\x5f\xbe\x7c\x99\x4b\xb3\xed\x34\x37\x92\x38\x9a\xa2\x2b\xe8\xc4\x42\x48\xa9\x3b\x4b\x52\x31\x97\xbd\x89\xd7\x7b\xf9\x94\x13\x2a\xf2\xf5\x7f\x67\x97\x89\x94\x23\x59\xcf\xc4\xa8\x6b\x39\x78\xbe\x07\x3a\xe1\x95\xf8\xff\xb4\xcf\xbf\xc2\x2b\x40\x66\x13\xcb\xec\x6c\x89\x7f\xad\x5e\xbd\x24\x72\x37\xc5\x4c\xe3\x79\x6d\xcc\x24\x08\xf2\x8e\xc8\xa9\x26\x4b\x4d\x09\x34\x0f\x56\x8b\x40\xeb\x97\x87\xab\xec\xa9\x32\xf3\x02\x1c\xc0\x52\x69\x1e\x53\x18\x38\x87\x2c\xf8\x07\xa2\xd6\x9a\x41\xfa\x0e\x51\x96\xb3\x17\xc3\xe1\x1e\x48\x6b\x75\x12\xcf\xdd\xce\x1a\x01\xaa\x05\xb2\x8c\x34\x85\x24\x88\xc6\x36\x90\x0e\x51\x67\xf8\x12\x92\xa9\xc7\x7d\x9d\x56\xc8\x0d\x82\x95\x2b\x63\xd9\x5c\x9e\xb8\x48\xf9\x28\x88\x5b\x4e\xb7\xb7\x3a\xad\x72\xaa\x2a\xba\x83\xa2\xee\x3d\x55\x6d\x94\xd9\xb2\xce\xd1\x75\x2b\xe9\x6c\xab\xf2\x5f\xb9\x29\x76\x4e\xea\xde\x18\xb2\xb7\xb3\xdd\xd9\xed\xac\xce\xc9\xa7\xda\xff\x4c\x7a\x74\x6e\xcb\xc9\xb5\x30\x4b\x73\x72\x35\xa6\x30\x15\x9d\x7d\x47\xc1\x14\xeb\x7d\x55\xc6\x3e\xf5\x2d\xcf\x4b\x81\x78\xfe\x15\x81\x3c\x8b\xcc\xe0\xb0\x2c\x07\x54\xcf\x4a\xa6\xb5\x76\x3c\xae\xc9\x33\x39\x18\xcb\x13\x3e\x98\x69\x13\xae\x83\x2e\x01\x72\x46\x86\x87\x67\x28\x78\x89\x89\x6f\x00\xdb\xb3\x28\x1d\x19\xc9\x74\xa4\x0d\xcd\x02\x85\x67\x67\x0c\xa4\xb5\x9a\x1c\x98\x00\x75\x7b\x25\x80\xf1\xa8\x62\x8c\x7f\x83\x9e\x8d\x7d\xb8\xdf\x76\x7b\x25\xc8\x9e\x35\x81\x1e\x98\x62\x32\x32\x44\x17\xfb\x44\x21\x09\x2c\x1f\x26\x12\x52\xd2\xc4\x81\xb7\x30\xc6\xf1\xce\x7e\x5b\xe0\x94\xd0\x12\x04\x80\xd6\x89\x85\xd2\xe5\x64\xf5\x56\x6f\x80\xd0\xb3\x6c\xe8\x62\xcf\x81\x64\x64\xa4\xb0\x04\xfe\x8c\x10\x81\x8e\xc8\x50\x31\x0f\xac\xa1\xc1\x69\xdc\x80\x0b\x1a\xc4\x3b\xeb\x69\x90\x34\xf2\x54\x8b\x02\xe9\x15\x5a\xa4\xb0\xb1\x16\x25\xfc\x32\x81\x71\xc8\x82\x84\xf0\x84\x05\x60\xc2\x82\xa6\x37\x13\x3f\x21\x41\xbe\x45\x16\xe2\x78\xe2\x61\xfb\xc2\x50\x42\xd2\x68\xe2\x23\x66\x8c\xcf\xd0\x8c\x8f\x21\xfb\x6d\x07\x5d\xe6\x22\xa9\xcd\x85\xd4\x82\x51\x82\x14\x83\xd3\x0a\xc3\x62\x58\xaa\x51\x3e\x1f\x9d\x31\x4e\xd2\x3e\xca\xe2\xb2\x10\xc2\xab\xa0\x05\x46\xd2\x11\x05\x6d\x3a\x3b\x17\x41\x54\x6a\x6b\x82\xe7\x74\x64\xf4\x8d\xf1\x7e\x3b\xc6\x5a\x42\x34\x96\x55\xf5\xb7\x82\x89\x6f\xb7\x2d\xbf\x7a\x2d\xb1\x6c\x99\x75\x35\x0b\x97\x0a\xa1\x3a\x46\x99\xb9\xb4\x58\xc0\x11\x5b\x4f\xd2\x72\x95\xdf\x09\x0a\x60\x9f\x86\x56\x20\x48\xce\x5d\x6c\xf9\x28\xa1\xc7\x8d\xc6\xf7\xca\xd4\x29\x89\x9f\xe5\xda\xe8\xcd\x5c\xbf\x90\x2d\xc1\x8b\xc3\x4e\xbb\xd8\xb5\x09\x0a\x59\x06\x87\x12\x3b\xbd\xb4\xb7\xb1\x03\x5b\x3f\x7e\x46\x90\x2c\xc4\x25\xbd\x3c\x6c\xf6\x5b\xdd\x56\x47\x5c\xc1\xff\x58\x79\x01\xdf\x7b\xba\xdd\xb4\x8f\x22\x78\xf8\x71\xfb\x88\x2e\x0e\x06\xed\x5d\x7a\x1d\x9e\x4e\xc9\xef\x7d\x3a\xd8\x7d\xbf\x7b\xf9\xe9\xd3\xd3\xc3\x19\x43\x6f\x7e\x74\xb6\xe7\xac\x4d\x47\x6b\x5e\xac\xa7\xfa\x48\xf9\xc7\xeb\x2b\xb4\xee\xbd\x8a\x1f\xf9\x5b\x15\xab\x15\xed\xef\x0c\x9a\xe7\xf6\xd3\xe3\xff\xa0\x49\xa7\xf7\xec\xe7\xe5\xe2\xc7\xd9\xfb\xe9\xeb\x1f\x27\xef\xad\x77\x17\xd3\xe8\xb7\xcf\x57\x7f\x5e\x7d\x3a\x0d\x8e\xde\x1c\x3c\xf3\x7a\xfe\xd1\x6f\x1f\x8e\xc3\x57\xbb\xfe\xab\xa3\xe7\x3b\xf3\x57\x1f\x8e\xed\xd3\xe7\xcf\xce\xaf\xac\x87\x52\x7e\x5c\x9d\x46\x81\x08\x71\x20\xb2\xc6\xac\x83\xbf\xaa\x95\x5a\x44\x21\xa0\x8c\x20\x9b\xd5\x86\xd5\x4a\xbb\xcd\x5c\x44\x5b\x56\x88\xc0\x08\xd4\xb8\x75\xf6\xda\xed\x6e\xef\x59\xab\xd3\xea\xb4\xba\x7b\x3b\x9d\x9d\x4e\xdb\x0a\x51\xbb\xa6\x66\x25\x0d\xf8\xcb\x97\x2f\xad\x83\xd3\xe3\x6f\xdf\xbe\xd5\x86\xd5\x74\x37\xa2\x90\xf0\x6d\x1d\x83\x97\x6f\xb5\x56\xad\x88\x15\x07\x4e\xa2\x19\x18\x81\xa9\xe5\x51\xa8\xa3\x7b\x78\x36\x83\xce\x71\x90\xee\x55\xe2\x75\x14\x9c\xc2\x80\xf7\xd1\x32\x3c\x5e\x51\xde\xcb\x14\x58\x01\xe5\x21\xca\x14\x14\x2d\x82\x69\x9c\xf8\xaa\xb2\x9e\xc9\x35\xda\x02\x5c\x87\x2d\x60\x5b\x9e\xc7\x67\xb9\x93\xb7\xe9\xf1\x87\x93\xb7\x75\x35\xbb\x88\x88\x98\x02\xb3\x28\xf2\xdf\x7f\x67\xf5\x13\xce\xa8\x54\x08\x64\x11\x09\x86\xd5\x4a\xe5\xa6\x5a\xad\x94\xaa\xca\x48\x14\xeb\x90\x37\x32\xff\x19\x66\x77\x94\xa9\xf9\x4f\x1e\xc9\xb7\x2e\xe0\x47\xf8\x33\x82\x94\x99\x9c\x77\xed\xd5\x8b\xf3\xda\x96\x38\x92\x45\xa9\x2d\xcf\x12\xc5\x29\xf4\xa6\x5b\xfc\x6a\x25\xc4\x01\x85\xba\x8e\x49\x74\x26\xe6\xc8\x03\x0f\x0b\xb0\x1c\x40\x77\xaf\xd4\xab\x52\xa9\x54\xe2\x9d\x12\x07\x57\x2a\x95\x9b\xbb\x0a\xf5\xa1\x4c\xaa\x75\xf9\xd5\x6a\xd5\x4a\x85\xc3\xdf\xe4\xa2\x93\x57\x73\x2d\x38\x62\x66\xba\x20\x65\x79\x50\x96\x0b\x99\xf5\x62\xe0\xe7\xf5\x11\xca\x28\xc3\xde\x24\xe1\xaa\x05\xbe\x2e\x97\x6a\x07\x2b\x23\xb6\x52\x49\x62\xb5\x98\x3e\x25\x01\x5a\x29\x70\x2c\x44\xa9\x02\x29\x84\xda\xe9\xc9\x59\x1c\x6b\xfc\x7a\x80\xf8\x16\x97\xf2\xf6\x80\x4b\xdc\xb5\x2a\xbf\x37\x89\xca\x75\x02\xea\x5e\x4c\x97\x45\x9d\xe4\xaa\xdc\x52\x1a\x5b\x5a\x6d\x2a\x8b\xb0\x25\x25\x27\xf5\x61\x49\x71\x5b\xee\x44\x0d\x78\x6d\x27\xa6\xf5\xe2\x6e\x3e\x5c\x59\x7d\x7f\x99\x13\xef\xc4\x75\xb5\x17\xe3\xe2\x10\x67\x62\xe2\xc3\x19\x64\x65\xb9\x48\x67\xc7\xce\x2d\x99\xb8\x91\xc5\x41\x03\x08\x9a\xa0\x01\x6a\xed\x1a\x78\xc8\xaa\xfd\x2b\x0a\x6e\xb9\xd1\xa4\xca\x36\x81\x16\x8b\x95\xce\x16\x30\xe6\x62\x67\x0b\x44\xc4\x93\x16\xba\xb4\x08\xb8\x72\x79\x4d\x0d\xe0\x1c\xfc\xfe\xfe\xdd\x6b\xc6\xc2\xd8\x5a\x82\x11\xcf\x84\xda\x1c\x31\xf7\x88\x40\x07\x06\x0c\x59\x1e\xad\x01\x14\x70\x34\x15\x12\x57\x2e\x69\xe1\x10\x66\xc8\x6f\x89\xb8\x17\x14\x6e\x00\xf4\x28\x94\x2d\x7c\x11\x42\x3c\x05\xbf\x3f\xc7\xbe\x85\x82\x58\xc0\x47\xa3\x11\xa8\x45\x81\x03\xa7\x28\x80\x4e\x2d\xa5\x1a\x8b\x95\x01\x97\x62\x95\x32\xd5\xd9\xe9\x34\x22\xcf\x93\x79\x1a\x1b\x97\x23\x53\xc8\x14\xc5\xd7\xd0\x72\x20\x31\x8d\x83\x88\xb9\x98\xa0\x6b\x11\x10\xc6\x16\x30\x0e\x2d\x8a\x6c\x60\x80\x06\x38\xb4\x28\xdc\x1e\xb4\x60\xc0\x87\x76\x33\xed\x45\x0d\x60\xec\xf1\xfd\xa4\x0b\xd5\x85\x0c\xb2\x36\x70\x36\xaa\x20\x15\x42\xb1\xcc\x2b\x7c\x61\x69\x44\x6f\x01\xc7\x62\x56\xea\xb5\x88\x78\xbc\xb8\xc4\xb3\x63\x43\x60\x0f\x33\x2e\x2d\xc6\x42\xc1\x58\xdc\x29\x8f\x52\x4f\xa2\xa9\x99\x4e\x92\x75\xf0\x17\xb0\x71\x40\xb1\x07\x79\x23\x35\x6b\x47\x27\x1f\xcf\x40\x80\x19\xa0\x51\x18\x62\xc2\xb8\xaf\x86\xe0\xa6\x64\xe4\xe2\x32\xf0\xa0\x55\x42\xf0\x75\xe1\xaf\xc0\xc3\x96\xa3\xab\x9e\xf2\x35\x85\x4f\x98\xc5\x22\x5a\x6f\x31\x7c\xc6\x08\x0a\x66\x66\xfd\x4b\xe7\x9b\x8c\x8f\x5e\x1c\x17\x95\x62\x66\x08\xcb\xc8\xac\xd0\x9c\x5f\x29\xa4\x24\x67\x11\x67\x11\xbf\xf4\xce\x60\x72\xd9\x45\x9d\x96\x92\x42\x42\x30\xc9\x88\x2a\x88\xea\xcc\xb9\x6a\x1a\x89\x14\x9b\xb7\x38\x33\x5e\xbf\x11\x37\xe4\xb8\x45\x64\x10\x81\x11\xf8\xab\x5a\x6d\xb7\x41\x48\xd0\xa5\xc5\x20\x08\x09\x0e\x21\x61\x8b\xea\xf7\x0b\xb8\x38\x63\x04\xec\x01\xe3\xe0\xf0\xe8\xf9\x8b\x97\xaf\x5e\x1f\xbf\x79\xfb\xee\xfd\x87\x93\xd3\xff\x7c\x3c\x3b\xff\xf4\xf9\xb7\xdf\xff\xf8\xd3\x9a\xd8\x0e\x9c\xce\x5c\xf4\xe3\xc2\xf3\x03\x1c\xfe\x24\x94\x45\x97\xf3\xab\xc5\x75\xa7\xdb\xeb\x0f\x9e\x6e\x3f\xdb\xd9\x6d\xb4\x47\xc6\x96\xe4\x11\x4d\x3c\x64\x03\xe9\x76\x30\xc5\x04\x88\x10\x46\xc1\xac\x2a\x63\x19\xec\x25\x1a\x02\x53\xdc\x03\x8a\x8b\x11\x97\x18\x47\x2c\x14\xa3\x99\x61\x0c\x93\x45\xdb\x25\xdd\x2d\xfe\xdd\x13\xdf\xfd\x2d\x4e\xb3\x2b\xbe\x7b\xe2\x5b\xae\x0c\x52\x0c\x7e\x69\xd3\x51\x25\x5d\xde\x81\x1a\xc5\x19\xf5\x3d\x62\xd3\x9d\xef\x2a\xaf\x24\x7f\x05\x38\x77\x91\x07\x81\x89\xc0\xbe\xc4\x69\x79\x30\x98\x31\x97\x8b\x97\x4e\x71\x2e\xe9\x82\x91\xda\xb7\x5d\x8b\x1c\x61\x07\x1e\x30\x13\x35\x1a\xda\xb0\xcc\x45\x5d\x0b\xaa\xbf\x1c\x2a\x01\xe3\xba\x82\x91\xe4\x3c\x1e\x83\xde\x50\xdf\xe1\x6c\x4c\x53\xec\x3d\x01\xfd\x3a\xd8\xdf\x07\x83\x3a\xf8\x1b\x98\x42\x84\xf1\x18\x0c\xea\x19\xf8\x7e\x0c\xdf\x03\x4f\x40\xf7\xa9\x40\xe8\xc5\x08\x7d\x8e\xb0\x9d\x45\x18\x48\xd6\x7d\xf0\x04\x6c\xf7\x35\xa9\x78\x1e\x23\xfa\xc1\xfa\x20\x88\xd5\xf3\x0d\x45\xb1\x52\x04\xb6\x07\x29\x4d\xad\x38\x27\xf8\xfd\x32\xfc\x02\x62\xca\x3c\x89\x12\x75\xd0\xc8\x0e\xe1\x2a\xac\x85\x4d\x0f\x98\xc9\x0d\x58\x8f\x4b\x66\x71\xaf\x57\xbf\x1d\xbf\xbf\x02\x7f\x10\x3b\x4b\x09\xa8\x4a\xb1\x14\x6d\x58\xbd\x59\x96\x17\x0e\x54\x79\x21\x0e\x1e\x22\x2f\xd2\x9d\xbb\xe4\x87\x8c\x43\x02\xc5\x3d\x57\xb3\xfd\xe5\x7f\x0f\x9a\x7f\x5a\xcd\xeb\x4e\x73\xf7\x6b\xe3\x6b\xfb\xeb\xe8\x5b\x7b\xb6\x05\x0c\x63\x93\x6c\x51\xa1\x9b\x31\x9c\x78\x0c\x73\x32\x35\xd3\xb0\x57\x21\x5f\x2f\x04\xf6\x5d\xf0\xfa\x77\xc4\x1b\x6c\x82\x97\x2f\x07\x22\xc4\xd2\x54\x12\xe2\xe7\x72\x4f\x95\x04\x53\x6e\x26\xb9\x37\x88\x11\x44\xee\xf5\x8a\xd5\xc1\x94\x9b\x71\x72\x6f\x73\x78\xe9\xcb\xe5\xd9\x00\x64\x2f\x6b\x4d\x09\xf6\x8f\x54\x5d\x11\x35\xa2\x9e\x4b\x60\x41\xfa\x11\x4f\xb3\x7c\xfe\xad\x4d\x53\x17\xf9\x26\x4f\x7c\x70\x4f\xe2\xfd\x7a\x21\xfd\xd5\x4f\x42\x22\x53\xd5\x65\x26\x99\x72\x33\x56\x36\x97\x8f\x49\x42\xaa\x66\xa8\x65\xa4\x78\xd7\x2a\xed\x57\x7a\xa7\xc8\x64\x27\x15\xd2\xc6\x5a\xc9\x33\x30\x52\x07\x69\x06\x7d\x25\x5f\x83\xf6\x6c\xcb\xf8\x1a\x18\xf5\x34\xf5\x22\x36\x15\x8f\xc9\x64\x2a\xab\x67\x65\x04\x98\x7c\x2f\x10\x69\x09\x02\xb0\x1f\x13\x93\x49\x35\x04\x41\xa3\x91\xc9\x2c\x91\xfe\x29\x4f\xad\x7b\x04\x79\x27\xdb\x60\x1f\x74\x7b\x3b\x79\x27\xc4\x72\x34\x46\xe5\xd6\x2f\xf5\xab\x2a\xdd\xa6\x69\x83\x31\xe8\xf6\x9e\xd5\xc1\x93\x27\x92\x43\xaf\x33\xd8\x29\x54\xf1\x5b\x78\x70\x2a\x63\x19\xd2\xdd\xdd\xfc\x63\xca\xdb\x71\x79\x47\x12\xb8\xbd\x9d\xe5\xc2\xde\x41\xa0\xae\xc8\xe2\x5e\x6f\xb0\xa1\x44\x89\x3a\xe5\x82\xdd\x4f\xa9\x92\x0e\xa3\x88\x0d\x6f\x8d\xe8\xa4\xd3\xe8\x59\x92\x89\x68\x45\x4a\xef\x38\x49\x58\xeb\x1d\x47\xf5\x0d\x3d\x02\xe5\x70\xd2\xd3\x1a\x8a\xea\x0c\x80\xb7\x06\x45\x58\xc5\x31\xc8\xce\x52\x60\x94\xec\xeb\xe3\xcf\x9a\x01\xac\x04\x5c\x23\x7e\x05\xa1\x46\xe3\xf6\x80\xde\xed\x6a\x01\xdd\x1b\x14\xe2\x59\xa8\x59\x26\x72\xa3\x9b\x63\xb7\x5a\x38\xe9\xe8\x7e\x37\x2d\xe9\xa6\x68\x0a\xdb\xfd\x7a\x5e\x6c\x4e\xa2\xb7\x66\x74\xaf\x2f\x9e\x68\x93\xa5\x90\xf9\x34\x5c\x47\x11\xd5\xcb\x64\xde\x98\x89\x2a\x9a\x72\xfd\xe5\xca\xf5\x6f\x89\x72\x29\x80\xb8\xa4\xb9\xd1\x1f\x76\x2d\x7f\x60\xa2\xb3\x10\x91\xac\xae\xea\xd5\xf3\x93\x61\xf6\xa9\xda\xbf\xcc\xda\xe3\xf8\x61\x71\xad\xde\xb2\x3d\x64\x5f\x98\xfa\x35\x18\xc8\x7d\x44\x11\x97\xb7\x85\xff\x65\x1a\x8f\xb3\x0f\xe3\xeb\xad\x4b\xcb\x33\x4b\xee\x9b\x73\x2c\x75\xd3\x38\xc1\x4a\x1e\x68\x2f\xc5\x52\x8f\x11\xcc\xc2\x06\xff\x88\x27\x1a\xa5\x3b\xe2\x29\x47\xe9\xce\x06\xf7\x83\x52\x03\x19\x8f\xd3\xd7\x2c\xea\x2d\xcb\x71\x8e\x3c\x8b\x52\xd3\x50\x4f\xb2\x4b\xe4\xd6\x71\xad\x30\x34\xea\x2d\x02\x7d\x7c\x09\x37\x42\x54\xcf\x58\xeb\x2d\x1e\xa5\xe6\x8a\x67\x12\xf1\x47\x78\x26\x74\x2c\x06\x3f\x23\x38\xd7\x2f\xa5\xc5\x3d\xf7\x15\x3a\x82\xe4\xc9\x86\x76\x9b\xb1\xdc\xea\xfa\xe7\x0e\xd6\xcc\xcb\xcb\x19\x82\x11\x78\x73\x76\xf2\xa1\x15\x5a\x84\xc2\x75\x14\xd5\x3f\xbc\x42\x4a\x22\xf2\xbe\xd3\xba\xcc\x41\x92\x62\x6b\x81\xdf\xac\x05\x95\x4c\x32\xb2\x51\x88\x16\xc0\x85\x4b\xe6\x18\x24\xe7\x98\x75\x05\x14\x3e\x49\x6f\xc7\xde\xee\x11\xfd\xc3\x19\x7f\x41\xdf\xca\x13\x61\xa9\x06\xf7\x73\xa9\xfe\xe1\x76\xf0\xe9\xec\x5e\xde\xcd\xd3\x73\x99\xef\x81\x11\xa8\x89\x17\x06\xd4\xcb\x07\xea\x69\x84\xa1\xbf\x40\xf0\x5d\xdd\x5d\x6e\x21\x71\x7b\xd9\x18\xd7\x36\x66\xa7\x58\x89\x9f\x46\x8e\xa3\x7a\xd7\xcb\x18\xc7\x6c\xe2\xb7\xe9\x38\xa0\x78\x2f\xe1\xc1\xf9\xa1\x60\x8a\x1f\x40\x8d\x58\x50\xe4\xec\xa9\xd7\x39\x74\x1e\x4e\xaa\x91\x34\x9c\x7a\xad\x03\xfc\x0d\xee\xcd\x39\xc4\x94\x41\x07\x4c\x16\x25\x8c\xad\x88\xb9\x29\x6b\x75\xa3\xf7\x21\x99\xf3\xa2\x58\xc2\x97\x2f\xa7\x7c\xf9\x99\xce\xf7\xfe\x3e\x7c\xa0\x50\x10\x64\xe2\xef\x09\xb9\x03\x45\xde\xe5\xf5\x77\x6b\x6a\xf5\x96\x15\x86\x30\x70\x4c\xce\xe4\x8e\xc9\x08\x3d\xe8\xc3\x80\x97\x6f\x07\xdb\x11\x3f\xe4\xb5\xea\x85\x5c\x3d\x5c\x1c\x3b\x66\x2d\xcb\x73\x73\x36\x8a\x45\x8b\xda\x04\x7b\xde\x39\x0e\xc1\x28\xb7\xf6\x5a\xbc\x77\xbe\x19\xe5\x9b\xff\xb7\x9a\x98\xb9\xbd\x1f\xe0\x8b\x4d\x6d\xb2\x5e\x1b\xe2\x9f\x35\x09\xdf\x4e\x70\x0d\x63\xdd\xd3\x40\x77\x31\xca\x6a\xb9\x57\x10\xb8\x59\xbe\x95\xce\x4e\xfa\x4b\x0a\x05\x02\x0f\x36\x59\xae\xab\x77\x51\xd7\x1c\xe0\xcd\x92\xa1\x5e\xbe\xf5\xb7\xd6\x54\x1f\xbf\x15\x52\x3e\x62\xac\x3d\x48\xe6\x87\xe5\x3b\x8c\xbd\x72\x5e\xde\x6c\xca\x2e\x14\x37\xe8\x87\x6c\x51\x76\x45\x01\xee\x6c\x4f\xf5\xb6\xe7\xda\xd7\x48\x72\xf4\xe1\x1a\xc5\xef\x9d\xae\xb8\xcc\xd1\xde\xd1\x28\xf7\x80\x4f\x67\x0f\x16\x77\x4a\xba\x63\x67\xe3\xd1\x6c\xa3\x79\x54\xbe\x30\xb0\x12\xe4\x1e\x85\xe3\xbe\xd3\xe5\x2f\x9e\x26\x7f\xf5\xf4\xf8\x40\xd3\xe2\x3f\x37\x1d\xfe\xa3\xd3\xe0\x3f\x33\xfd\xdd\x71\xda\xbb\xff\x74\x77\xef\x69\xee\x17\x4d\x6f\x0f\x33\xad\xdd\x32\x70\xdc\xa3\x66\x6c\x3a\x68\x2c\x1f\x32\xfe\xbb\x66\x83\xf8\x38\x7b\xd3\x70\xbf\x2d\xff\xa8\x68\xbf\x2d\xff\x18\xff\xff\x02\x00\x00\xff\xff\x61\x07\xe4\x9d\xa4\x3f\x00\x00")

func assetsIndexTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexTmpl,
		"assets/index.tmpl",
	)
}

func assetsIndexTmpl() (*asset, error) {
	bytes, err := assetsIndexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.tmpl", size: 16292, mode: os.FileMode(420), modTime: time.Unix(1474551228, 0)}
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
	"assets/index.tmpl": assetsIndexTmpl,
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
	"assets": &bintree{nil, map[string]*bintree{
		"index.tmpl": &bintree{assetsIndexTmpl, map[string]*bintree{}},
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
