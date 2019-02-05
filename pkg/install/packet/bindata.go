// Code generated by go-bindata.
// sources:
// lokomotive-kubernetes-packet.tar.gz
// DO NOT EDIT!

package packet

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

var _lokomotiveKubernetesPacketTarGz = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xbc\x13\x43\xec\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\x03\xed\x3d\xfd\x57\xdb\x48\x92\xf9\x99\xbf\xa2\xc7\x30\x0b\x99\x87\x64\x09\xcc\x97\x6f\x3d\x77\x04\x3c\x13\x5e\x08\xc9\x01\xb9\xbc\xdb\x5c\x56\x4f\x48\x6d\x5b\x83\x2c\x79\xd4\x92\xc1\x9b\x65\xff\xf6\xad\xea\x6e\x7d\x4b\xb6\x4c\x08\xb3\xb9\x75\xef\xdb\x09\x48\xf5\xd5\xdd\x55\x5d\x55\xdd\xad\xc2\xf5\x6f\xfd\xb1\x1f\x3a\x53\xaa\xdc\x46\x37\x34\xf0\x68\x48\x59\x7b\x62\x5a\xb7\x34\x6c\x0f\x5c\x33\xb4\xcc\x40\x71\x1d\x2f\xba\x6f\x67\xde\xbf\x58\xaa\x69\xd0\x0e\xf6\xf6\xf0\x5f\xfd\x60\x4f\xcb\xfe\x1b\xb7\x17\xfa\x6e\x67\x67\x5f\xdb\xdf\x81\x37\x2f\xb4\x9d\xdd\x83\x83\x83\x17\x64\x6f\x39\x36\x8f\x6b\x11\x0b\xcd\x80\x90\x17\xbf\xf9\x23\xd3\xf3\x28\xab\x83\x5b\xf4\xfe\x3b\x6d\xee\xa3\xe6\x9f\xb1\x91\x1a\x0e\x9a\xf2\xc0\x09\xde\xef\x74\xea\xe6\xbf\x73\x50\x9a\xff\x3d\x7d\x57\x7b\x41\xb4\x6f\xd9\xf1\xb8\xfd\x9b\xcf\xff\x3a\xb9\xa2\x56\x14\x50\x62\xf9\x93\x19\xa1\xa1\x65\x93\xeb\xf3\x2b\x62\x32\x46\x43\x46\x42\x1f\x9e\x7b\x61\xe0\xbb\x2e\x0d\x98\xba\x16\x50\xe6\x47\x81\x45\x49\xcb\x8b\x5c\xd7\x88\x7f\x6d\x91\x16\xa2\x2b\x29\xac\xc2\xa8\x15\x00\x85\x16\xf9\xb2\x46\x80\x46\xe4\x85\xa4\x47\x5a\x1b\x5f\xa6\x66\xa0\xa6\x60\x06\x7f\xf3\xd0\x5a\xe3\x40\x30\xbc\x56\xe8\xf8\x1e\xc7\x21\x24\x9c\x4d\x28\xfe\x0b\x78\xa0\x6f\x2d\xfe\x6c\xe4\xb3\x50\x3e\xdb\xf8\x42\x5d\x3a\xa6\x5e\xb8\x25\xb4\xd5\xb0\xe9\xd4\xb1\xa8\x9a\x15\xf8\x27\xd5\xb4\x2c\xca\x98\x31\x89\x6e\x5c\xc7\x32\x9c\xc9\xb4\xb3\x2d\xc4\x51\x1d\xcf\xa6\xf7\x2f\x1f\x04\xdd\x88\xd1\x40\xd2\xb5\xfc\x80\x8a\x87\xa1\x33\xa6\x7e\xc4\x05\xd7\xf7\xc6\xf8\xec\x01\x05\x9d\x04\xfe\xd4\x61\x20\x27\xa0\xb4\x06\x8e\x4b\x5b\x52\x60\xe4\x0c\xf2\x90\x44\xc0\xb1\x6f\x47\x2e\x55\x6f\x7c\x3f\x44\xe3\x51\x71\x78\x0d\xcb\x34\x2c\x1a\x84\x92\xb1\x4d\x59\xe8\x78\x26\xef\x36\xe2\xbc\x7e\xf7\xb6\xdf\x46\x38\xc5\x72\x1d\x20\xa6\x58\xa6\x6a\x05\xe1\x93\x31\xe7\x44\x97\x11\xe0\x1b\x70\xbf\xa5\xb3\xa6\xcc\x01\xf4\xa9\x98\xc3\x0c\x4f\x51\xe3\x9a\x74\x5d\xc0\x3e\x65\xd7\x25\xf7\x26\x5d\x97\xcc\x9f\xb0\xeb\x13\xda\xb4\xe3\x08\xf9\x94\xdd\xe6\x9c\x9b\x74\x9a\x33\xae\xed\x72\x40\xc1\x47\x51\x85\xde\x53\x2b\xe6\xef\x78\xe0\x95\x28\x50\xf9\xc4\x7f\x25\xb0\x48\x44\xb6\x4f\xc6\xb7\xb6\x13\x10\x65\x42\x90\x2c\xf8\x29\x97\x93\xe7\xff\x69\x6d\xe7\x21\xa7\x24\xa3\x69\x3f\x15\x10\x0a\xc0\x56\x15\xc1\xbc\x89\x56\x00\xb4\xc5\x54\xc6\x36\x5c\xcd\x3f\xd5\xb5\x7a\x0a\x8b\xd1\x61\xe8\xe6\xa0\xe3\xc0\x3e\x49\x87\x70\x9a\xe6\x77\x27\xd6\xa0\x3a\xec\x45\xa8\x35\x1d\x49\xf4\xa3\xd0\x8d\x91\x7f\xe7\x11\xe5\x92\x13\xe8\x72\xf7\x95\x43\x2e\x81\x83\x96\x22\xf8\x9e\xa6\x55\x02\x7e\xe6\xea\x07\x0a\x98\xf7\x8b\xb1\x52\x67\xfc\xe2\xbb\x8b\x7e\xc6\x37\x12\xd3\xb3\x09\x46\x12\x61\x0a\x0b\x40\x13\x1a\x0c\xfc\x60\x0c\xd4\x40\x8f\x15\xf4\x27\x84\x51\x77\xa0\xa0\x1b\xa3\x36\xb1\x5c\x88\x3e\x00\x19\x51\x58\x18\x98\x93\x89\xe3\x0d\xe7\x79\xd9\x98\xb6\xc2\x59\x09\x53\xb0\xe9\x84\x7a\x36\x33\xb8\x49\x09\x63\x68\x15\x6c\x51\x76\xae\x65\xde\x31\x23\x00\x8f\x46\xf7\x76\x81\x2a\x38\x3a\x5b\x35\x27\x8e\xd0\x11\x16\x03\xe5\x78\xaa\x75\x7e\x1d\x81\x3f\x3f\xce\x6d\xd7\xbb\x6b\xad\xc2\x5d\x3f\xa5\x83\x96\xc3\x9a\xae\x59\x18\x8d\xf0\x29\x35\x60\xd5\x98\xb7\x46\x89\x79\xff\xea\xc5\x69\x4a\xb2\xe4\x48\xdb\x9f\x84\xed\xc2\x24\xc5\xb0\x6c\x06\xaa\x31\xb6\x42\xb7\xa0\x56\x05\x45\xfd\xda\xf8\xef\x71\xf1\x3f\x0c\xfa\x24\x0a\x59\xc3\x1c\x60\x41\xfc\x0f\xe9\x9e\x5e\x88\xff\xf7\xb5\x3d\x7d\x15\xff\x3f\x47\x13\x33\x49\x5a\x38\xb9\x60\x8a\x03\x67\xa8\x98\xf6\xd8\xf1\x84\x26\x4f\x4d\x37\xa2\xd5\xee\xbd\x88\x00\xd6\x03\xa6\x51\x26\xb7\x0c\x21\xfc\xd1\xa5\x21\x27\xf5\x47\x0f\xcc\xbf\x49\x7b\x9c\xfd\xdf\xf9\xc1\x2d\x2e\xd8\x4f\x62\xff\x3b\xda\xc1\x6e\xc9\xfe\x57\xf9\xff\xf3\xb4\x34\xd6\xc8\xb9\x65\x88\x35\xc4\x1c\x1b\x9e\x0f\x0e\x31\x9b\xc6\x67\x5a\xe2\x43\x25\x6c\x9c\xcd\x0b\x87\xef\x99\x63\x5a\x06\x95\x31\x8f\x81\x6f\x1f\x14\x81\xa8\x6c\x7c\xc9\xe4\xe4\x9c\xc0\xc4\x35\x3d\x32\x8f\x17\x86\x19\x1c\x72\x60\x5a\x8e\xeb\x84\xb3\x32\x64\xfc\x86\x83\xf9\x10\x8b\x81\x53\xf7\x86\x86\xf0\xac\x3c\x92\x00\x51\xfc\x31\x84\x19\xf7\x3c\xa0\xb8\x71\x5c\x17\x01\xac\x99\xe5\xc6\x21\xcc\x08\x46\xc7\xe5\x59\x09\xf8\xfc\xdf\x20\xd0\x31\x1c\xbb\xc8\x27\x7d\xc3\x39\x21\x39\x83\x59\x81\x33\x09\x0d\x40\xce\x00\x16\xde\x70\x68\xd3\xbd\x33\x67\x10\xf1\xdc\xd3\x94\xec\xc0\x74\x19\x97\x08\xa3\x1e\xc3\x36\x43\x33\xd7\xb7\x78\xdf\x03\x5f\xa8\xc0\x57\xac\x9e\x72\x60\x14\x67\xe8\x39\x18\xbc\xe0\xee\x47\x00\x91\x21\x0d\xa8\x5d\xda\xf3\x80\xa5\x9a\x93\x6d\x25\xe8\xc9\x8c\xa7\x04\xf2\xb3\x5e\x3f\xd9\x71\x2a\x58\x16\x0e\xc6\x19\x26\x32\xa4\x06\x46\x61\xb1\x80\x82\x5f\x33\xf1\x72\x04\x52\x11\x25\x89\x82\x5a\xd6\x4b\x18\x93\x11\x20\x48\x6b\x8b\x07\xa2\xe1\x48\x15\xee\xe8\xa1\x6d\xb9\x72\x55\x53\x67\xe6\xd8\x55\x43\x40\x68\xbd\x14\x3b\x53\x40\x92\xc9\xc0\x2e\x75\x56\x25\xdd\x44\xe1\xa1\xdf\xba\xb6\x4d\x16\xbb\xb8\x78\xdf\x09\x22\x64\x4c\x8f\x59\x4e\xd5\x05\xbd\xdf\x18\x84\x97\x9e\x05\xe6\xb7\x25\xbb\x15\x03\x3f\xb4\x62\xf4\xdb\x43\x66\xd8\x1e\xe3\xbb\x0b\x60\xb5\xa0\xc9\x09\xba\xe5\xd8\x01\x5a\xe1\x16\xc7\x94\xef\xf1\xe1\x36\xd1\xb5\x18\x3f\x36\x46\xdb\x1f\x9b\x8e\x67\xb0\x68\x30\x70\xee\xcb\xb6\x9a\x7b\xfd\xd0\x7a\xa2\xd0\x73\xd5\xfe\x05\xda\xe3\xfc\x7f\x36\x69\x5b\x1c\x03\x2c\xf0\xff\x7b\x7a\xa7\xb8\xff\xbf\x8f\x47\x02\x2b\xff\xff\x0c\x6d\x9d\x9c\x3a\x0c\x53\x7a\x4a\x4e\x2f\xae\x88\xd8\x15\x60\x64\xe0\x07\x84\x9a\xd6\x28\xb3\xc5\xb1\xc9\xc0\x05\x3a\x53\x5c\x45\xcf\xde\x4f\x3b\x02\x04\x37\x5c\x22\x66\x0e\x69\x66\xd3\xa2\xbc\xcb\x00\x0b\x37\x42\x2e\x73\x1a\xb0\xce\xc5\xf9\x0b\x2c\x81\xe4\x6e\x04\x3e\x42\x4a\x46\x18\x38\x64\xd7\x26\x37\x94\x80\xd0\x20\x8b\x0d\xb0\x7f\x03\x28\x74\xcb\x09\x49\x5c\x12\xe5\x43\x41\x8d\x87\x22\x62\xed\xf7\x83\xb1\x19\x6e\xb5\x7e\x64\x0a\x8a\xf4\xa3\xad\xfe\xc8\xd4\xd6\x36\x29\x86\x26\x39\x97\x24\x5e\xc7\x54\xc5\xe2\xc9\xb7\x39\x80\xe4\x31\xff\x25\xe4\x6e\x7e\x57\xd3\x84\xec\xb9\x81\x32\x6d\x1b\x06\x87\x25\x03\x06\x10\xf1\x30\xf7\xc8\xa7\x25\xcf\x31\x04\xe1\xca\x83\x8c\xcf\x62\xdf\x2a\x9d\x46\xce\x31\x1c\x51\x72\xfc\xfe\x8c\xc8\x5d\x9e\x45\xf3\x94\xd9\x0f\xe2\x93\xf5\xb8\xb1\xad\x1d\xd4\x66\xe3\xd8\x82\x81\x94\x5a\x70\xfd\xee\xf4\x1d\x51\x08\xb8\x4e\xdc\x8c\xc3\xcd\x1e\x93\x40\xc0\x84\x9b\x6b\x62\x48\x8a\x1d\x44\x75\xf1\xf0\xe1\x8c\xd8\xbe\xb7\x19\x92\x91\x39\xa5\x80\x24\xb6\x94\xf8\x8c\x64\x09\x03\x8a\x33\x00\x72\x23\x13\x29\x8f\x23\x37\x74\x32\xdb\x5d\x40\x34\x8c\x26\xc0\x29\x8c\x4c\xd7\x9d\x11\x8c\x10\x58\x69\xfe\x96\x3b\x7f\x92\x13\x55\x1f\x78\x67\x08\x2c\x8a\xbb\x2b\x6c\xa7\x79\xec\x9d\xd9\xd5\x5b\x26\xfe\xce\xf0\x5c\xc5\xe0\xc5\x18\x3c\x33\xa6\x8f\x8f\xc3\xab\x88\xd4\xc4\xe2\x95\x0a\xd0\x34\x1e\xcf\xf0\xf9\x9a\x98\xbc\x4c\xa6\x2e\x2e\xaf\x94\xb6\x59\x6c\x9e\xa2\xce\x8d\xcf\xd7\xc9\x09\x78\x69\x3f\xc4\x59\x23\xa8\x43\x60\xf2\x62\x83\x1c\x42\x69\x87\x32\xe2\x7b\xd9\x13\x6d\x22\x16\x48\x27\xc8\x3a\x3f\x4e\x88\x9f\x5e\x49\x2b\xea\x09\xf7\x55\x36\x12\x09\x26\xc2\xe3\x3a\x33\x2b\xa3\xaa\xf9\x95\x54\x74\x02\x65\x47\x50\xad\x37\x0a\xc3\x09\xeb\xb6\xdb\x92\x0c\x27\xa0\xa9\xf4\xde\x84\x2e\xe3\xa4\x8d\xb7\xf1\x89\x5e\x09\xa7\xe7\xe0\x54\x55\x4d\xa5\x74\x50\x93\x4c\xd7\x88\x0f\x1d\x44\x92\xe1\x3b\xde\x56\x6b\x1b\xd6\xea\x0a\xfd\xe0\x3e\x3b\xa3\x12\x2f\x63\x49\x57\x59\x50\x9a\x05\xd5\x18\x45\x29\xde\x59\xca\x0e\xb8\xce\x88\xc1\xbc\x7f\x48\x26\x1a\x9e\x54\xa8\x56\x0c\xa6\xc2\x4f\x89\x4a\x75\x77\x76\x0f\xb5\xa2\x71\x70\xb8\xdc\xe8\x96\x15\x3a\xcb\xa0\x46\xa3\xe5\x31\x89\x64\x95\x5f\x86\x33\x3a\xbd\x4a\x12\xbf\x87\xf6\xc8\xfc\xcf\x5d\xe6\x0a\xe0\xf2\xf7\xff\x3a\xbb\x07\x7b\xab\xfb\x7f\xcf\xd1\x1e\x3d\xff\xc5\xcd\xb2\x39\x3c\xe6\xe7\xff\x3a\xcc\xfa\x41\x61\xfe\x41\x5b\x76\x57\xf9\xff\x73\x34\x45\x51\xd6\x44\x30\x6e\x77\x31\xd8\x85\x20\x81\x75\xf9\x02\xaf\xf0\xcc\xae\x0b\x59\x94\x85\x13\x2d\x9d\xa7\x3c\xa3\xa6\x9e\x79\xe3\xc2\xcb\x30\x88\x68\x0e\xda\x05\x68\x36\x76\xc2\x91\x5d\xc0\x18\x9b\xec\xb6\x02\xfe\xce\x74\x42\x05\x12\x47\x05\x3c\xc7\x22\x1e\x49\x64\xcb\xba\xe4\xef\x6b\xb1\x23\xfb\xf4\x01\x64\xfe\x9c\xfc\x7a\x4a\x45\x70\x0f\x41\x73\xef\x23\x10\xe7\x49\x30\x46\x77\x80\x17\x40\x04\x98\x00\x7e\x34\x81\x50\x4f\xf6\x5d\xc1\x74\xcc\x9d\xd2\xa2\xd4\x84\xbc\xa2\x40\x80\xf6\x64\xc4\x52\x7a\xfd\xe9\x4a\x3c\x48\x05\xb8\x86\xa4\xa8\x07\x2e\x90\x8d\xfc\x30\x79\x78\x49\x31\x76\x38\x1e\x80\x13\xed\xdf\x3b\x61\x2f\xd3\x27\x42\xfa\xf7\xd4\xba\xc2\x43\xfe\x5e\xfb\xc6\xf1\xda\x6c\x44\x14\x8b\x6c\xde\x8d\x20\x92\x20\x3f\x90\x76\xc4\x02\xfe\x7c\x18\xd0\x09\xd9\xfc\xeb\xa7\xbf\xae\x7f\xea\x32\xb0\x50\xda\xfd\xfc\x79\x53\xdc\x62\x11\xd2\x63\x5c\x31\x20\x3f\x93\x36\xa4\x92\x6d\xbc\xc6\xf1\x1f\x04\x2f\x12\xb8\x14\xf0\x74\xfc\xd9\xa3\x9b\xa9\xe0\x67\x1e\x68\xbe\xeb\x7e\xce\xc8\xf8\x7b\xe4\x40\x98\xf7\x6a\x56\xd9\xdb\x78\xc6\xaa\x47\xe2\xeb\x27\xeb\x8d\xa0\x4b\xa6\x8e\x49\x5e\xc3\x18\x06\xc8\xa8\x30\x5b\xc1\xc4\xc2\xab\x2f\x61\x79\x9a\xca\xf3\xd0\xf7\xa6\x4e\xe0\x7b\x98\xfc\xfc\x02\x23\xd9\xe3\x03\x95\x59\xc3\xe2\x7e\x50\x6f\x5a\x85\xd3\x6b\x5d\xbe\xb9\x36\x2e\x3f\x5c\x18\xc7\x97\xbf\x5e\xf5\x14\x25\x8a\x1c\x5b\xc1\xe8\x4e\x61\xe6\x14\xc8\x41\xb8\xd3\xb6\x4c\x6b\x44\x63\x4a\xca\xc4\xb7\x55\x84\x22\xff\x97\x10\x84\x61\x53\xa6\xbe\x1b\x8d\x69\x4f\xcc\xd1\xf6\x2d\x44\x5b\x3d\x0c\x49\xb7\x45\xfe\xdf\x2b\x4d\x60\x1e\x7b\xcc\xa3\xc7\x3c\x0d\xd0\x95\x21\x0d\x17\x61\x0a\x1c\x0c\x00\x61\x01\xbf\x51\x2c\xcf\xa9\x60\x8e\xbd\x80\xb7\x6d\x78\x3b\x8f\x71\x96\x48\xcc\xbd\x1e\xb5\xc8\xd9\x84\x34\xcc\x9f\xc7\x9c\x03\x34\xe2\x2f\x48\x95\x44\xa8\x22\x20\xa5\xf0\x27\x21\x8a\xad\x80\x09\x55\x88\x80\x37\x70\xe0\x2d\x1a\xd8\x3c\xfe\x59\x22\x31\xf3\x7a\xd4\x6c\xff\xfd\x61\x5d\xc7\xfd\xe1\xc2\x1e\x03\x72\xae\xab\x25\x14\xc7\x63\xfc\x9a\x9a\xe2\x73\x23\x62\x3d\x67\x6c\x0e\xe5\xed\x28\xae\xd0\xf1\xd2\xf2\x1e\xd6\x30\xbe\x8a\xa4\x17\x24\x33\xf2\x37\x82\x2f\xd8\xcf\xd8\xf4\x9c\x01\x65\x21\x7b\x0c\x32\xf2\x85\x9f\x54\xbb\x11\x72\x46\xd3\x96\x83\xe7\x6a\xb1\x14\x8a\xb4\xe5\xb6\x98\x85\x89\x1b\x0d\x61\x88\x6b\x28\xc4\x0b\xf3\x8d\x29\x96\xec\x96\x58\xa1\xf1\xa2\xad\x33\x70\x2c\xc8\xe5\x14\x33\x0a\x47\x7e\xe0\x84\x33\x05\xd3\xc4\xcd\xd2\x38\x64\x32\xe9\xbf\x13\xf3\xee\x96\x6c\x7e\x99\x04\x0e\x68\xc1\xc6\xce\xc3\x26\x3c\x02\xd2\x74\xbf\x43\x14\x1b\x97\xf5\xe2\x18\x26\x37\xc3\x2b\x84\x53\x12\xe9\x82\xdb\x90\x04\x63\x92\x59\xc0\xe6\xad\x5d\x55\x4e\x09\x29\xf1\xe1\x04\x47\xe8\x27\x0b\xa7\x72\x87\x17\x16\x69\x50\xd0\x48\xd3\xf3\xbd\x19\x68\x32\xe3\x7d\xef\xf1\x6d\xb4\x22\x08\xbc\x80\x15\x16\x47\x08\x74\x56\x09\xfd\x5b\xea\x29\x77\xf4\x66\xe4\xfb\xb7\x15\xa0\x30\x7e\x7f\x13\x90\x63\xdf\xa6\xbd\x8f\x95\x80\xc9\xcd\x55\xd9\xc3\xca\xb1\x2a\xe1\xc8\xc4\x1e\xec\x66\xe3\x4b\x79\x17\xe1\xa1\x0e\x9e\x6f\x04\xf4\xd2\x6c\x3c\xbf\x33\x50\xc4\xf2\xf8\x8e\xf2\x40\x01\x55\x2b\xcb\x15\xdb\x41\x01\x89\x42\x90\xa0\x40\x8f\x31\x9c\x52\xa4\x23\xc5\x0b\x89\x79\xa8\x54\x7b\x2a\xbd\x9b\x54\xac\x3c\x0e\xa7\x98\x6a\x41\x10\x79\x6d\x7c\x94\xb8\x43\xfc\xa5\x80\x02\x14\x31\xda\x57\x84\x35\xf4\xca\xeb\x3d\xde\x11\x51\x5c\x13\x08\xb0\x1e\xff\x39\xf0\x5d\xb1\xdf\x23\xa4\x51\x1d\xbf\x8d\xcf\x0b\x68\xa0\x73\x4a\xbc\x8a\x28\xb8\xe1\x57\xea\x45\xb2\xc6\x14\x30\x03\x6a\xda\x30\x3c\xee\x0c\x68\x80\x86\x6a\x95\xeb\xaf\x94\x57\x8c\xfb\x32\xf6\xed\x4f\x0a\xe6\xc3\xe0\xd1\xd2\x06\x74\x49\xf9\xc5\xcd\x9e\xd8\x5c\x2e\x3e\xbe\xa2\x56\x6f\x6f\x4e\x30\x86\xf1\x0e\x0f\xc5\xc4\x91\x04\x6e\x42\xab\xc2\x15\xe4\x83\x73\xe4\x4d\xf9\x04\x7c\x93\xe8\xd9\xf1\x86\x78\xd4\x22\xd8\x90\x37\xc9\xcc\x10\x3e\x9d\xa0\x90\x6c\x14\x85\xb6\x7f\x97\x7a\x90\x6f\x10\x10\x57\xbc\x83\x29\x2a\xea\x4a\x66\x28\x1e\x37\xae\x30\xcb\x01\x38\x4f\xcc\x7e\x70\x92\x93\x34\x08\x35\xb3\x3b\x67\xe5\x96\xb4\x39\x0e\xcf\x26\xba\x24\xf0\x93\xae\xe2\x9a\xd5\x25\x98\xff\x16\x67\x21\x91\x49\xdc\x16\xce\xce\x0a\x21\xb0\x1e\x25\x0c\x1e\x16\x0a\x92\x8f\x63\x9f\x56\x92\x37\x1f\x5e\xf5\xcf\xfb\xd7\xc6\xd9\xdb\xe3\x5f\xfb\xc6\x87\xcb\xf3\x9e\xc8\x08\xbb\xed\x36\x2c\x99\xea\xd0\x0a\xd0\xbc\x47\xa5\x88\xbd\x88\x79\x7d\xfc\x6b\x6f\xaa\xab\xfa\xae\xaa\x97\xfb\x03\xe2\x5a\xa1\xab\xda\x60\xf1\xf7\x7c\x52\x94\x3b\x33\x04\xf3\x62\x3c\xb0\x9d\xdf\xaf\x66\xdd\x18\xc0\x2a\xe4\xf9\xe0\x98\x67\x2a\xf0\x30\xf8\xa9\x8e\xe4\xd1\xd3\xf7\xf5\xc3\xce\xdc\x41\x2e\x2b\xd7\xfc\x41\x3e\x58\x72\x90\xd7\x7f\x48\x22\x89\xcc\x53\x06\xe9\x90\x92\x1d\x51\xbc\x63\x4e\xf2\x9e\x3d\xca\x3b\x05\x5c\xfe\xc0\x5e\x60\x41\xc5\x3d\x72\x65\x10\xf8\x63\x85\xef\x1f\x97\xa0\x64\x90\x2a\x54\xac\x26\x2d\x49\x47\xa0\x84\x9e\x0b\x56\x25\x91\x6c\x5e\x32\x07\xb5\x3a\x68\x2d\x80\xcd\x55\xb2\xae\x54\xa4\x12\x69\x60\xc8\x7b\x51\x7a\x81\x3e\xbe\xf2\x05\x0e\x69\x4f\x58\x73\xe8\xc2\xef\x8d\xdd\xaa\x5c\x13\xf9\x42\xb8\xb1\x15\x1f\xba\xbe\x5c\x9b\x98\x8c\xdd\x89\xcd\x14\xd0\x31\x3e\xef\xf1\x6a\x8d\x01\x54\x72\xe0\x11\x87\x36\xd4\xe6\xc7\x19\x5d\x30\xf9\xe4\x68\xe3\xf9\xb6\xcd\x1f\xbd\xff\x57\x75\x20\x57\xc3\x63\xc1\xfe\xdf\x1e\xbc\xcc\xef\xff\xed\x6a\x9d\x9d\xbd\xd5\xfe\xdf\x73\xb4\xc5\xfb\x7f\xfc\xd3\xac\x31\x1d\xdf\x2c\xde\x04\x04\xab\x0d\xfc\x09\x98\x77\xba\xd8\xc5\x54\x3a\x9a\x22\xbf\x30\xe3\x21\x73\x76\x59\xc7\x56\x15\x9a\x60\x2b\x47\x13\xd8\x72\xdb\x34\xfd\xeb\x93\xd3\xac\x8b\xd9\x55\x61\x61\xd0\x5a\x0b\x10\x2e\x8e\xdf\xf6\x21\x82\x4f\x8e\x79\x1f\x16\x21\x1c\x9f\xfe\x4f\xff\xf2\xfa\xec\xaa\x6f\x9c\x9c\x9f\xf5\x2f\xae\xd1\x13\x5e\xa5\x47\x73\x5f\x32\x27\xc1\x78\xfa\x76\x70\xb4\x88\xe0\xd9\xc5\xd9\xf5\xd9\xf1\x79\x86\xf0\xfb\x7e\xff\x72\x01\xd9\xc3\x85\x1d\x3b\x3f\xbb\xba\xee\x5f\x54\x0a\xa9\xa9\xfc\x7f\x8d\xa4\x93\x64\xca\x22\xa5\x44\x1a\xcb\xf2\xb6\x7f\x7d\x79\x76\x72\x95\xd2\xc9\x93\xd1\x9b\x8e\xd4\xc9\xf9\x07\xa0\x77\x19\x4f\x5b\xe1\x40\x7b\xe1\x0c\x5e\xa1\x14\xd7\xc6\x65\xff\xe4\xdd\xc5\x2f\x67\xbf\x1a\x27\xaf\xfb\x27\x6f\x78\xe0\xb9\x10\xf3\xea\xdc\x38\x3d\xbb\xec\xe5\xbf\x18\x5c\x80\x74\x7d\x89\xd2\x9e\x1a\x27\xc7\xc6\x2f\x67\xe7\xfd\x14\x19\xf7\x07\x58\xd5\xf7\xa1\x0b\x08\x9e\x80\x9a\x2c\x22\xd5\x84\xce\x9b\xfe\xff\x2e\x24\x23\xbf\xc4\x9d\x2b\x8e\xd0\x30\x2e\xd5\xf1\x87\xeb\xd7\x8d\x46\x92\xeb\x53\x93\x91\xc9\x7e\x68\xda\x84\xe6\xfc\xc1\xc9\x7e\xd6\xbc\x90\xd4\xdc\xf1\xc9\x7e\xa7\xbc\x58\xa8\xfa\x21\x5a\x9d\xae\xf0\xb6\x3a\x5d\xc9\xf4\x36\x03\x52\xe7\x74\x57\x87\x30\xab\x43\x98\xd5\x21\xcc\xea\x10\xe6\x0f\x38\x84\x29\x22\x8f\xa8\x75\x3b\xf1\x1d\x2f\x8c\x6b\x02\x3c\x86\x8a\xe3\x99\x16\x4f\x3e\x97\xeb\xc8\xea\x40\x68\x75\x20\xb4\x3a\x10\xfa\xbe\x0f\x84\x20\x64\x0d\x4b\xca\xd2\x04\x31\xdd\x79\xea\xb5\x78\x40\xfd\x5c\xc7\x4a\x01\x1d\x3a\xfc\xa6\xf5\x1d\x44\xe0\x4a\x08\x33\x1c\x2e\xea\x5e\xaf\x7b\xe1\x5f\x81\x59\xe1\x4d\xe8\xef\xf3\x90\x4a\xd7\xe6\x04\xb5\x0d\x4f\xa9\x92\x2b\xe0\xf9\x80\x71\xc9\xd8\xf4\x55\x5c\x25\x87\x98\xd9\x93\x28\x69\x73\x09\xca\x89\xef\xd9\xfc\x9b\x89\xf7\x30\xf5\x90\x0e\xc0\x34\xf7\x7e\xc8\x55\x5a\x69\xe3\xce\x81\x91\x08\x85\x91\xfb\x93\x66\x1c\x1f\xc1\x36\x1c\x6f\x78\x0a\xb1\xbc\x15\xfa\xc1\xac\x97\xe3\x5e\xb5\xa0\xe6\xa4\xcb\x17\xf9\xa9\x58\xcd\x21\x88\x91\x07\x63\x7e\x64\x8d\x48\xf3\xbe\xad\x4e\xc3\x56\xa7\x61\x12\x79\xb9\xd3\xb0\xc5\xfa\x39\x7f\x94\xf7\x92\x51\x46\x4e\x19\xd1\xec\x2e\x96\xe2\x92\xbf\x0f\x03\x3f\x9a\xd4\xbd\xfc\x9a\x73\xb4\x75\xf2\x51\x86\x26\xb8\x5d\x91\x14\xe8\xca\xdb\x57\xc5\x71\xdb\x3a\x79\xeb\x4f\x29\xa1\xf7\x80\xea\x60\x3e\x6a\xba\xa4\x1c\xa8\x82\x61\x11\xc5\x23\xad\x8d\x2d\xb7\x50\xd2\x49\xd6\x79\x4a\xfd\x8d\xf2\x53\xfb\x27\xb2\xf3\x73\xb2\x85\xf0\xb2\x45\x3e\x93\x3f\xfd\x09\x0b\x43\x35\x40\x9c\x0f\x82\x74\x30\xac\x0b\x06\x0b\x69\x65\x44\xff\x06\x27\x8a\x82\x5f\x4d\x9a\x57\x90\x6a\xfe\xb1\xa2\xa4\x14\x67\x5f\x35\x28\x92\x6d\x52\x44\xed\x2b\xcf\x32\x53\x3a\x8d\x8e\x33\x37\x36\xbe\xe0\x1e\xc5\xbb\xf7\xd7\x57\x0f\x85\x57\xbf\x47\xe6\x4c\x04\x2c\x3c\x4c\x8e\xbb\xde\x9d\x6a\xaa\xde\x51\xb5\x92\x18\x8f\x3d\xbc\x4c\x14\x5a\x51\x64\xb1\x30\x88\x93\x71\xac\x44\x54\x21\x87\xad\xb5\xf1\x5f\xad\xef\xec\x70\x72\xd5\xbe\x79\x7b\xdc\xf9\x2f\x44\x96\x0e\x6e\x35\x36\xab\x00\xb5\xa8\xfe\x43\x67\xbf\x54\xff\x69\x57\xdf\x59\x9d\xff\x3e\x47\x5b\xcf\x7d\xf9\xba\x16\x4f\x2c\x69\x65\x3f\xf9\x13\xdf\x31\xc6\x85\x1d\xb1\x61\x71\xc7\x30\x80\x20\x17\x77\x1b\xec\x34\x48\xc7\x17\x10\xc1\xff\x1e\xd1\xa4\xb8\x25\xff\x92\x70\x6b\x12\x88\xaf\x70\x6d\x7e\x9d\x2c\xae\x00\xc0\xbf\x2a\x4e\x79\xc6\xcf\x97\xe3\x77\xfc\xf1\x8a\x5c\x8a\x3a\x06\x69\xd1\x88\x2d\xaa\x0e\x55\x62\xde\xb1\xec\x07\xb1\x75\xfc\x0c\xc7\x7e\x02\x96\x67\xa7\x92\xeb\x5f\x76\xdf\x1f\x1f\xbf\x7a\x75\xf2\xcb\xf1\x9b\xfe\x89\x56\x64\x9a\x7e\xa9\x1e\xd7\xef\xcc\x91\x7e\xcf\x4d\x2f\xfe\xd2\x3d\x25\xda\xd1\xf6\xe8\x80\x1e\x59\x8a\x65\xd1\x23\xa5\x63\x1d\xe8\xca\xe1\x81\xa5\x2b\x47\x9d\x23\x6b\xe7\x48\xbb\xd9\x39\xb0\x2d\xc1\x6a\x9d\x5c\x60\xcd\xae\xdc\x64\x16\x3e\x38\x5d\xd8\xdb\x81\x09\x09\x41\xfc\x42\xaf\x18\x81\x8b\x08\xcf\x26\x88\x3f\xc8\x7d\x4e\xbd\xe5\xa8\x54\x25\x22\xf5\x65\xc5\x8e\x67\xab\x32\x3d\xa5\x00\xb2\x18\x5d\x81\x5b\xa1\x4c\xc1\x72\x0c\x6f\xcc\x80\x8e\x29\x84\x7b\x86\x56\xc1\x5a\x4e\x91\x83\xd9\x94\x67\x51\x41\x15\x23\xcb\x6c\xe5\x86\xca\xbe\x3f\x9b\x24\xd5\x63\x52\xa8\x7d\xb0\x9c\x24\xf1\x05\x80\xc0\xbc\x53\x87\x4e\x38\x8a\x6e\x30\x84\x90\x61\x39\x5a\x57\x1b\xe2\x2e\x88\xa1\x6e\x13\xb7\x81\xec\x14\xc1\x8e\xb5\x3b\x03\xba\x7f\xb4\xb7\xdb\x19\xec\x1f\x69\xfa\xee\xcd\xd1\xfe\xa1\x6e\x1f\xda\xe6\x01\x85\x6c\x03\x1e\x6b\x07\xb4\x63\xe9\xd6\xde\x91\xf4\x3d\x6a\x5c\x10\xa2\xd0\xe3\x73\x5f\x6c\x12\xe2\x32\xe2\xfa\xa6\xcd\xcb\x7f\x60\xe1\x06\x0c\x82\x88\x00\x25\x18\xa3\x16\xba\x1e\xd7\xa1\x58\xce\xce\xe5\xf8\x26\xe5\x2d\xf8\x55\xd8\x89\xeb\xcf\x38\xdb\x78\x85\x73\x3c\x69\x77\x27\x3c\x41\x8d\x02\x2e\x61\x96\x7b\x1c\x3b\x55\x72\x77\x1d\x16\x56\xf0\xbe\xba\x7a\x1d\x17\x2d\xe1\xdf\xc6\xe3\xb4\xf2\x2a\xb8\x9b\x18\xa9\x6d\x16\xfa\x97\x54\xb1\xad\x59\x54\xc2\x11\x2f\x99\x42\xec\x78\x53\x42\x96\xd5\x19\x52\x0f\x6b\x72\xc0\xba\x2c\xc3\xc4\xb4\xc0\xce\xc4\x35\x2d\x78\xbe\x85\x53\x6c\x82\x8a\x11\xb9\xcf\xff\xb2\x55\x3f\x82\x39\x99\xe4\x36\x21\xbe\xa8\x12\xea\x64\xe4\x3b\xa0\xb3\x60\xc1\x29\xa0\xa8\xac\x6b\x43\x2f\xb7\x40\x8b\xc0\xdf\xba\x58\xa2\x41\x6c\xdb\xcf\xe3\x5b\xd2\x56\x81\x52\x2d\x90\x31\x0e\xa3\x6a\x89\x2e\xce\x60\x32\x61\x4a\x61\xc2\x29\x79\x7b\xfd\x81\x6c\x41\xda\xe8\x62\xdd\x08\xfc\x4b\x08\xe2\x7c\x0b\xf7\x07\x97\x12\x45\xef\xe0\x4d\x99\x4a\x41\x9c\x09\x46\xda\xb0\x5a\x87\xa2\x84\xb2\x01\xa6\x3e\xf2\xab\x1d\xc3\x5b\xfe\x8a\xcf\x62\x82\xc2\xd5\x90\x67\x06\xb9\x4a\x43\x4f\x21\xf5\xc0\x09\x18\xde\x1d\x88\x3c\xbb\xe8\xc0\x7c\x9b\x97\x50\xa8\x1e\xc2\xb3\xd3\x4b\x21\x4c\x60\x7a\x43\x5e\x04\x1b\x14\xcb\x19\x7a\xd9\x2d\x3b\xa0\xc0\x96\x1a\x41\x4d\xdd\xc1\xdb\x42\x6d\x7d\xbf\x20\x4b\xb6\xa4\x43\x95\x3c\x7f\xfe\x73\xff\xdd\xe9\x5a\x23\xa1\x24\x29\xa6\xae\x5d\xc3\xa0\xea\x7c\x4c\xc9\x9d\xe3\xba\x68\x0b\x30\xac\x78\x31\x46\x94\x54\xc2\x40\xd8\x48\x0a\x25\x6d\xf3\x49\xd0\x35\xb0\xb0\x3a\x04\x34\x58\xbc\x81\xb1\x86\xc2\xac\xe5\x4b\x64\x97\xfa\x2c\xfa\xbb\x5b\xdd\xdf\xca\x03\x86\xca\x89\xf8\xef\x88\xe2\x85\x0c\xce\x5f\x40\x33\x82\xfb\xd6\x5c\x5a\x59\xe2\x22\x16\xd6\xf4\xd8\x1d\x16\xf8\x20\x37\xb3\x44\x56\x72\x2a\xe5\x71\x92\x2d\x56\x3c\x0a\x30\x5d\x19\x8f\x0c\x7c\x5f\x95\x22\xab\x6c\x6a\xa9\x39\x98\x97\x64\x29\x5b\xcd\xa2\x16\x3a\x2c\xae\x34\x18\x10\x3d\xfa\x41\x98\xac\x24\x4d\xd7\xef\x3e\xc7\x16\x15\xca\x70\x21\x31\x3d\xd3\x9d\x85\x8e\xc5\x48\x42\x10\x95\x21\x9a\x00\x05\x6a\x8e\xc1\x72\x4e\xd2\xa5\xa6\x68\x10\xa2\x20\xd0\xff\xb3\xf4\xf7\x71\xf9\x5f\x20\x6e\xaa\x3c\x51\xfd\xef\x7d\xad\x5c\xff\x5b\x5b\x7d\xff\xff\x2c\x6d\x9d\x5c\xd3\x20\x30\xb1\x74\x1b\xd6\x42\xc3\x9a\xf6\xfc\x6f\x19\x88\x43\xa7\xf8\x11\x24\x13\x61\x02\x86\x06\x28\x15\xc0\x36\x62\x1c\x30\x90\x9f\x7b\x44\x53\x75\x5d\x15\xce\x2e\x71\xe4\x2d\x4b\x46\xfb\x19\x50\xbe\xbe\x15\xc0\x84\xf5\x17\x21\xff\xf1\x33\x29\x53\xc4\xad\xd3\x66\x90\x71\x51\x9d\x86\xd0\x2e\x6b\x06\x28\xec\xa3\x1a\x76\xa7\x00\x0b\x39\x68\x35\xe0\xde\xc1\xbf\xc0\x6a\xf2\x38\xfb\x4f\xce\x97\x1a\x2d\x00\x0b\xec\x5f\xdf\xd7\x3a\x45\xfb\xd7\x3b\xab\xfb\xff\xcf\xd2\x44\x91\xae\xf4\x2f\x8b\x08\x55\x95\xa5\x09\x41\x53\x21\xd3\xeb\x76\xe3\xd4\x4f\xa4\x7d\xb9\x5c\x2f\x59\x16\x14\x51\x25\x4c\x89\x09\xfd\x67\x40\x07\xbd\xdd\x81\xbe\x73\x68\x59\xfb\xba\x3e\xe8\x68\xb6\xa9\xdb\x47\x16\xfc\x7f\x17\x7e\xd1\x0f\x0e\x06\xe6\x91\x45\xf5\x9b\x4e\xe7\xe6\x50\xfc\x01\xb0\x26\xa5\xa7\xd6\x9e\xac\xc6\x1b\x04\x70\x46\x5c\x33\xb2\xaa\x89\xda\x8e\xb9\x9a\x96\x4d\x4a\x5a\xe2\xe1\x6a\xe6\x2f\x3e\x55\xd2\xe6\x9d\xab\xf8\x63\x2c\x71\xc1\xb5\xe4\x83\x84\x24\xab\xab\x94\xb0\xf6\x8f\x98\x64\xf2\xa9\xf9\x68\x29\x60\x16\x0f\xf3\xa3\x46\x78\x08\x98\x43\xac\xc9\x67\xca\x88\x35\x80\xa2\xf4\xa4\xcc\x2f\x6a\xfa\x9c\xa9\xf9\x28\x01\x39\x56\x36\x13\x98\x8b\x95\x05\x14\xe5\x12\x2b\xcb\xbf\x55\x60\xd6\x17\x82\x2b\x46\xa9\xb5\xcc\x8b\x80\xab\x3f\x75\xb1\x6a\xab\xb6\x6a\x7f\x58\xfb\x27\x18\x12\x53\xc2\x00\x78\x00\x00\x01\x00\x00\xff\xff\x57\x0e\x9e\xaf\xbc\x13\x00\x00")

func lokomotiveKubernetesPacketTarGzBytes() ([]byte, error) {
	return bindataRead(
		_lokomotiveKubernetesPacketTarGz,
		"lokomotive-kubernetes-packet.tar.gz",
	)
}

func lokomotiveKubernetesPacketTarGz() (*asset, error) {
	bytes, err := lokomotiveKubernetesPacketTarGzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lokomotive-kubernetes-packet.tar.gz", size: 5052, mode: os.FileMode(420), modTime: time.Unix(1549296655, 0)}
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
	"lokomotive-kubernetes-packet.tar.gz": lokomotiveKubernetesPacketTarGz,
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
	"lokomotive-kubernetes-packet.tar.gz": {lokomotiveKubernetesPacketTarGz, map[string]*bintree{}},
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
