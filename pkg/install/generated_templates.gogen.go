// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 7129,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\xdd\x6f\x1b\x37\x12\x7f\xf7\x5f\x31\x50\x0e\x48\x0c\x48\x2b\xbb\x6e\x8b\xc3\xf6\x54\x5c\x9a\x0f\x37\x97\x26\x35\xe2\xe4\x0e\x7d\xaa\x29\xee\x48\x4b\x88\x4b\xee\x71\xb8\x52\x05\xa1\xff\xfb\x61\xc8\xfd\xe0\xca\xb2\x5d\xe4\xed\xf2\x10\x4b\xbb\x33\xc3\xf9\x9e\xdf\x50\xb3\xd9\xec\x4c\xd4\xea\xdf\xe8\x48\x59\x93\x83\xa8\x6b\x9a\x6f\x2f\xcf\x36\xca\x14\x39\xbc\xc6\x5a\xdb\x7d\x85\xc6\x9f\x55\xe8\x45\x21\xbc\xc8\xcf\x00\x8c\xa8\x30\x87\x95\x6e\xfe\x38\x1c\x40\xad\x20\xfb\x28\x2a\xa4\x5a\x48\x84\x3f\xff\x6c\xdf\x87\xaf\x39\x1c\x0e\xe3\xb7\x87\x03\xa0\x29\x98\x8c\x6a\x94\x2c\xcc\x61\xad\x95\x14\x94\xc3\xe5\x19\x00\xa1\x46\xe9\xad\xe3\x37\x00\x95\xf0\xb2\xfc\x45\x2c\x51\x53\x7c\x90\x9e\xcd\xd4\xde\x09\x8f\xeb\x7d\x7c\xe9\xf7\x35\xe6\xf0\x09\xa5\x43\xe1\xf1\x0c\xc0\x63\x55\x6b\xe1\xb1\x15\x96\x58\xc0\xff\x84\x31\xd6\x0b\xaf\xac\xe9\x85\x03\xd4\xce\x56\xe8\x4b\x6c\x28\x53\x76\x5e\x5b\xe7\x73\x98\x5c\x5d\x5c\x5d\x4e\xe0\x19\x78\xd4\x3a\xa1\x00\x6f\x81\xa4\x13\x35\xc2\xbc\x42\xef\x94\x24\x36\xae\xb6\xca\xf8\xe7\x04\xcc\x9c\xb5\x82\xf5\xc8\x86\x23\x2b\x00\x3a\x5f\x84\xcf\xe8\xb6\x4a\xe2\x4b\x29\x6d\x63\xfc\xc7\x31\x21\xc0\xd6\xea\xa6\xc2\x5e\xd4\xac\x15\xb5\x56\x7e\xb6\xc1\x7d\x7f\x00\xb1\x17\xfc\x70\x60\xf7\x64\x90\x37\x63\x96\x22\x04\x38\xa1\x2a\x70\x25\x1a\xed\x3f\xd8\x02\x73\xb8\xf8\xf6\xe2\x02\x9e\xc1\xae\x44\x03\x15\x6b\x83\x05\x38\x14\xc5\xcc\x1a\xbd\x9f\xc2\x0e\x61\x67\xcd\x73\x0f\x4b\x04\xb1\xd4\xc8\xfe\x90\x65\x65\x8b\xb3\x56\xe0\x33\xf8\x5c\x2a\x02\x45\x20\xc0\x57\xf5\x8a\xa0\x21\x2c\x60\x65\x1d\xac\xd1\xa0\x13\x5e\x99\x35\xdc\xde\xfe\x0c\x1b\xdc\x53\x06\xef\x0c\xbc\xff\x3b\xc1\x8f\x0b\xb8\xcc\x2e\x2f\xa6\xbd\x94\xee\xec\x68\x02\x81\x70\x98\xea\x41\x96\x55\x31\x88\x05\x08\x20\xac\x05\x27\x45\xeb\x28\xd8\x61\x2f\x46\x0a\x03\x3b\xa7\x3c\x2b\x9a\x9d\xf6\xdf\x1a\x4d\xef\x0c\xac\x6a\xbf\x7f\xad\x5c\xea\xc4\x0a\x0b\xd5\x54\x39\x7c\xc0\xca\xba\x7d\x6a\x27\xc2\xca\x6a\x6d\x77\x6c\x51\x7b\xb4\xa2\x60\x6a\x43\xfc\x4c\x80\x6c\xc8\xdb\x4a\xb1\x07\x36\xc6\xee\xcc\xef\xa5\x25\x4f\xbd\x88\x95\xd2\x38\x85\x5d\xa9\x64\x09\x7b\xdb\xc0\x4e\x69\x1d\x8d\xf2\x16\x0a\xcb\x75\xc6\x8f\x99\x89\x3f\x38\xb0\x3b\xc3\x6a\xf7\x02\x1c\xd6\x16\x9c\xf0\x25\x3a\xf0\xa5\x30\xed\xc1\x6b\xe5\xcb\x66\x09\x96\x1f\x22\x68\xb5\xc1\x0c\x7e\xb3\xcd\x73\xad\x41\x68\xb2\xdd\x11\x63\x67\x83\xf2\xa0\x8c\xb7\x81\x47\x5a\xe3\x85\x32\xe8\xa6\xb0\x44\x6d\x77\x19\xdc\xe2\xe0\xd5\xd2\xfb\x9a\xf2\xf9\xbc\xb0\x92\x32\x4e\x2c\x59\x70\xe9\xa0\x99\x73\xe9\x91\x9f\xaf\x1b\x55\x20\xcd\x1b\xc2\x59\xed\xd4\x56\x78\x0c\xa9\xc7\x86\x64\xa5\xaf\x74\x2f\xa9\x8b\x05\x51\x39\x93\xd6\xac\xd4\xba\x7f\x05\x10\x1f\x7c\x10\x75\x9e\x3c\x4c\x0b\x69\x96\xb0\x7d\x6d\x5c\xb2\x4d\xb3\xc4\x79\x14\x32\xa4\xdf\x93\x31\xd9\x29\x2a\xf9\x49\x29\xb6\x08\x02\x0a\xb5\x5a\xa1\xe3\xa6\xd9\x49\x68\xab\x6a\x68\x8c\x21\x04\x51\x5c\x1a\x04\x6e\x2e\x5b\x55\x60\xe7\xf6\x95\x5a\x57\xa2\x1e\x14\x51\xbe\x04\x61\x00\x8d\x77\xfb\x60\xc3\x5d\x24\xba\x9b\x82\x30\x05\x34\x46\xda\x8a\xbb\x75\xe0\x8f\xd6\x7e\x08\xe1\x14\xa6\xe8\xa5\xa0\xd9\x06\x09\x0a\xa9\x8d\xe7\xbd\x08\xb0\x1b\xbe\x22\x02\x09\xdb\x93\x11\x08\x9d\xc0\x5b\x50\x15\xf7\x49\xb8\xbe\xb9\x0e\x4d\x00\x5e\xb0\x59\xa4\xd6\x46\x99\xe1\x70\x36\x6e\x8b\x4e\xad\x94\x0c\x0d\x1b\xea\xc6\xd5\x96\x90\xce\xff\x82\x23\x7b\x29\xb1\x7d\x44\x2f\xb2\x83\xf8\xbc\xbf\xe0\x38\x10\x6e\x3d\x94\xe9\x03\x1e\x5b\xd7\x6b\xee\x1f\x94\xb8\x66\xdc\x82\x9f\x3d\xd0\x84\xef\xf3\x9d\x68\xc2\x9d\x3b\xfb\x4a\xbc\xd7\xff\x93\x09\xd1\x7a\xdd\x61\xe8\x93\xc6\xc2\x24\x8f\x95\x38\x01\x55\x89\x35\xc6\xec\x67\x86\x0c\xde\x2a\x53\x04\x9b\x2b\x6e\x2b\x0e\xe5\x90\xb5\xb1\xa5\x68\x14\x84\xdc\x3c\x02\x2b\x07\x81\x71\x02\x08\xdf\xd7\x7d\xd9\x2c\xb3\xc2\xca\x0d\xba\x4c\xda\x6a\xee\xe6\xb1\x07\x84\x3f\x73\x2f\x7a\xd7\x75\x71\xe4\x79\xcf\x58\x80\x4f\xf5\x62\x0d\xac\x69\xd6\xd3\x84\x63\x72\x68\x05\x2a\x9b\x4a\xcb\x2f\xb3\xcb\xef\xb3\x8b\x31\xed\x4d\xa3\xf5\x8d\xd5\x4a\xee\x73\x78\xb7\xfa\x68\xfd\x8d\x43\x4a\xad\x70\x48\xb6\x71\x12\x29\xed\xe3\x0e\xff\xdb\x20\xf9\xd1\x33\x00\x59\x37\x39\x7c\x77\x51\x8d\x1e\x56\xa1\xd5\xe7\xf0\xfd\xb7\x1f\xd4\x00\x13\xac\x4b\x99\x67\x43\x64\x6e\x02\x64\xb8\xba\xb8\xe2\xc9\xa9\xcc\xca\xba\x2a\xa4\xac\xd0\x3d\xb5\x56\x5b\x34\x48\x74\xe3\xec\x12\x53\x0d\xd8\xa5\xd7\xe3\xa9\x1d\x8f\x8a\x02\xc7\x8f\x85\x2f\x73\x98\x8b\x5a\x45\x4f\x6f\xbf\x9f\xab\x02\x8d\x57\x7e\x9f\xd5\xcd\x32\xa1\x55\x46\x79\x25\xf4\x6b\xd4\x62\x7f\xcb\xf5\x59\x50\x0e\xdf\x25\x04\x5e\x55\x68\x1b\x7f\xe2\x1d\x0f\x59\xf5\xff\xa1\x6a\x52\xb4\xa3\xc0\x9c\x86\x47\x10\xc7\xdc\x4d\xd4\x0c\xbd\x0c\x9a\x15\x73\xa2\x92\x71\x9e\x8d\xc8\x13\xb4\x6d\xfb\xcd\x9a\x43\x06\xca\xc4\x9c\x7b\x4e\x91\x87\xa8\x9c\x8f\xda\x64\xe7\xb3\x5f\x8d\xde\xe7\xe0\x5d\x83\x2c\x8d\x31\x50\xe8\x50\xcb\xb6\xb1\x73\x49\xd5\xe8\x56\xd6\x49\x64\xa1\x11\xf4\x30\xe6\x79\x48\xf1\x14\x97\x8c\x75\xdf\x0a\xd7\xea\x1e\xc9\xbe\x4e\xfd\xa4\x46\xdf\x19\xa9\x9b\xd0\x39\x19\xba\xc5\x01\xd7\x75\xd5\x88\x0d\x9e\x80\x32\x1d\x98\xf9\x81\x59\x8f\x60\x46\xdf\x5d\xa1\x40\xa9\x85\x63\xc8\xb6\xb4\xdb\xa4\x01\x3c\x02\x03\x62\x7b\x4c\x8d\x77\xd6\xfa\x79\x46\x54\x3e\x68\x80\x30\xa3\x53\x27\xc3\x88\x9a\xc4\x93\xa7\x1d\x49\x22\x01\xcd\x56\x39\x6b\xc2\x40\x88\xb3\x76\xf2\xfe\xcb\x4f\x6f\x5e\xfd\xfa\xf1\xed\xbb\xeb\x49\x1c\x01\x53\xf6\x87\xdd\xa2\x73\xe3\x79\x9d\x88\x09\x23\x6e\xb9\x8f\xd3\xd4\xeb\x53\x36\xde\x1b\xb4\xf7\x6d\x1c\x92\x93\x89\x1f\x34\x94\x67\x1e\x2f\x1e\xdd\x69\xdc\xa2\x13\x28\xd2\x6a\x17\x62\x92\x88\x38\x06\x34\x69\xd0\x03\x9a\xe9\xa0\xb7\x30\x20\xb4\x47\x67\x18\x5a\xdf\xd3\x78\xe5\x6c\xc5\x69\xd1\x21\x96\x29\x08\xe2\x74\x6b\xa7\x2a\xbb\x41\x5b\xb9\xa1\xfb\xc1\x46\xb3\xcd\x4f\xf8\x65\x70\xf7\xc8\x2f\x5b\xa1\x1b\xbc\xe7\x93\xa7\x92\xf8\x38\x07\xba\x99\xfb\x48\x06\xf0\xc8\x1f\x8f\xfa\x47\x86\xfd\x03\x79\xc9\x54\x11\xdd\x8c\xe8\xc6\xfd\xe1\xa9\xca\xdb\x09\x06\x25\x16\xa8\xa9\x6b\xbd\x87\x9f\x3f\x7f\xbe\x81\xa5\x20\x25\x41\x34\xbe\x04\xe9\x30\x74\x52\xa1\xe3\x54\x1f\xf6\x01\x16\xb8\x55\x22\x18\x7e\x77\xfd\xee\xf3\xef\x2f\xbf\x7c\xfe\xf9\xcb\xed\x9b\x4f\x77\xc1\xdc\xfe\xd1\xfb\x37\xbf\xdd\x8d\x12\x7e\x2b\x9c\xe2\x6d\x8e\x3a\x80\x9c\x08\x8c\xf0\xe5\x28\x7e\x6f\x9d\xad\xc6\x31\x8c\x64\x9f\x70\x95\x8f\x2c\x1f\x61\x45\x6e\x6c\x6c\xc2\xe0\x00\xf6\x79\x3e\xf2\x47\x74\x41\xdc\x51\xb1\xe0\x49\x2c\x85\x2c\xb1\xe0\xd4\x4a\x73\xbb\x87\xd5\xec\x29\x96\x3e\x4d\xa4\x58\xd7\xe2\xe6\x84\xa1\xdd\xb1\x03\xe3\x34\x1c\xc2\xbb\x61\xeb\x63\x5f\x22\xa5\xb9\x30\xa0\x57\xbf\xb3\xac\x65\xc3\x7e\x0a\x15\x17\x2e\x04\x42\x22\x42\x69\x77\x61\xff\xb5\xc6\xa0\x0c\x21\x53\x7e\x9c\x3b\xb3\x59\x6f\x40\x58\x7e\xf8\xf0\x45\xff\x28\x6b\x41\x5f\x46\x5b\x99\x49\xdd\x90\x47\x97\x71\x03\xd7\xa9\x4b\xbe\x50\xec\x35\x83\x2b\x5e\x45\xd2\x77\x37\x23\xa3\xb8\xed\x10\xfa\xb0\x5f\x8f\x33\x7b\xd0\xa1\xa3\xe7\xec\xf2\x8e\x29\xc3\xc6\x9b\x8c\xa0\x54\xe3\x96\x7a\x71\x36\x42\x99\x8a\xa0\x6a\x28\xdc\x00\x04\xef\x29\x2c\x62\x39\x2d\xc3\x60\x0b\x18\x2f\x2c\xfe\x2f\xba\x6d\xfa\x3c\xd5\xa5\x6b\x2e\xb1\x0c\x39\x81\x93\xfd\x7f\xa4\x08\x0f\x83\x38\xe0\x66\x85\x72\x8b\x7b\x63\x2f\x55\xeb\x53\x82\x30\x87\xe0\x7d\xf9\xf4\x4b\xbc\xa0\x10\x66\x1d\xdf\x5d\x2b\x1f\x96\x66\x52\xde\xba\x7d\xdf\xae\xdf\x32\x32\x4e\xc4\x3d\x56\x73\x9c\x36\x89\xed\x6d\xc9\x9c\x2c\xa7\xb4\x16\x3a\xec\xfc\xb7\x17\x69\x65\x9e\xe7\xc3\xf7\xf7\x6f\x7e\x3b\xff\x67\x5c\xdd\x03\xac\x6e\x08\xdd\x7c\x50\x36\x4b\x0b\x9d\xfd\xc3\xe5\xd4\x38\xbd\x38\x1c\x20\xbb\x56\x9e\x8d\x0d\x57\x71\x63\x8a\xa5\x13\x46\x96\x1d\xd1\x4f\xe1\x5b\xbc\x94\x53\xab\xf0\x88\xfb\x17\x9d\xe2\x64\x0c\xc7\x7c\xb7\x21\x53\xe8\x5f\x56\x99\x84\x61\x32\x9d\xb4\x77\x7b\x9a\x30\x65\x7f\xbc\xa9\x39\xe4\xc4\x93\x71\xeb\xaa\x84\x51\x2b\xc6\xe4\x5c\x43\xa4\x0a\x74\x31\x1c\x47\x9b\x4d\xb8\x93\xb0\x84\xd0\x98\x02\xdd\x51\x8c\x1d\x6a\xe1\xd5\x16\x03\xe4\xa4\x2e\x03\xd7\xa3\x38\x1f\xd5\x64\x6f\x1c\x35\xcb\x42\xb9\xcb\x69\xfc\xfb\x4d\x7f\x51\x39\x38\x27\x5c\x44\x9e\x72\x4e\xb8\xdd\xeb\xbc\xda\x51\x9d\x10\xf0\x85\xd0\x9d\xe2\xe7\xe0\xf6\x91\x63\x1a\x38\xcd\xff\xa6\x12\xea\xa4\x02\xc8\x2f\x3a\x09\x1d\xd5\x70\xd5\x7a\x32\x1c\xc8\xad\x64\x67\xd9\xa1\x68\xc2\xf5\x1d\xfb\x89\x27\xb6\xf2\x47\x0b\x78\xea\xab\x76\xf6\xb5\x93\x6d\xf1\xc8\xa8\xeb\x38\x5a\x59\xcc\xb5\xf8\xc7\x06\xf7\xa0\x8a\x1f\x7b\xb2\x47\xe0\x4c\xa2\x15\x8b\x10\xbe\x71\x38\xba\x05\x38\x71\x56\x78\xbd\x9f\xf5\xf4\x34\x6a\x57\x5d\xb7\x06\xe5\xa1\x14\x14\x46\xb1\x35\x7a\x0f\x42\x4a\xa4\xd8\xd1\x4b\x8c\x17\x69\x2f\xba\x3b\x9b\xbb\x95\xd0\x84\x77\xe7\x67\x87\xc3\xac\x0b\xc4\xa7\x76\x86\x9f\x8a\x45\x27\x34\xd0\xdf\xaf\x87\xd3\x64\x27\xe2\x44\xde\x35\xd2\x47\x7d\x77\x61\x9d\x67\x88\xd7\x78\xa0\xbd\x91\xb0\xb4\x76\xb3\x41\xac\x39\xeb\x7b\x55\x27\x6b\xe5\x27\x53\xa8\x50\xb0\xc3\xb9\xa1\x81\x08\x3b\x76\x5b\x08\x4d\x4d\xde\xa1\xa8\xfa\x8a\x38\x3f\x52\x8c\x45\xcf\xc8\x0b\x8f\x0b\x6e\x30\x0f\xe6\x8d\xc1\x3f\x7c\x97\x3c\xc9\xc4\x13\x06\x26\xdd\x19\x93\x6e\x1e\x25\x42\x5e\x60\xb6\xce\xa6\xf0\x1f\x64\x64\xf9\x4a\xdb\xa6\x38\xcf\xc2\x05\x91\xb7\x1b\xde\x4f\x08\x6a\xe1\xbc\x92\x8d\x16\xae\x0b\x46\x2b\xe5\x78\x94\xb6\xa7\x2e\x76\xc4\x7d\x54\xb2\xac\x6c\xc7\x72\xb3\x9d\x75\x1b\xea\x97\xcd\x23\xb6\x70\xd0\x42\x2c\xe5\xe5\x37\x57\xf7\xff\x4f\x0d\x7e\x13\xb3\xaf\xeb\x4a\xfd\x85\xb5\x35\x8f\xa4\xc6\x87\x96\xfa\x7a\x20\x3e\xca\x90\x4e\xde\x6c\x90\xb7\x08\x38\xf0\xe1\x6c\x39\xc5\x12\x0e\x7e\x20\x75\x6e\xd1\x6d\x4f\xfc\x22\xc1\x0b\xc1\x80\x80\xb8\x56\x7f\x48\x47\xb1\xd8\xf0\x18\x8b\x59\x46\xe8\x93\x9f\x39\x9e\x27\xbf\x94\x24\x3f\x79\x70\x70\xc2\xd5\x5d\x00\xe5\xd9\xc8\x4a\xad\xc8\xa3\x99\xb5\x2a\x2c\xf2\xab\x8b\xab\xcb\xb3\xb6\x8f\xbd\x2c\x0a\x15\x2f\x44\x78\xd0\xbe\x64\xa0\x3d\x32\x79\x78\x3f\x60\xad\xc3\x01\x5c\x18\xdb\x4f\x70\xcf\xc2\xef\x4d\xa3\xde\x37\x7c\xea\x0e\xf8\xb5\x6e\xc5\xbf\xfe\x78\xdb\x81\x24\x9a\xb6\xcb\x4b\xe3\x5a\xc8\x04\xa6\xb0\x9e\xc0\x06\x62\xa8\xc4\x3e\x5c\x24\xe9\xed\x70\x9d\x68\x48\x5b\xbb\x69\x6a\x50\x44\x0d\x12\x58\x03\x64\x2b\x84\xf7\xcd\x12\x9d\x41\x8f\xc4\xd2\x9b\x9a\x86\xdb\xc2\xc2\x50\x77\x57\x35\xf9\x68\x0d\x4e\xd2\x37\xaf\x82\x02\xe9\x7d\x61\x3c\x9c\xc6\x57\x88\xdd\x12\x12\xf4\x1b\xbd\xe9\xf7\xa3\xc9\xe5\xe4\xec\x7f\x01\x00\x00\xff\xff\xf1\x09\xea\xfa\xd9\x1b\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 874,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xcd\x8e\x9b\x40\x10\x84\xef\x3c\x45\x49\xbe\x06\x27\x58\xf2\x85\x5b\x94\x4d\xa2\x95\x92\x95\x2f\x9b\x7b\x7b\x68\xf0\x28\xf3\x97\xe9\xc6\x31\x41\xfb\xee\x11\xfe\xc5\xbb\x3b\x27\xa0\xea\xeb\xa9\x29\xa0\x2c\xcb\x62\x01\xcf\xde\x90\xd9\x71\x83\x86\x93\x8b\x83\xe7\xa0\xe8\x85\x1b\x6c\x07\x7c\x73\xfd\x01\x1a\x71\x74\x14\x0b\x98\x18\x94\x6c\xe0\x0c\xeb\xa9\x63\x78\x56\x6a\x48\x69\x59\x50\xb2\xbf\x38\x8b\x8d\xa1\x06\xa5\x24\x1f\xf7\x55\xf1\xdb\x86\xa6\xc6\xc3\x75\x6c\x71\xb1\xd7\x05\x10\xc8\x73\x7d\xdb\x7d\x1c\x61\x5b\x2c\x9f\xc8\xb3\x24\x32\x8c\x97\x97\xb3\xe9\x78\x5b\x63\x1c\xef\xd5\x71\x04\x87\x66\xb2\x49\x62\x33\x4d\xcc\x9c\x9c\x35\x24\x35\xaa\x02\x10\x76\x6c\x34\xe6\x49\x01\x3c\xa9\xd9\xfd\xa0\x2d\x3b\x39\x3d\x78\x13\xa0\x00\x94\x7d\x72\xa4\x7c\x46\x66\x61\xa7\xe5\xee\xe8\xf7\x78\xe0\x12\x65\x5a\xd7\xae\xae\x4c\xf9\x2e\x33\xad\x63\x9b\x33\xa1\xae\x96\xeb\xe5\xea\xd3\xbd\xbe\xe9\x9d\xdb\x44\x67\xcd\x50\xe3\xb1\x7d\x8a\xba\xc9\x2c\x53\xad\x17\x17\xe5\x6e\x96\xaf\x44\xe9\xb1\xae\x56\x00\x16\xf8\x49\x07\xeb\x7b\x3f\xed\x10\xf3\x30\xbd\xd2\x5e\xf8\x03\x6c\x80\xe7\x8e\xb6\x83\xb2\xcc\xc1\x47\xac\x3d\xee\x40\xb1\xff\x18\x6d\xcc\x88\x81\x61\x95\xfd\xdc\x9e\x50\x55\xab\xaa\xc2\x02\x0f\xdc\x52\xef\x14\x29\xe6\x5b\xae\xc5\xe4\xd9\xef\x4f\x97\xcf\xc1\x44\x7f\xfc\xc8\x34\xa2\x63\x85\x8b\x9d\x20\xb6\x60\x32\x3b\x64\xfe\xd3\xb3\x28\x28\x34\xc8\x2c\x29\x06\xe1\xe5\x75\xd0\x34\xf5\xee\x84\xa7\x3e\x8d\xb3\x1c\xf4\x76\x80\x59\xf7\x9b\x98\xb5\x3e\xa5\xbb\xca\xc2\xa6\xcf\x56\x87\x2f\x31\x28\x1f\xb4\x9e\x71\xb9\x0f\x9f\xe5\x59\x38\xbf\x66\xce\xd2\xf7\x1c\xfb\xf4\x56\x23\xe7\xe2\xdf\x4d\xb6\x7b\xeb\xb8\xe3\xaf\x62\xc8\x91\x1e\x7f\x85\x96\x9c\x70\xf1\x3f\x00\x00\xff\xff\x15\x98\xb4\xee\x6a\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
