// Code generated by go-bindata.
// sources:
// rpc.swagger.json
// DO NOT EDIT!

package logger

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

var _rpcSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5c\xff\x73\x1a\x39\xb2\xff\x3d\x7f\x85\xca\xef\xbd\x32\x6c\x30\xc6\x76\xd6\x5b\x71\x2a\x3f\x38\x36\xc9\x71\x71\x0c\x67\xf0\xe5\xd5\x1e\x5b\x58\xcc\x08\x98\xcd\x20\x71\x92\x06\x42\xf6\xe5\x7f\x7f\xdd\x92\xe6\x2b\x33\xd8\xc4\x76\xdd\x5d\xd5\x6e\xed\x6e\x3c\xfa\xd2\xea\x6e\xb5\xba\x3f\xdd\x92\xf3\xc7\x0b\x42\xf6\xd4\x8a\x4e\xa7\x4c\xee\x9d\x91\xbd\xe3\x66\x6b\xaf\x81\x6d\x01\x9f\x08\x68\xc0\x7e\xf8\xd2\x81\x0e\x19\xf6\xcb\x85\xd7\x5c\x48\xa1\x85\x19\x05\x3d\x4b\x26\x55\x20\x38\xf6\xb9\x1f\x09\x17\x9a\x28\xa6\xf7\x60\xc0\x77\x43\x4b\x79\x33\x36\x67\x0a\xc6\xfc\xc3\x4e\x9a\x69\xbd\x88\x09\xe0\xcf\x0a\xc7\xfe\x66\xc6\x7a\x82\xab\x28\x37\x98\x2e\x16\x61\xe0\x51\x0d\xa4\x0f\x7f\x57\xb0\x54\x32\x16\x18\xf1\x23\xef\x81\x63\xa9\x9e\xa9\x54\xa0\xc3\xe5\xd1\x61\x28\xa6\x87\x4c\x4a\x21\x93\x66\x1c\x27\x94\xce\x7c\x23\xf7\xd1\x7c\x4e\xe5\x1a\x25\xbc\x12\xd3\x36\x4e\x20\x81\x22\x94\x93\x9b\xde\x05\x61\xdc\x5f\x88\x80\x6b\x32\x81\x66\xa0\x38\x0d\xf8\x14\xfb\x0c\xe1\xa6\x13\xd2\xd0\x11\x0b\x26\x0d\x67\x1d\x3f\x4b\x2b\x3b\x44\x32\xb5\x00\xf9\x99\xca\x71\x00\x1d\xc7\xad\x56\xa1\x09\x1a\x7d\xa6\x3c\x19\x2c\xb4\x53\xff\x65\xe6\xb3\x91\x1f\x69\x36\x80\x6e\x50\x80\x9e\xff\x96\x6c\x82\x93\xff\xeb\xd0\x67\x93\x80\x07\x38\x5b\x1d\x9a\x1d\x1e\x47\x93\xf6\x7c\xa1\xd7\x7b\xb9\x39\xdf\x5f\x94\xfd\xfc\x3d\x23\xc4\x82\x4a\x3a\x67\x1a\x8c\x21\xd9\x17\xfb\x4f\x81\x7d\x0e\xa3\x70\xe9\xb1\xf0\xd7\x45\x86\x03\x5e\xd5\x23\xd9\x3f\xa3\x40\x32\x54\xa1\x96\x11\x7b\xa4\xa0\xb8\x61\x4c\xda\x7d\x78\x80\x98\xbf\x65\xc4\xd4\x74\x5a\x14\x10\xf7\x14\xcf\x51\x3a\xfe\x45\x96\x82\x53\x52\x6a\x7b\x4b\xc6\xf5\x4e\xb6\x87\x13\x1e\x62\x7b\x38\xee\x3e\xdb\x33\x8b\xff\x69\x7b\xff\x72\xdb\x33\xfb\xf0\xfc\xb6\x97\x38\xe3\x0c\x0b\xa9\x3b\x34\x27\xe0\x42\xf8\x2c\x6b\x8e\x7a\xbd\x30\x4a\x52\x5a\x82\x61\x25\xca\xd8\x63\x3c\x9a\xe7\x56\xdf\xbb\xbd\xfe\x78\xdd\xfd\x7c\x9d\x35\xa7\xfe\xc5\x4d\xa7\x37\x18\x5d\x75\xcf\x2f\x3b\xd7\x1f\xb2\x3d\x60\xb9\xa3\x8b\xf3\xab\xab\x6c\xdb\x5f\x06\x83\xde\x46\xe3\xf9\xe5\xa8\xdf\xbe\xf9\x7b\xfb\x26\xdb\x78\xd9\xfd\x34\xba\xee\x0e\x46\xef\xbb\xb7\xd7\x97\xb1\xb4\x89\x6e\x50\x3a\x1a\x85\x78\x8a\x36\x98\x4a\x83\x18\x0a\x4a\x8c\x1e\x98\x22\x7a\xc6\x88\x87\x0d\x62\x62\x7e\x66\xa9\x33\x88\x0f\xec\x07\x26\x3a\x8b\x8e\x5a\x94\x29\x47\x8c\x7f\x67\x5e\x7a\x90\x30\x26\xc1\x59\xd3\x41\xe1\x1c\xed\xd1\x48\x0b\x2e\xe6\x22\x52\x23\xb5\x56\x9a\xcd\x47\xa0\xc5\x31\x93\xc5\xd3\x16\xd3\x85\x63\xcd\x70\x3f\xb3\x46\xb6\x07\xe7\x7c\x4e\xb5\xeb\x3e\x7d\x95\xef\x2c\x1c\xc9\xaa\x05\x09\xf8\x8a\xf3\xfe\x35\xa9\x51\xe5\x74\xe0\x13\x08\xda\x9f\x83\x2f\xc1\x82\xf9\x01\xad\x1b\x07\x33\xe4\x9e\x08\x43\x90\x0c\x03\x3a\x68\x06\xc2\x32\x87\x2f\x18\xdb\x01\xc6\x24\x67\x9a\xf4\xf0\xa0\xc2\x28\x52\xeb\xf4\xea\x44\x8a\x48\xa3\xfb\x59\x80\xad\x07\x5f\x41\xaf\x11\xf7\x99\x1c\x72\xab\x5e\xae\x25\x0c\x04\x3a\x82\x33\x64\x60\x2e\x24\x23\x40\x64\x25\xe4\x17\x62\x9d\x93\x90\x0a\xf9\x18\xb3\x19\x0d\x27\x38\x94\x12\x05\x04\x43\x36\xe4\xd4\x9f\x83\xc9\x82\x19\x82\x0b\x5b\xc2\x0e\x71\xd8\xc9\x35\x92\xf1\xc5\x9c\x06\x1c\xb6\x8d\x6a\x5c\x58\x41\x0f\xf0\x0e\xeb\xcd\xe7\x82\x37\x88\x17\x32\x2a\xc3\xf5\x90\xc7\x72\x26\x4c\x0a\x40\x09\x6b\xa2\x85\xd9\xf1\x58\xa2\xe6\x5e\xa9\x4f\xd9\x54\xa4\x90\x53\xca\x83\x6f\xd4\xa9\xba\x74\xff\x0a\x87\xa6\xb0\x7d\x65\xbd\x89\x79\x6e\x5f\x10\xb7\x07\xb9\xb6\x6d\xca\xb6\xd5\x06\xdd\xcb\xee\x99\xdb\x4f\xe8\x0e\x54\xbd\x5c\x18\xe7\xf8\x9e\x88\xe5\x82\xc9\x21\x71\x22\x59\xb2\x13\x46\xb9\xfd\x1e\xc1\xf6\x0a\xe5\x3e\x8f\x2a\x0b\x7c\xe5\xd4\x57\xe0\xcf\xf0\xec\x8e\x7e\x56\xa5\x59\x7e\x5f\x14\xf8\x2e\xd2\x07\xd7\xb0\x21\x36\x34\x39\xaa\x91\x62\xb2\x99\x77\x28\x3d\x3a\x65\x7d\x03\x8a\x9f\xcc\xdf\xa2\x07\x2d\x7e\xf7\x77\xf3\x8f\x05\xa9\x2c\x83\x89\x9f\x84\xe3\x08\x41\xcf\xb3\x60\x03\xe5\xb2\xa8\xbe\x20\x99\x0d\x6a\xe7\xfe\x0d\x84\x4a\x96\x43\x33\x45\x8f\x59\x32\xed\x92\x2d\x03\xaf\x54\x25\x0f\xf5\xb2\x81\xff\x6c\x46\x14\xf8\xc4\xa3\xe8\x9f\x88\x88\x24\xb8\x18\xf1\x25\x60\x24\xf0\x1b\xa4\x73\xf9\xfe\xdc\x38\xd5\xf3\xce\x65\x85\x95\x3b\x26\x72\x8c\x6d\x03\x05\x56\x11\x03\x9c\xb5\x8d\x25\x1c\x50\x34\x3c\x5c\x2a\xb6\x3c\xc7\x24\x3a\x4a\x43\x10\xd8\x85\xe1\x9e\x90\x3e\xf8\x42\xe3\x37\xb1\x69\xc8\xc7\x2c\x14\xab\x5d\x0c\xde\xf2\x97\x5d\x9a\xc6\x4b\x74\x2e\x71\x3d\x1a\xaf\x0d\x9f\x3e\x5b\x00\x50\x45\xb7\x0b\xa7\x6f\x35\x63\x92\xd9\x38\x8b\xc0\x67\xc8\x57\x10\x86\x90\x04\x99\x48\x31\x2f\x35\xa6\x8c\x32\x9e\xea\xb0\x5c\xb6\xfb\x1f\x07\x5d\x40\x1c\xdd\xee\xc7\x4e\x3b\xdb\xf3\xa9\xfb\xae\x73\xd5\x2e\xe9\xc0\x5d\xce\x81\x13\xd8\xed\x47\x1d\xaf\x03\xe2\x46\x9c\xc5\x3f\xc4\x9e\xdd\x11\xc9\x68\x94\x72\x9f\xa8\x99\x88\x42\x1f\xed\x0f\x53\xe4\x00\x82\xab\xd9\xc4\x25\x0d\x03\xbf\x39\xe4\xe4\x80\xe4\xa5\x3a\x2b\x7c\xe7\x77\x2b\xdd\x1e\x13\x3c\x73\xbb\x60\x36\x53\x7d\xd1\x62\x01\xa6\x21\xc5\x0a\xfd\x97\x59\x20\xa7\x9c\xb3\xfc\xe7\x4e\xe4\xe7\x62\x1c\x84\x8c\x2c\x66\x00\x07\x0a\x6b\xa0\x9e\xcf\xec\x99\xca\x52\x84\xb0\xe7\x63\xdc\x9f\x04\x80\x60\xd0\x03\x51\x7f\x89\x0e\x00\xe6\x29\xa3\x1e\x04\x2d\x4b\x1a\x84\x74\x0c\x84\x05\x0f\x01\x1f\xf0\x21\x0f\xba\x7d\xa7\x46\x65\xc9\xe3\xb6\x9d\x99\xff\x3f\x05\xf9\x73\xee\x4b\x01\x87\x2a\x5e\xa2\x04\x6b\xde\x77\x46\x93\x4d\x2e\x33\xfd\x76\xb1\x3a\xb1\xab\x3f\xf4\xf2\x90\x7e\x8b\xe3\x49\x33\x80\x6d\x1e\xc7\x00\x65\x9b\x80\x06\x5c\x69\xca\x3d\x23\x48\x06\x50\xfb\x84\x8e\xc5\xb2\x2a\xdc\x7b\x91\xd2\x62\x5e\xe5\xa3\x0b\x52\x99\x3e\xea\xfb\x86\x43\x1a\xf6\xca\x65\xdc\xe6\x0d\x5c\xef\x86\x9f\xcf\x26\x58\x5b\xc5\x35\xec\x1a\x81\xc9\x9c\x2e\x86\x51\xab\x75\xe2\x59\x1a\x0d\x62\xff\x34\x6d\xcc\x9a\x79\xc0\xbd\x30\xf2\x31\x56\xf2\x35\xf1\xa9\xa6\xb6\x99\x92\x45\x34\x0e\x03\x35\x43\x3c\xbc\xc2\x3f\x15\x42\x4f\x37\x9a\xac\x02\x3d\xcb\x64\xed\xa5\x7a\x9b\x33\xa5\x00\x33\x3c\x5b\x70\x73\xf4\x63\x0f\x64\x72\x20\xe2\x1a\xab\x90\x1b\x1f\x01\x09\xc6\x2a\x61\xdb\x58\x08\xc0\xdf\xbc\x8a\xab\xd2\xee\x22\x70\x8b\xd7\x68\x90\x60\x62\x73\xee\x34\x76\x10\x63\xbb\x98\x38\x18\xff\xb2\x0c\x54\x60\x0f\xa7\x43\x27\x38\xaf\x82\xf7\x64\x3f\x76\x09\xca\xbd\x64\xd2\x36\x96\x13\xd2\x25\xc7\x24\xa1\x90\x9c\x15\x08\xc2\xc9\x04\x5b\x57\xad\xe0\x18\x88\x78\x5f\x46\x90\x01\x79\xcf\x67\x04\x99\x35\x90\x79\xe3\xe5\x52\xa7\x07\x1b\x10\xe8\x7d\x23\x13\xfb\xea\x31\x33\x2b\xf5\xee\x1e\x8d\xa6\xb3\x2a\xeb\xd5\x01\x98\x92\xa6\xf3\xc5\xc3\xd4\x1d\x57\x80\x06\xc9\xb4\x6d\x5c\x27\xc4\x63\xe3\x4d\x1b\x40\xe5\x80\x37\x38\x99\xca\x85\x77\x60\x37\x91\xcc\x80\x59\xc0\x40\xe0\x75\x41\xfd\x38\x7c\xbb\x95\x23\x78\xdf\xc5\x48\x6e\xd5\x3d\xf6\x81\x04\x4b\x4c\x03\xe7\x65\xad\xc2\x24\x0d\x1b\x06\x71\x1f\x2e\xb3\xc5\xe9\x0d\xc4\x9e\x84\x1f\x57\x1f\xdc\x37\xe7\x7b\xbf\xe1\x5c\x90\x69\xb5\x35\x10\x0c\x75\x34\x0c\x51\x31\x43\x2e\x59\x48\xb5\x61\xc7\x9a\x51\x9a\x15\x15\xa2\x54\xb1\x8e\xb9\x73\x6d\xc4\x1f\xc9\x8d\x7c\xe1\x5e\x55\xa7\x49\x46\x45\x1a\x9d\x50\x2d\xc6\xdf\x4c\x8f\x31\x60\x00\x54\x1e\x93\x5c\x99\xc4\x39\x31\x08\x32\xc0\xaf\x40\x0d\x79\xe1\x24\x20\x3e\x38\x69\xb5\x88\x0d\x07\x6f\x6d\x38\xb7\x1f\xe4\xe7\x56\xeb\xcf\xc8\xf7\x94\x91\xef\xcf\x30\x53\xca\xf2\x33\x85\x99\xff\x60\x4f\xbd\x6b\x8e\x6d\xbc\xd6\xbd\x29\xb6\x39\xdb\x9b\x2a\x36\xa8\x1e\x53\x4d\x92\xa9\xa2\x6e\x03\xbe\xff\x61\x71\xc4\x18\xfe\x03\xe2\x08\x8e\xdb\x2f\xba\x00\xeb\x10\x5d\x28\xb1\x15\xe0\x9d\x42\xc9\x93\xa6\xfa\xef\xba\xdd\x01\x5e\x42\xb4\x6f\xcc\x5d\x44\xfb\xb2\xe2\x92\xa2\xd0\xd1\xee\xf7\x3b\xdd\xeb\x51\xaf\x70\x77\xd1\x3b\xff\xd0\x1e\xfd\xbd\xd3\xfe\x9c\xbb\xd0\x68\xff\xed\xb6\xdd\x1f\x64\x9b\xde\x77\xae\xae\xf2\x14\x3b\x9f\x7a\x37\x96\x68\x6e\x9d\xc1\xf9\x4d\x6e\x62\xff\x63\x27\x57\xc6\xbb\xb8\xea\x5c\x7c\xcc\x36\xfc\xed\x16\x66\x40\xda\x7d\x54\xd6\x78\x5c\xd6\x78\x92\xa3\xd7\xfd\xd4\xbb\x6a\x0f\xda\x8f\x2a\x5d\x18\xeb\xcf\xde\x9f\x64\x13\xda\xac\x33\xcf\x6f\xae\xb9\x46\x79\x54\x36\x1b\xe8\xf5\xb3\xa1\x5f\x24\x1e\x3b\xa6\x49\x04\x18\xc8\x34\x6c\x29\x5d\x7b\x22\xe2\x5a\x3e\x23\x43\x96\xbe\x8d\xc1\x7a\x25\x88\x37\xa3\x80\xcd\x35\xc3\xea\xa3\xed\xb2\x98\x4d\x65\x0f\x7c\xa7\xdf\x25\x27\x47\xa7\xbf\x1c\x1c\x55\xb0\x1d\xa8\x07\x3a\xf7\xe4\xd6\x6b\x6b\x55\xb4\xa4\xf4\xdd\xef\x15\x78\xca\x38\x6f\x53\x18\xdf\xe6\x28\xe7\x4c\x4b\x31\x2a\xab\x5b\x3c\xd9\xcd\x58\xba\x44\x91\xf5\x1e\x16\xa4\xc8\xb9\x64\x16\x0e\x57\x85\x75\x01\xfe\x36\xdc\xca\xe3\xa3\x37\xdf\xae\x71\xe1\xca\x2d\xc8\x9b\x6d\xd9\xc6\x17\x86\xd3\xd1\x37\x90\xe0\x19\xee\xa4\xe2\x70\x6d\xe8\x17\xd5\x36\xf8\x95\xe4\x2e\x13\x87\xdc\x3c\xea\x39\x3b\x3c\x04\x68\xb4\x8a\x2f\x16\x9b\x42\x4e\x0f\xf1\xeb\xf0\x2a\x50\x7a\x24\x26\x23\xfd\x6d\x84\x58\x71\x4c\x15\x1b\x25\xc4\x55\xb9\x70\xcb\x57\xcf\xa6\xeb\xe5\xab\x0d\x13\xee\x41\x1b\xe0\x6f\x68\x53\x15\xca\x5e\x9e\x3e\x1f\x3f\xa7\x25\xfc\x9c\x96\xf1\x73\x5f\x1c\x37\x47\x38\x4b\x6c\xca\x04\xc0\x2b\xcf\x5d\x24\xa6\xc1\x18\x8f\x64\xa4\xcd\x5a\xb7\x9b\xd7\x55\x0e\x98\xe6\x0b\x50\x3b\x7b\x6f\x7b\x59\xfb\x6c\x5a\x73\x77\xc1\x71\x0d\xdd\x7e\xd5\x32\x19\x2e\x64\x26\x2a\x1a\x1f\xd8\x9e\x7a\x72\x8d\xb5\xa8\x86\x94\xa1\x80\x4c\x11\x0d\x73\x37\x34\x7c\x19\xd9\xb7\x36\x5b\xd9\x4d\x68\x27\x1c\xbb\x59\xc4\xa3\xa1\x17\x59\xb8\x34\x66\x7a\x05\xe9\x05\xd9\xe7\x74\x19\x4c\x4d\x77\x5f\x53\x09\x90\x0b\xd0\xd5\x90\xef\x23\x11\x03\x99\xda\xdc\xdf\xcf\x15\xa8\x0d\xc9\xfd\x55\xc0\x7d\xb1\x6a\xc2\x76\x18\x55\x02\x46\x6c\xc2\x92\xa0\x8d\x7d\xe2\x8a\xee\x43\xfe\x57\x98\xd6\x37\x9c\x11\xbb\x93\x55\x5e\x0f\x9f\xd6\x6c\x3c\x0e\xfa\xb7\x4d\x60\x2d\xbb\x0f\x4e\x60\x0b\x07\xee\x9f\x11\x83\xf0\x6a\x69\x54\xea\x43\xcf\x9e\xcf\xfd\x03\xf1\xc4\xf1\xe3\xcf\x0e\x5c\xa1\xb1\x92\x1a\xe6\xcd\xf1\x71\x75\x86\xee\x8c\x39\xcb\x77\xbd\x82\x71\xb0\x5d\x26\x65\xf5\xcb\x93\x47\x33\x1f\x2f\x10\x0b\x90\x7c\x3f\xe0\xc8\xa9\xe2\xad\xb8\x69\x2d\x3f\x6f\x99\x5b\xf4\xad\x65\x4d\x7b\x8f\x0d\xcc\x30\xd0\x1b\x26\x96\x10\x9c\x50\x61\x26\x48\xc5\xb7\x3b\x05\xb8\xe2\x2e\xbf\xb7\xa6\x74\x32\x7c\x36\x15\x02\xed\x1c\x16\xbd\xbd\xb9\xca\x14\xeb\xd2\xdb\xf7\x46\xc6\x06\x1a\xa9\xb5\xa0\x4c\xa6\x84\xb7\x69\xc1\xf7\x85\x0c\x54\xea\x66\xe6\xa7\x01\x7d\xda\xf7\x40\x71\x79\xb0\x24\x7a\x50\x62\x26\x87\xc1\x17\x66\x57\x07\xae\x53\xa6\x1a\x19\x5b\x00\x1c\xc3\xb4\xd7\x6c\x96\xc7\x99\x92\xaa\xc9\xce\xc1\xc6\xca\x34\x7a\xc6\xf7\x00\xc9\x0a\x85\x9b\x4f\x57\x38\xf3\xad\x5f\x31\xd2\x27\xf5\x18\x7c\x40\xa0\x18\xb6\x4a\x11\x4d\x67\x56\x4b\xbf\x83\x03\xb6\xc4\xc8\x58\x08\x8d\x3e\x1d\x22\x30\xe9\x18\x55\x83\x93\x2e\xd0\x80\xad\xf0\xf1\xdd\x12\xf0\xa7\xac\xae\xcf\x0f\xdf\x41\x96\xcd\x94\x36\xbe\x0d\xd4\x71\x60\xdc\x04\xc2\xdc\xc0\x53\xb6\x45\x05\x3a\xd3\x62\x94\x5f\x61\xd5\xd1\x33\xea\x0c\x89\x17\x5d\x2d\xb4\xc4\xce\x2d\x91\xd1\x56\x7e\x92\x77\x0a\xc6\x22\x3c\xba\xaf\x48\x0c\x17\x77\xb2\xe8\x84\x6c\x11\xed\x27\x1d\x35\xf0\x06\x76\xe3\x98\xac\x27\x25\x34\x8f\x96\xda\x67\x1f\x40\x58\xfe\xb5\xd2\x0f\x14\xbc\xc7\x80\xc2\xbe\x54\xfb\xe0\x27\x28\x6a\x26\x6b\x6c\x56\xbf\xe3\x0e\x48\x28\x74\xa4\x62\xf5\x2b\x2b\x57\x55\xda\xf8\x84\x56\x91\x64\x15\x9b\xd6\xe0\x98\x00\xab\xa8\xca\x5e\x47\x63\xb1\x71\x4d\xf0\x84\x6a\xb3\x0b\xc4\x75\x60\x38\x32\x94\xab\xe4\x19\x97\xa9\x00\xfb\x4c\xdb\x97\x90\x14\xcf\x3b\x0c\xde\x52\xf2\x1b\xc1\x41\xe4\x95\xec\x3e\xfe\x3c\x25\x4b\x14\xf5\x68\xb8\xb5\x3d\xf1\xdb\x8b\xf1\xda\xf4\xc4\xaf\x2d\xca\x99\x5e\x45\x72\xb2\x11\xd6\xb6\x15\x2a\x3f\x9b\x09\xdb\x98\x34\x24\x4b\x4a\x95\x66\x66\x36\xee\x9a\x81\xbb\x17\x2b\xdd\x79\xcc\x6a\xa0\x2c\x36\x19\x8d\xe4\x6d\x3c\x7f\xae\x6f\xd5\xe3\x42\x8e\x5f\x7c\xbd\x76\xaf\xee\xdc\x7b\xb7\xad\x59\x8d\x7b\xb8\xb5\xa1\x3d\xdc\x4a\xf7\x0c\xcb\xc5\xe8\x87\x94\xa1\xe3\xdc\x6f\x17\x2e\x6d\xd9\x6e\x7b\x2e\xe3\x25\x4f\x53\x4b\xd8\xb4\x59\xe8\x0e\x5c\x2e\x4a\x9e\x3a\x6c\xbd\x32\xc1\xf1\xdb\xe1\xf4\xb4\x4a\x87\x06\xb2\xec\xc2\x5b\x1c\x34\x46\xcf\xb2\xdf\x45\xf2\x4f\xb6\xf3\x6a\x23\x6a\xdd\xcb\x6e\x1c\xe8\xb6\xa2\xa0\xd8\x5f\x97\xb2\x19\x1f\xcd\xfb\xf8\xbc\xef\x88\x9b\x5b\x8d\x82\x87\x63\xdc\x3f\x30\x67\xda\xad\xb5\x62\x63\x04\x39\xa5\x47\xfb\x73\xc1\xab\x6d\x7d\x8b\xba\x91\xc6\x3f\xc2\x27\x70\xca\x45\x65\xce\x7c\x7f\x21\xf3\xe4\x78\xab\xf2\xfb\xc1\x14\xd5\x39\xc1\xaa\x30\xee\x9b\x7b\x35\x0f\xc8\x09\xef\x5e\x34\x31\xab\xdb\x2f\xd0\x9c\x08\xa3\xf8\x35\xbf\x89\xb4\x0b\xca\x87\x1c\x3f\x82\x39\x6b\x92\x58\x5a\xc0\x94\xb0\x69\x08\x5e\xb9\x79\xac\x1f\x53\x93\x99\xaa\x1f\xac\x69\xaf\x6e\x49\x6b\xc8\xef\xec\x08\x75\x47\x26\x01\x0b\x7d\x7b\xef\x83\xd5\xca\xc0\x3c\xd6\x07\x64\xc5\xd9\xd4\x3e\xdc\xbf\x33\x0c\xb9\x81\x4d\xf2\x1e\xdf\x2a\xc4\xcb\x1a\x56\x32\x0b\xba\x5f\x12\x68\x00\x2d\x2e\xf8\xc1\x37\x26\x05\x59\xd2\x30\x62\x49\x16\x99\xa3\x46\xe6\x00\xdf\xc8\x98\x59\x89\x50\x3c\x7c\xd7\xad\x40\x41\x18\xff\xcc\xf0\x3c\x9f\x4d\xf2\xc9\xce\xb0\x6f\x12\x0f\x5e\xbf\x7e\xdd\x70\xff\x01\x98\x16\xe4\x65\xa6\xc1\xa6\x5e\x2a\xd8\x72\xb8\x0c\xe9\x1f\x8b\xf3\xf7\x17\xac\xdd\x3e\xbb\x55\xb2\x1b\x48\x92\xfd\xcb\x4b\x73\x72\xf4\x73\xe3\xe7\x5f\x4e\x1b\xad\x56\x0b\xff\xb3\x12\x15\x1a\xcb\xa5\xba\xef\x28\x9e\x27\x96\x92\xcf\x7b\x94\xe1\xb1\x41\xf0\x97\x3f\xfc\x83\x90\xf1\x29\x18\x48\x96\xc5\xac\xf9\x0c\x39\xb5\xaf\x44\x23\xc0\x25\xd0\x1f\x4b\x86\xb6\x93\xb3\xe6\xa4\x23\x6b\xcc\xf8\x32\x25\xb6\x66\x93\x27\xe1\x33\x0d\x6e\x9f\x19\x33\x4b\x10\x8b\x7d\x1e\x05\x2e\x7c\x2a\x0d\x55\xf3\xc4\x63\xa1\x5d\xc6\x34\xdc\xf3\xe9\x7a\xb8\x07\xc6\x22\xe1\xe7\xb9\xe0\x7a\x36\xdc\x8b\x49\xc5\x77\x95\xa0\xb2\x41\x7a\x63\xcd\xd3\x7c\xce\x0f\x26\x90\xc4\x32\xf4\x73\xae\x40\x07\xfa\x5d\x65\x47\x1b\x43\x75\xc5\xa7\x44\x5d\xa6\xd0\xa0\xe3\x47\xe4\xd4\xc7\xe4\x06\xd6\x57\xd1\x58\x9b\x6b\x1d\x90\xcb\x3d\x8f\x4d\x08\x35\xc9\x0d\xe5\x2e\x74\x2d\xc0\xc7\x7c\x0d\xc0\x66\x58\xb8\x26\x2f\x0f\x8e\xec\x1e\xae\x19\x95\xf8\xb4\x75\xc8\xdb\x5f\x61\x42\xc8\xc8\xd1\x19\xb9\x10\xf3\x45\x04\xf9\x5e\xb2\xb4\xa1\x9b\x63\x51\x99\x37\x01\x8a\x45\xbe\xb0\x57\x0a\x48\x02\xf7\x3b\x15\x42\x61\xa5\x91\xbc\x25\x90\xa6\xbf\x29\xf6\x81\x62\x73\x3d\xc9\x4a\x49\x15\x33\xee\xb5\xfd\x71\x73\x33\xde\xd0\xb7\x48\x22\xf9\x3a\xb0\x8b\xc5\xdf\x6f\x0a\x93\xcc\xd6\xbb\x29\xf6\xe7\x78\x82\xf9\x4a\x56\x01\xd0\x5e\xdb\x58\xc9\x3d\xcc\xb1\xcf\x75\x8e\x4f\xed\xff\x8b\xb4\x5d\x29\xb0\x55\x27\x7f\x58\x52\x25\x2c\xbf\x7c\x4b\x8e\xde\x6c\xf4\x3a\x7e\xa0\xaf\x15\xff\xe3\x06\x7d\x27\x2c\x54\x2c\xc7\x94\xca\x73\xc5\x1e\xc2\x95\xb7\x9d\xab\x83\x2d\x5c\xbd\x2c\xe3\x2a\x6b\x2b\xc7\xa9\xad\xa4\x7b\x6b\x8c\x25\xfd\x7c\x99\x6e\xee\x0f\x98\x4c\xb5\x61\x54\x1b\x94\xed\xcb\xda\xc7\xdb\xbc\x7d\x00\x4f\x45\x3d\xbc\x49\x27\xc5\xd6\x92\xb1\x90\xec\x84\x4d\x93\x49\x27\x95\x28\x3c\x67\xa6\x59\x5d\xa7\xb3\x4a\xd5\x9c\x6e\x7e\x81\x3c\xcb\x8e\xae\x58\xe8\x65\xf9\x42\xa5\x56\x56\x0e\x5e\xec\x6f\xcf\x3e\x00\xb9\x6c\xe0\x39\x69\x90\xe5\x7b\x21\x52\xd6\xe4\xc2\x23\xef\xa8\xac\x4d\x85\x98\x86\xac\x19\xaf\xd1\x34\x8b\xd4\xc1\x61\xea\x08\x5f\xcf\x55\xf4\x67\x0d\x6f\x00\xee\xf3\xaf\xfd\xee\x75\x1a\x0c\x9c\x83\x02\x47\x78\x67\x86\xdf\x99\x42\x2d\xfe\x64\x07\x5a\x7e\xc9\xdd\x1f\xdf\xef\xca\x9e\xd7\x9f\x13\xc8\x71\x99\x0c\x3c\x37\x27\xbe\x66\x36\xce\x7a\x2d\x22\xe3\x6e\x25\x43\xa4\x88\x0e\x9d\x2e\xed\x73\x7d\x84\xba\x7c\x0a\x46\x61\xff\xba\x00\xf4\xbd\xb9\xf9\xc6\x3d\xae\xf1\x17\x7d\xce\x7b\x1d\xd5\x24\xe7\xf8\xc4\x01\x46\x86\x84\xb9\xa3\x13\x98\x97\x6c\x48\x36\xd0\x31\xc4\x70\x8f\x0a\x4d\x58\xb1\xdf\xf6\x57\xaa\x93\xf7\x11\xc0\x0b\xd0\xc3\x3a\xdc\x4c\x38\x00\x14\x43\xe6\xb3\xf2\x7d\x1c\x94\x3c\xc8\xfa\x37\x43\xa1\xd7\x00\xcf\x12\x88\xb7\x1b\x16\x6d\x92\x6b\x37\x71\xc8\x5d\xa7\x0b\x9d\x06\x5f\xa6\xc4\x0c\xba\x53\x3a\x08\x43\x32\xa3\xb0\x0c\xcf\xae\x69\x0f\x87\x9d\x88\xcf\x8d\xcc\x8b\xce\xc8\xfe\x5a\xfc\x8a\x4a\x5b\xd3\xdf\x04\x49\x2d\xdc\xc0\x1c\xec\xfb\x17\xc3\xbc\x9b\x14\x4f\x65\xa0\xde\xed\xe0\xc2\xe2\x27\x15\x20\xe2\xb8\xe5\xc1\x57\xc2\x16\xc2\x9b\x0d\xf9\xd1\xeb\x5f\x5a\x07\xad\x23\xf8\x77\xd0\x6a\x9d\x99\x7f\x7f\x2d\x08\x69\x25\x6d\x99\x41\xb9\x71\x20\xfc\x90\x83\xd4\xaf\x0f\x8e\x8e\x01\x2e\x0e\x8e\x4f\xce\x7e\x7e\x0d\xff\xfe\xfa\xa3\xb0\x30\xf5\xe5\x39\x5c\x68\xff\x82\x02\xb7\x03\x65\x30\xcd\xb4\x7f\x33\xbf\x08\x84\x55\x58\x07\xda\x1a\xb9\x6c\x83\xaa\x87\x40\xc4\x21\x2f\x4f\x78\x60\x71\x54\x62\x1b\x95\xe6\x0c\xc1\x22\x3d\x80\x70\x02\x51\x58\xa4\xdc\xad\xca\x90\xf7\xa4\x08\xf1\x6d\xbb\x47\x3e\x48\x36\x15\x32\x80\x23\x7b\x11\x03\xc9\xd5\x2c\x00\x0a\xec\xab\x66\xb8\xa0\x29\x6d\x24\x83\x62\xc6\x87\x7c\x4c\xbd\x2f\x68\x76\xc6\x41\x20\x3e\xc3\xb4\xa6\xb8\x24\x55\x2a\x9a\x9b\x6b\x69\x30\x69\xf8\x01\x42\xb0\x32\x49\xd6\x69\x2b\x3e\x0a\x80\x55\x05\x16\xf4\x83\x26\xcc\x0e\x19\x5d\xa4\xa2\x4a\x84\xb0\x6a\x0e\xb4\xc1\x73\xed\x11\x25\xac\xbf\xe3\x22\x3b\x0e\x8e\x82\x7d\xb8\xaf\x20\x03\x63\xb8\xe8\xc4\xf8\x1b\xcd\x24\x68\xd6\x3a\xde\x0c\xc0\x44\x4b\x19\xf2\x0a\x53\x21\x25\x96\xd2\x7c\x1d\xff\xf3\x2b\x40\x81\x77\x6b\x54\x39\xde\x2f\x98\xd7\x81\xda\xb1\x24\x91\x7c\x03\x92\x73\x90\x5c\x45\xd2\x39\x66\xf8\x44\xbf\x0c\x5c\xe2\x6f\x59\x19\x53\xb4\xfb\x0a\xc6\x4a\x6e\xde\x5f\x90\x93\x93\x93\xd7\x58\xec\x67\xee\xaa\x16\x21\x6e\x9f\x31\xf2\x8f\xf8\x61\xc9\x6a\xb5\x6a\x06\x4c\x4f\xcc\xa3\x12\x39\xf1\xf0\x3f\x9c\xd4\xd4\x5f\xf5\x6f\xb5\x87\x8c\xaa\x57\xa1\xe6\x02\x12\xea\x75\xfb\x9d\xff\x25\x77\x68\x37\xb5\xfa\x5d\x09\xe8\x49\x5e\xb1\xba\x40\x97\x7c\x43\x38\xd7\x23\xb7\x67\x35\x33\xff\xfa\xf6\xea\xaa\x5e\x2f\x1d\x67\x4c\xb7\xd6\xaa\xbf\x79\x18\x3e\x73\x5c\x4d\x99\x46\x32\x62\x02\xf9\x4c\x96\xbb\xf8\xda\x0e\xfa\xc0\x2f\x12\xbd\x74\x6b\xe6\xc6\x5b\xd0\xa9\x97\x0d\x62\xf8\x7a\xf3\xc3\xa2\x2d\x9b\x7a\x89\x5f\xdb\x24\xb3\x83\x20\x5e\x7a\xe4\x27\x83\x62\xf2\x92\x9e\x54\x4a\xfa\x39\xe0\x27\xc7\xe4\xee\x03\xd3\x7d\xf3\xcb\xf0\xd8\x7d\xae\xde\x07\x21\x1b\x14\xb6\xe4\x7d\xe7\xaa\x3d\xe8\x7c\x6a\x93\x89\x76\x8c\x54\x4d\xb2\x92\x4f\x74\xcc\xf0\x6d\xe7\x7a\x70\xfa\x0a\xf8\xf6\xbe\x20\x72\xac\xd5\x6a\xb6\xa5\x3e\xd1\x4d\x7f\xf5\x97\x60\x3a\xbb\x04\x5b\xc4\xa9\x75\x87\x11\x1d\x52\x3c\x39\xae\x93\xff\x23\x66\xd4\x95\x58\xc5\x83\x12\x4d\x1e\x1e\x02\x72\xf8\x6c\x1e\x61\x28\x43\x1d\x0f\x1a\x08\x9f\x89\x86\x00\x2e\xe2\x01\xc6\xab\x93\xa3\xd3\xcd\x13\x98\x50\xc3\xe9\x47\xa7\xaf\x5e\xbd\xfa\xe5\xe4\x14\xc8\xc4\x0e\x61\xcc\x26\xf8\x17\x2a\xa4\xb1\x81\x94\x86\x86\x1f\xdb\xde\x9a\x55\x05\x68\xc5\xea\xe7\x30\x01\xa1\x75\xc8\xc3\x32\xec\xdc\x63\xdb\x48\x07\xf5\x15\xd3\xf9\x9f\x0c\x1d\x63\x12\xf5\xbc\x49\xbc\xaa\x34\x09\x7c\xb2\x42\xee\xec\xce\x36\xbd\x48\x4a\x7c\x34\x0c\x43\x3e\x01\x3e\x08\x54\xd6\x24\xd0\x83\x82\x7b\xc5\x66\xd8\xd7\xea\x19\xdb\x6c\x1f\x26\xa6\x09\x39\x67\xab\x77\x51\x10\xfa\x4c\xd6\xea\x28\x5b\xdf\x29\xc9\xad\x61\x75\x53\x8f\x81\x33\x21\x38\xe6\xda\x8a\x0f\x9e\x17\x85\x77\x23\xad\xf4\x4e\x72\xa3\x84\x7a\x73\x8c\x94\x1d\x33\xa9\x1e\x7e\xae\xd4\x83\x13\x24\x0e\xad\xa4\xb7\x06\x5c\xc9\x13\xe1\xb9\x58\x01\xef\x26\xe0\x59\xef\xe5\x7c\x43\x92\x54\x01\x47\x35\x18\xe4\xda\xe3\xbc\x09\x5b\xb1\x19\x93\x6c\x3b\xd4\x72\xf9\xd3\x4f\xaf\xeb\x85\xdd\xcd\xaa\xa6\xe6\x06\xbf\x75\x7f\x36\x2c\xc1\xb7\xe6\xff\xee\x6f\xa3\x30\x7f\x0b\xcd\x8b\xef\x2f\xfe\x3f\x00\x00\xff\xff\x8a\x3f\x67\x65\x76\x4c\x00\x00")

func rpcSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_rpcSwaggerJson,
		"rpc.swagger.json",
	)
}

func rpcSwaggerJson() (*asset, error) {
	bytes, err := rpcSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rpc.swagger.json", size: 19574, mode: os.FileMode(420), modTime: time.Unix(1466624904, 0)}
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
	"rpc.swagger.json": rpcSwaggerJson,
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
	"rpc.swagger.json": &bintree{rpcSwaggerJson, map[string]*bintree{}},
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
