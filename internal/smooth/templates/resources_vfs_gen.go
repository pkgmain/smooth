// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package templates

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

// VFS statically implements the virtual filesystem provided to vfsgen.
var VFS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 4, 15, 18, 31, 37, 170384599, time.UTC),
		},
		"/app": &vfsgen۰DirInfo{
			name:    "app",
			modTime: time.Date(2019, 4, 18, 19, 35, 59, 821683946, time.UTC),
		},
		"/app/.gitignore.tmpl": &vfsgen۰FileInfo{
			name:    ".gitignore.tmpl",
			modTime: time.Date(2019, 4, 18, 19, 35, 59, 820729955, time.UTC),
			content: []byte("\x2e\x69\x64\x65\x61\x0a\x62\x69\x6e\x0a"),
		},
		"/app/cmd": &vfsgen۰DirInfo{
			name:    "cmd",
			modTime: time.Date(2019, 4, 18, 19, 22, 57, 265363818, time.UTC),
		},
		"/app/cmd/migrate.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "migrate.go.tmpl",
			modTime:          time.Date(2019, 4, 18, 19, 22, 57, 264221390, time.UTC),
			uncompressedSize: 895,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x91\x3f\x4f\xf3\x30\x10\xc6\x67\xdf\xa7\xb8\x37\xc3\xab\xe4\x55\x94\xaa\x7a\xb7\x48\x1d\xa0\x65\x2c\x03\xa8\x13\x62\x70\x62\x37\xb5\xa8\xff\xe8\x62\xd3\x21\xca\x77\x47\x0e\x21\x2d\x50\x28\x2b\xe3\x9d\xef\x9e\xe7\xf7\x9c\x1d\xaf\x9f\x78\x23\xb1\xd6\x02\x40\x69\x67\xc9\x63\x0a\x2c\x69\x94\xdf\x85\xaa\xa8\xad\x9e\xb5\x6e\x3b\xff\x3f\xab\x6d\x45\x3c\x01\x60\x49\xd7\x15\x6b\x2b\xc2\x5e\xf6\xfd\xac\xf5\x96\x64\x02\x19\xc0\x33\x27\xd4\xaa\x21\xee\xe5\x52\x0b\x5c\xe0\xdf\x61\xa3\x58\x5a\xad\xb9\x11\x1d\xb0\x4d\x2b\x4b\x44\x4c\xc6\xa9\x24\x07\x76\xbf\xb3\xe4\xcb\xa9\x85\x7e\x27\x51\x70\xcf\x2b\xde\x4a\x0c\x0e\x2d\xa1\xb0\x07\x93\xe4\xd0\xbf\x3a\x04\x77\x41\x3c\xb8\x73\xba\x27\x9a\xf1\xf9\x2e\x98\x9b\x12\xdf\x70\x37\x6e\x92\x8f\x66\x17\x0c\x46\x9e\xef\xd1\xe3\x10\xce\xb1\xf5\xf2\x8c\xdf\xca\x1e\xcc\xe0\xb8\x0d\xa6\x46\x65\x94\x4f\x33\xec\x80\x91\xb5\x7e\xa9\x45\x71\x25\xc4\x68\x9b\x1e\x0f\x9a\x01\xb0\x63\x75\x3a\x33\x9c\x24\xfb\xea\x75\x4c\x94\x4d\x7e\x53\xe8\xb4\xd6\x02\xff\xbd\x8b\x99\x23\xa7\xa6\xc5\x87\xc7\xd6\x93\x32\x4d\x86\x92\xc8\x52\x64\x13\x55\x1e\x0b\x2c\x17\x38\x7c\x79\xb1\xba\x4e\x33\x60\x6a\x3b\x74\xff\x2c\xd0\xa8\x7d\x9c\x63\x24\x7d\x20\x13\xbb\xc0\xfa\xc8\xfc\x61\xef\x56\x1e\xd6\x03\x81\xa5\x54\x54\x3f\x92\x18\x6b\x5d\x6c\x5c\xfa\x29\x47\x3c\xe6\xef\x4b\x32\x50\xc7\x2c\x2f\x01\x00\x00\xff\xff\xe5\x42\x20\x2b\x7f\x03\x00\x00"),
		},
		"/app/cmd/root.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "root.go.tmpl",
			modTime:          time.Date(2019, 4, 3, 19, 54, 59, 705615384, time.UTC),
			uncompressedSize: 265,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcd\xc1\x4a\xc4\x30\x10\x80\xe1\x73\xe6\x29\xc6\x80\xd2\x62\x49\x11\x6f\x85\x1e\x96\x45\x6f\x9e\x16\x1f\x20\x4d\xd3\x34\x98\x74\xc2\x34\x5d\xc5\x92\x77\x97\x0a\x0a\x7b\xfe\x7e\xf8\x93\x36\x1f\xda\x59\x34\x71\x04\xf0\x31\x11\x67\xac\x40\x48\xe7\xf3\xbc\x0d\xca\x50\x6c\x79\x6d\xbf\x2d\x53\x20\xd7\x06\x72\xf2\x16\xd7\x34\x3d\x3d\xb7\x86\x06\xd6\x12\x6a\x80\xab\x66\x64\xa2\x7c\x8e\x23\xf6\xf8\xf0\x0b\xea\x4c\x31\xea\x65\xdc\x41\xbc\xaf\xb6\x43\x44\xb9\xef\xea\x94\x52\x29\xb2\x01\x71\x99\x89\x73\x87\xf2\xb4\xa0\x4e\x09\x87\xcd\x87\x11\x3f\x7d\x9e\xf1\x12\x89\xf2\x2c\x1b\x28\x00\xd3\xb6\x18\x7c\xf9\xb2\x66\xcb\xb6\xaa\x71\x07\x61\x99\xb1\xeb\xff\x6e\xea\xdf\x40\xf8\x09\x0f\xbc\xeb\x71\xf1\xe1\x48\x45\x20\xa7\x5e\x75\xd6\xa1\xaa\xd5\xdb\xea\xa6\x4a\xde\x3f\x5e\x65\x73\x64\x35\x88\x02\x05\x7e\x02\x00\x00\xff\xff\x3e\xd1\xe4\xf6\x09\x01\x00\x00"),
		},
		"/app/cmd/serve.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "serve.go.tmpl",
			modTime:          time.Date(2019, 4, 3, 17, 46, 45, 519552069, time.UTC),
			uncompressedSize: 594,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x31\x6f\xab\x30\x14\x85\x67\xdf\x5f\x71\x9f\x87\x27\x78\x42\x8e\x9e\xba\x21\x31\xb4\x49\x87\x0e\xe9\xd0\x28\x53\xd5\xc1\x60\x43\xac\xc6\x36\xba\x36\xe9\x80\xf8\xef\x15\x98\x0e\x89\x3a\x64\xe4\x70\xce\x77\x7d\xcf\xed\x65\xf3\x29\x3b\x8d\x8d\x55\x00\xc6\xf6\x9e\x22\x66\xc0\x78\x67\xe2\x69\xa8\x45\xe3\xed\x26\xf4\xed\xff\x87\x4d\xe3\x6b\x92\x1c\x80\xf1\x71\x14\x7b\xaf\x86\xb3\x9e\xa6\x4d\xe3\x5d\x6b\x3a\x7e\xa3\x92\x1f\xa2\xa6\x5b\x35\x44\x4f\x9a\x43\x0e\x70\x91\x84\x41\xd3\x45\x6f\xad\xc2\x0a\xff\x2e\x6c\xb1\xf5\xd6\x4a\xa7\x46\x60\xc7\xa0\x4b\x44\xe4\x8b\x87\x17\xc0\x0e\x27\x4f\xb1\x44\x1e\xa2\xa4\x88\xf1\xa4\x53\x9c\xe6\x7f\x6f\x83\x7b\x2e\xb1\x1d\x5c\x93\x35\x56\xe1\xbf\x2b\x58\x81\x92\xba\x80\xef\x1f\x21\x92\x71\x5d\x8e\x9a\xc8\x13\x8e\xc0\x98\xaa\x8b\xf9\x0b\xcb\x0a\x97\x97\x89\xdd\x53\x96\x03\x63\xa6\x5d\xe4\x3f\x15\x3a\x73\x5e\x9c\x8c\x74\x1c\xc8\xcd\x32\x30\x36\x41\x32\xa5\xd5\xc5\x4b\xd8\xe9\x4b\x96\x27\xa3\xbd\x41\xbe\xea\xaf\xbd\xe9\x48\x46\x4f\x99\xaa\x67\xfa\x6f\xf8\x6b\x7e\x1a\xc0\x66\x53\x85\x56\x1c\xfb\xec\xee\xdc\x9a\x4d\xd1\x74\x04\x71\x98\x2b\x5b\x87\xdf\xb5\xda\xaa\x38\x73\x06\x36\x15\x30\x01\xcc\xdd\xa2\x71\x26\xa6\x35\xc9\xfb\xb8\xb5\x4a\x3c\x2a\xb5\xb6\x9c\xfd\x1c\x33\x87\x09\xbe\x03\x00\x00\xff\xff\x81\x43\x85\xf5\x52\x02\x00\x00"),
		},
		"/app/config": &vfsgen۰DirInfo{
			name:    "config",
			modTime: time.Date(2019, 4, 3, 20, 43, 8, 169557762, time.UTC),
		},
		"/app/config/env.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "env.go.tmpl",
			modTime:          time.Date(2019, 4, 3, 20, 43, 8, 167952338, time.UTC),
			uncompressedSize: 256,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xce\xb1\x6a\xc3\x40\x0c\xc6\xf1\xd9\x7a\x0a\xa1\xc9\x86\x12\x1b\x3a\x67\x8b\x09\x5d\xea\x50\x4a\xd7\x70\xb1\xe5\xb3\x1b\x9f\x64\x2e\xba\x5b\x4a\xdf\xbd\xa4\x5e\xda\x21\xab\xd0\x8f\xef\xbf\xba\xfe\xea\x3c\x63\xaf\x32\xce\x1e\x60\x0e\xab\x46\xc3\x12\x8a\x33\x92\x9f\x6d\x4a\x97\x5d\xaf\xa1\xfe\xd4\x49\x6b\xaf\x83\x1a\x4b\xae\x5d\x32\x5d\xd4\x0d\x04\xc5\xdf\xa7\xf5\xea\x83\x9b\xa5\xbe\x05\x55\x9b\x08\x2a\x80\xec\x22\xb6\x92\x71\x8f\xdb\x71\x77\x64\x6b\x25\x77\xf1\xc0\xa3\x4b\x8b\x95\x74\xec\xce\xed\xeb\x07\x3d\x21\x0d\x9c\x79\xd1\x35\xb0\x18\x55\xbf\xf2\x74\x6f\x79\x4c\x4f\xdd\xdb\xfb\x1d\x3e\x37\x4d\x43\x15\xc0\x98\xa4\xc7\x97\xdb\x81\x73\x59\xe1\x45\x75\xc1\x2f\x28\x22\x5b\x8a\xb2\x45\xec\xff\x8f\xc0\x37\xfc\x04\x00\x00\xff\xff\x19\x12\x00\x55\x00\x01\x00\x00"),
		},
		"/app/config/logger.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "logger.go.tmpl",
			modTime:          time.Date(2019, 3, 28, 18, 31, 25, 0, time.UTC),
			uncompressedSize: 169,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcc\xbd\xaa\xc2\x40\x10\x40\xe1\x7a\xe7\x29\x86\x54\xbb\x4d\xf2\x04\xb7\xba\x36\x82\x68\x69\x1d\x97\xc9\x38\x98\xec\x84\xc9\xac\x82\x92\x77\x17\x7f\x5a\xcb\xc3\x07\x67\xee\xf3\xa5\x67\xc2\xac\x65\x10\x06\x90\x69\x56\x73\x8c\x10\x1a\x16\x3f\xd7\x53\x9b\x75\xea\x6c\xe9\xee\x64\x3a\x2a\x37\xbf\xa0\x7b\x63\x02\x18\x6a\xc9\x28\x45\x3c\x26\x7c\x40\x90\x01\xb7\xcb\x86\xae\x9f\x0a\xa3\x72\xbb\x53\x66\x32\xfc\xc3\x57\x1c\xaa\xcf\xd5\xe3\xf7\xd2\xee\xe9\xf6\xaf\x65\xd1\x91\x8e\x26\x4e\x16\x53\x82\xb0\xc2\x0a\xcf\x00\x00\x00\xff\xff\xad\x5d\xc7\xd5\xa9\x00\x00\x00"),
		},
		"/app/docker-compose.yml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "docker-compose.yml.tmpl",
			modTime:          time.Date(2019, 4, 1, 14, 30, 2, 0, time.UTC),
			uncompressedSize: 226,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcd\x4f\xcb\x83\x30\x0c\x06\xf0\x7b\x3f\x45\xf0\xfe\xbe\x87\xb9\x5d\x7a\x73\x38\x76\x54\xec\x60\x47\x29\x2e\x48\x41\xd3\x92\x86\x8e\x21\x7e\xf7\xa1\xfb\xc3\xf0\x12\xc8\x2f\x0f\x4f\x12\x72\x74\x9e\x34\x64\x79\xa6\x54\x44\x4e\xae\xc3\xa8\x95\x02\x08\x3e\x4a\xcf\xcb\x02\x00\xe0\x46\xdb\xa3\xfe\xe2\x6a\x9d\x27\xb1\x8e\x90\x5b\xb2\xe3\xf6\xc8\x18\xc5\xb2\x68\xb0\xc3\xdd\x3e\x5e\x16\x3c\xcb\xbb\x0f\xe0\x0f\x0e\xfb\x7c\xa7\x97\xb1\x0a\x52\x72\xec\x69\x44\x92\x4f\xa4\xae\xcc\xe5\xdc\x9c\x4c\x5b\x17\xc6\x5c\xab\xa6\xdc\xfc\xf8\x49\x94\x47\x0d\xd3\xf4\x5f\x84\x30\xcf\xed\x0d\x13\x0e\x3e\x2c\x55\xea\x19\x00\x00\xff\xff\x14\x32\x1a\xa7\xe2\x00\x00\x00"),
		},
		"/app/main.go.tmpl": &vfsgen۰FileInfo{
			name:    "main.go.tmpl",
			modTime: time.Date(2019, 4, 18, 19, 3, 58, 357407375, time.UTC),
			content: []byte("\x70\x61\x63\x6b\x61\x67\x65\x20\x6d\x61\x69\x6e\x0a\x0a\x69\x6d\x70\x6f\x72\x74\x20\x22\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x7d\x7d\x2f\x63\x6d\x64\x22\x0a\x0a\x66\x75\x6e\x63\x20\x6d\x61\x69\x6e\x28\x29\x20\x7b\x0a\x09\x63\x6d\x64\x2e\x45\x78\x65\x63\x75\x74\x65\x28\x29\x0a\x7d\x0a"),
		},
		"/app/refresh.yml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "refresh.yml.tmpl",
			modTime:          time.Date(2019, 4, 18, 19, 4, 17, 143767457, time.UTC),
			uncompressedSize: 225,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\x8c\x4b\x6a\x03\x31\x0c\x86\xf7\x3a\x85\xc8\xde\x43\xe8\xd2\xbb\x9e\xa3\x14\xa3\x19\x29\x53\x83\x2c\x19\x3f\x42\x43\xc8\xdd\x43\x66\x36\x3f\xfc\xaf\x8f\x6a\x4d\xcd\x7d\x44\x5c\x20\xef\xe6\x4d\x38\xdd\x5c\x59\x5a\x8f\x10\xf0\x2e\xc6\xde\x20\xa0\xfa\x7e\x6a\x87\x80\xa3\x54\x08\x68\xce\x92\x8a\xf3\x54\xf9\x84\x6b\xb6\x4f\x25\xa5\x2a\x0d\xe9\x90\x6d\xd3\xc9\xc2\x49\xfe\x87\x58\xcf\x6e\x07\x71\xd9\x1d\xd6\x99\x95\x53\xa5\xf1\x17\x8f\xdb\xe9\x59\x94\x1e\x11\xbf\xae\x57\xeb\xb0\x66\xa3\xf6\x48\x46\x45\x22\x3e\x9f\xcb\x77\xad\xaf\x57\x38\x86\xb0\x79\x29\x64\x9c\x6e\x4a\x7b\x8f\xf8\x73\xe9\xd2\xee\x72\xf9\x05\x31\x5a\x55\xd2\xe6\xea\xad\x47\x1c\x6d\x0a\xbc\x03\x00\x00\xff\xff\x92\xe5\x9c\xe5\xe1\x00\x00\x00"),
		},
		"/app/resources": &vfsgen۰DirInfo{
			name:    "resources",
			modTime: time.Date(2019, 4, 18, 19, 21, 10, 210369035, time.UTC),
		},
		"/app/resources/config": &vfsgen۰DirInfo{
			name:    "config",
			modTime: time.Date(2019, 4, 18, 19, 21, 10, 210368986, time.UTC),
		},
		"/app/resources/config/database.yml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "database.yml.tmpl",
			modTime:          time.Date(2019, 4, 1, 15, 18, 30, 0, time.UTC),
			uncompressedSize: 365,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xcd\xc1\x4a\x03\x31\x10\x06\xe0\x7b\x9e\x62\xc8\x03\x64\xdb\x6a\x11\x72\x32\x62\x6f\x82\x60\xd7\x73\x88\xcd\xa0\x85\xb8\x13\x32\xd9\x7a\x08\xf3\xee\x12\xa5\xad\x3d\x78\xea\x6d\xf8\x99\xff\xff\x22\x1e\x30\x51\xfe\xc4\xa9\x5a\x05\x10\xf7\x21\xe1\xae\x5a\xc8\xc4\xf5\xbd\x20\xf7\x2c\xd4\xf0\x16\x18\x2d\xb4\x66\x5c\xce\x22\xfe\x4f\x4b\x01\xcc\x8c\xe5\xa2\x91\x03\xf3\x17\x95\x78\x11\x7e\x10\x57\x0b\xcb\xd5\x9d\x59\x98\x85\x59\xf6\x37\xa2\x64\x61\xad\x54\x45\xfe\xd1\xe7\x92\x3a\xa2\x5b\xd3\x22\x38\x1d\x9e\x0b\xe8\x71\xb3\x1d\xfd\xa3\x1b\xdd\x83\xdb\x6e\xfc\xeb\xcb\x93\x06\x7d\x5c\xb5\xc3\x70\x3a\x8f\xc7\xfd\x49\xb0\xeb\xdb\x9b\xd5\xd0\x9a\x11\xf1\x5d\xf8\x9d\xd5\x22\x4a\xe5\x42\x71\xde\xd5\x3d\x4d\xff\xb1\x57\x8b\x67\xe2\xec\x7e\x07\x00\x00\xff\xff\xc6\x97\x97\x8f\x6d\x01\x00\x00"),
		},
		"/app/resources/migrations": &vfsgen۰DirInfo{
			name:    "migrations",
			modTime: time.Date(2019, 4, 18, 19, 20, 53, 806386068, time.UTC),
		},
		"/app/resources/resources.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "resources.go.tmpl",
			modTime:          time.Date(2019, 4, 18, 19, 20, 29, 649967541, time.UTC),
			uncompressedSize: 224,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\x3d\x0e\xc3\x20\x0c\x46\xe7\xf8\x14\x88\x29\x48\x55\x90\xba\x77\xea\xdc\x4e\xbd\x80\x83\x1c\x8a\x1a\xe2\xc8\x40\x7a\xfd\x8a\xa4\xff\x4c\x7c\x4f\xef\xc9\x33\xba\x1b\x7a\x52\x42\x89\x8b\x38\x4a\x00\x21\xce\x2c\x59\xb5\xd0\x68\x1f\xf2\xb5\xf4\x9d\xe3\x68\x3d\xf7\x65\x18\x70\x64\x5b\x13\xb1\xcb\x5e\x83\x01\x58\x50\xaa\x79\xe4\x69\x08\x5e\xd5\x77\x50\xab\xd0\x9d\xe9\xde\x6a\xb7\x72\xbd\x53\xba\xb3\xcf\xbf\x81\xe6\x14\xbc\x60\x0e\x3c\xa5\x5f\x3b\xbe\xf9\x56\x7c\x6d\x03\xcd\x85\xe2\x3c\x62\xa6\xf4\x77\x23\xbf\xf8\x16\x7d\xa6\x01\x03\x8f\x00\x00\x00\xff\xff\xb5\x30\xe4\x28\xe0\x00\x00\x00"),
		},
		"/app/resources/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(2019, 4, 18, 19, 21, 2, 592697454, time.UTC),
		},
		"/app/root": &vfsgen۰DirInfo{
			name:    "root",
			modTime: time.Date(2019, 4, 3, 17, 48, 18, 740046379, time.UTC),
		},
		"/app/root/root_handler.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "root_handler.go.tmpl",
			modTime:          time.Date(2019, 4, 3, 17, 48, 18, 738759160, time.UTC),
			uncompressedSize: 197,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcd\xb1\x8a\x84\x30\x10\xc6\xf1\xda\x79\x8a\x21\x95\x1e\xa2\xfd\xbd\xc0\xc1\x15\x77\xa0\x85\xc5\x71\x45\x36\x3b\x1b\x83\x9b\x4c\x76\x32\x62\x21\xbe\xfb\xb2\xb8\xd5\x07\x1f\xfc\xf8\x67\xeb\x16\xeb\x09\x85\x59\x01\x42\xcc\x2c\x8a\x35\x54\x26\x91\xf6\xb3\x6a\x36\x00\x95\xf1\x41\xe7\xf5\xd2\x39\x8e\x7d\x5e\x7c\xb4\x21\xf5\x25\x32\xeb\x6c\xa0\x01\xb8\xad\xc9\xe1\x17\x69\xbd\xe1\x4b\x74\x03\x95\xcc\xa9\xd0\x24\x41\x49\x5a\x14\xfc\x78\xff\x8f\x95\x8a\x36\xb8\x43\x75\xf2\x6e\xa0\x74\x25\xf9\x1e\x7f\x7f\xea\xad\x45\x69\x31\xda\xfc\x57\x54\x42\xf2\xff\xe7\xec\x26\x52\x29\xd6\x93\xf9\x44\x33\xd1\xdd\x71\x24\x54\xc6\xf1\xec\x1f\x0d\x1c\xf0\x0c\x00\x00\xff\xff\x27\x27\x2c\x65\xc5\x00\x00\x00"),
		},
		"/app/router": &vfsgen۰DirInfo{
			name:    "router",
			modTime: time.Date(2019, 4, 3, 20, 49, 51, 878401154, time.UTC),
		},
		"/app/router/middleware.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "middleware.go.tmpl",
			modTime:          time.Date(2019, 4, 3, 18, 1, 33, 993546186, time.UTC),
			uncompressedSize: 303,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4d\x4a\x04\x31\x10\x46\xd7\xa9\x53\x84\xac\x3a\x83\x26\x87\xd0\x8d\xe0\x80\x0b\x05\xb7\xe9\xfc\x16\xd3\x49\x85\x74\x62\x0b\x4d\xdf\x5d\x70\x84\x81\xc1\xed\x7b\x54\xbd\xaf\x1a\x7b\x31\xd1\xf3\x46\xa3\xfb\x06\x80\xb9\x52\xeb\x7c\x02\x26\x22\xf6\x34\x66\x65\x29\xeb\x48\x8f\x36\xa1\xb6\x09\x05\xb0\xbc\xf1\xff\x9d\xce\xe8\xdc\xe2\x37\xd3\xbc\xb8\xbf\x9f\x47\x08\x66\x21\x5d\xa9\xde\xb9\x7a\x89\xd9\x60\xd1\x6b\x26\xea\x49\x00\x30\xb1\xef\xea\x4c\x6e\x2c\xfe\x38\xb4\xa5\x12\x30\x0a\x90\x00\x61\x14\xcb\x6f\x8d\xa9\xf1\x93\x4d\xa8\xce\xe3\xfb\x81\xbb\x99\x9f\x2a\x55\xf5\x44\xa5\x78\xdb\x91\x8a\xe4\x3b\x30\x0c\xfc\xfa\x40\xbd\xac\xcf\xfe\x6b\xfa\x85\xac\xa9\x8f\xd5\x4f\x79\x53\xaf\x14\xa3\x6f\x12\xd8\x01\xf0\x47\xaf\x33\xd4\x1b\xd5\xf7\xcf\xf3\xad\xe5\x66\x29\xe1\x80\x9f\x00\x00\x00\xff\xff\x66\xaa\xd8\x77\x2f\x01\x00\x00"),
		},
		"/app/router/router.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "router.go.tmpl",
			modTime:          time.Date(2019, 4, 3, 20, 49, 51, 877693407, time.UTC),
			uncompressedSize: 353,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xc1\x4a\xc4\x30\x10\x86\xcf\x9d\xa7\x18\x02\x42\xaa\x6b\x72\x5f\xf0\x20\x9e\x04\x57\xc4\x7d\x82\x36\x9d\xa6\xc1\x6c\x26\x4c\xa7\x2e\xb8\xec\xbb\x4b\xeb\x45\xf6\xfc\xcd\x7c\xfc\xff\x5f\xbb\xf0\xd5\x45\x42\xe1\x45\x49\x00\xd2\xa9\xb2\x28\x5a\x68\x4c\x21\xf5\x93\x6a\x35\x00\x8d\x89\x49\xa7\xa5\x77\x81\x4f\x3e\xf2\x63\x98\x92\x0f\x53\x32\xb7\xa0\x5f\xc6\xb1\xcb\xec\x2b\xd7\x1b\x26\xb3\xff\x21\xe1\xcc\xd1\x67\x8e\x9b\xf2\x72\x71\x07\x1e\x96\x4c\xd7\xab\x0f\x5c\xc6\x14\x0d\xb4\x00\xe3\x52\x02\x1e\xb5\x13\xb5\x43\x8f\xf7\x95\xab\x7b\xe1\x52\x28\x68\xe2\xd2\x22\x89\xb0\xe0\x05\x1a\xc1\xfd\x13\x86\x29\xb9\x77\x3a\x7f\x6e\xe9\x6d\x0b\xd0\x9c\xd2\x30\x64\x3a\x77\x42\x56\x76\x38\xf4\x2d\x34\x5b\xb7\xd9\xca\x8a\x33\x47\xf7\x5a\x46\xb6\xad\x3b\xcc\x71\xb4\xe6\x48\xf2\x4d\x82\x39\xcd\x4a\x25\x95\x88\x5c\x70\x9b\xe0\x6e\x76\xce\x99\x1d\xfe\x45\x73\x1f\x2c\xba\xba\x48\x17\x29\xb8\xee\xe2\xde\xb6\x9f\xe7\x32\x6c\x0e\x6b\xf6\xe6\xe1\xdf\xf1\x0e\xa5\x85\x2b\xfc\x06\x00\x00\xff\xff\x45\xae\xf5\xd3\x61\x01\x00\x00"),
		},
		"/app/router/routes.go.tmpl": &vfsgen۰FileInfo{
			name:    "routes.go.tmpl",
			modTime: time.Date(2019, 4, 1, 15, 35, 19, 0, time.UTC),
			content: []byte("\x70\x61\x63\x6b\x61\x67\x65\x20\x72\x6f\x75\x74\x65\x72\x0a\x0a\x69\x6d\x70\x6f\x72\x74\x20\x28\x0a\x09\x22\x67\x69\x74\x68\x75\x62\x2e\x63\x6f\x6d\x2f\x67\x6f\x2d\x63\x68\x69\x2f\x63\x68\x69\x22\x0a\x0a\x09\x22\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x7d\x7d\x2f\x72\x6f\x6f\x74\x22\x0a\x29\x0a\x0a\x66\x75\x6e\x63\x20\x72\x6f\x75\x74\x65\x73\x28\x72\x20\x2a\x63\x68\x69\x2e\x4d\x75\x78\x29\x20\x7b\x0a\x09\x72\x2e\x47\x65\x74\x28\x22\x2f\x22\x2c\x20\x72\x6f\x6f\x74\x2e\x47\x65\x74\x29\x0a\x7d\x0a"),
		},
		"/app/store": &vfsgen۰DirInfo{
			name:    "store",
			modTime: time.Date(2019, 4, 18, 19, 6, 41, 677201390, time.UTC),
		},
		"/app/store/connection.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "connection.go.tmpl",
			modTime:          time.Date(2019, 4, 18, 19, 5, 58, 530573306, time.UTC),
			uncompressedSize: 734,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x41\x6b\xdc\x30\x10\x85\xcf\x9e\x5f\x31\xf5\x49\x2e\x8b\x7d\x2f\xec\x25\xdd\x16\x02\x29\x2d\xf4\xd6\x9b\x64\x8f\xb4\x22\xf2\x8c\x90\xe4\x4d\xc3\xe2\xff\x5e\xb4\x4e\x37\x74\x0b\x21\x47\x3d\xcd\x7b\xf3\xcd\x4c\xd4\xe3\xa3\x76\x84\xb9\x48\x22\x00\x3f\x47\x49\x05\x15\x34\xad\xf3\xe5\xb8\x98\x7e\x94\x79\x70\x62\x16\x6b\x75\x90\x21\x4a\x6c\xff\xfd\x8b\x8f\x6e\xd6\x9e\x87\x3c\x8b\x94\x63\x0b\xd0\xb4\xe7\x73\xff\x4d\xa6\x25\xd0\xba\x0e\xa3\xb0\xf5\xae\xbd\x51\x13\x65\x59\xd2\x48\xb9\x85\x0e\xe0\xa4\x13\x4e\x06\x3f\x46\x89\xfd\x67\x61\xa6\xb1\x78\x61\x80\x61\xc0\xc3\x1d\x26\x2a\x4b\xe2\x8c\x9a\x91\x7e\xfb\x5c\x3c\xbb\xdb\xd2\x1e\xef\x2d\x6a\x1c\xaf\x02\x1e\x75\x46\x96\x82\xcf\x54\xd0\x10\x31\xce\x7a\xa2\x5d\x4d\xd4\xc8\xf4\x84\xc2\x84\x4f\x3e\x04\x34\x84\x94\x8b\x36\xc1\xe7\x23\x4d\x3d\xd8\x85\x47\x3c\xdc\xa9\x0e\xd5\x4d\x93\x1d\x52\x4a\x92\x3a\x3c\x43\xe3\x6d\x05\xfe\xb0\x47\xf6\xa1\xbe\x9b\x0d\x12\x27\xb3\xab\x12\x34\x2b\x40\x43\x29\xe1\xa7\xfd\x5f\x2a\xd5\x5d\x6c\x55\xfc\xdf\xc7\x3e\x5c\xe2\x37\xe3\x4d\xd8\x0a\x1b\xd5\x35\x68\x03\xa9\xfe\x4a\xf8\x93\xca\x83\x38\x47\x49\x6d\x27\xe8\x7f\x51\x92\x1f\x12\x37\xb1\xdb\x8a\x0e\x64\x16\x87\x17\x18\xeb\x5d\x7f\x9f\x0f\x74\x52\x1d\x40\x63\x2f\x7d\x2b\xe7\xf5\x26\x75\xe6\x5a\xf4\x3d\x12\xab\x76\xd2\x45\x1b\x9d\xa9\x7f\x9e\x43\xfb\xd6\x0c\x57\xfc\xfa\xbd\xc7\xda\xf5\x41\xf4\xf4\x35\xc9\xac\xec\xbb\x8c\x75\xe0\x57\xf3\xcb\xe6\xd5\x0b\xf2\x17\x3e\xbd\x2b\xe4\x75\xa3\xb0\xc2\x9f\x00\x00\x00\xff\xff\x36\x28\x33\x8d\xde\x02\x00\x00"),
		},
		"/app/store/migrator.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "migrator.go.tmpl",
			modTime:          time.Date(2019, 4, 18, 19, 6, 41, 676171424, time.UTC),
			uncompressedSize: 692,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x8f\x31\x4f\xf3\x30\x10\x86\xe7\xdc\xaf\xb8\x2f\xc3\x27\xbb\xaa\x12\xb1\x22\x75\x01\xd6\x22\xa4\x0a\x31\x27\xc6\x49\xad\x36\x3e\xeb\x6c\xab\x41\x51\xfe\x3b\x72\xda\xa6\x65\x28\x30\x30\xda\xef\xdd\xf3\x3e\xe7\x2a\xb5\xab\x5a\x8d\x3e\x10\x6b\x00\xd3\x39\xe2\x80\x02\xb2\xbc\x35\x61\x1b\xeb\x42\x51\x57\xb6\x54\xc7\xa6\xa9\xf6\x54\x3a\x72\xf9\xd7\xcc\xed\xda\x52\x33\x13\xfb\x1c\x20\xcb\x87\xa1\x58\xd3\x7b\xdc\xeb\x71\x2c\x59\x7b\x8a\xac\xb4\xcf\x41\x02\x84\x0f\xa7\xf1\x85\xdc\xda\xb4\x5c\x05\x62\xf4\x81\xa3\x0a\x38\x40\x56\x53\x8f\x8e\x5c\x71\x8c\x0c\xd9\x07\xea\x61\x04\x68\xa2\x55\xf8\xac\x0f\xe7\x15\xa1\xc8\x5a\x5c\xa4\xd1\x47\xb2\x56\xab\x34\x2b\x51\x2c\xae\xb0\x4b\x9c\x6c\xe4\x89\x3b\x3d\xf1\x7e\x35\xf1\x67\xd4\xb1\x42\xcc\x82\x97\x66\xbf\xc4\x54\x22\x21\x33\xcd\xb4\xfa\x6f\x85\xd6\xec\x13\x2d\x63\x1d\x22\xdb\xf4\x9c\xa8\x90\x8d\x00\xe7\xcf\xff\x57\x0a\x43\x4d\xfd\xb8\x4c\x73\xf3\x11\xa2\xbb\x3e\x5d\xe2\xab\x13\xf2\x28\x9a\xc8\x27\xc5\xae\xa8\xa9\x2f\x52\xf4\x4d\xfb\x5c\x9c\xe2\xf3\xce\x26\x54\x21\xfa\xdf\xed\x5d\xae\xb8\x69\xf7\x44\x07\x7b\xd3\x6f\x0a\xef\x7e\xa8\x22\xf6\xc5\x9b\x09\xdb\x4d\xa8\xd4\x4e\x68\x66\xf9\x87\xd2\x9f\x01\x00\x00\xff\xff\xeb\xab\x5e\x03\xb4\x02\x00\x00"),
		},
		"/domain": &vfsgen۰DirInfo{
			name:    "domain",
			modTime: time.Date(2019, 4, 5, 17, 49, 34, 111917585, time.UTC),
		},
		"/domain/domain.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "domain.go.tmpl",
			modTime:          time.Date(2019, 4, 5, 17, 49, 22, 794104843, time.UTC),
			uncompressedSize: 892,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x6a\xdb\x40\x10\x3d\x6b\xbe\x62\xd0\x49\x6b\x8c\x7c\x0f\xe4\xd0\x3a\x2d\x04\x4a\x30\x6d\xdc\x8b\x31\xcd\x5a\x1a\xcb\x53\x4b\xbb\xdb\xd5\xac\xd3\x56\xec\xbf\x17\x59\x52\x68\x9d\x04\x1a\x5d\xc4\x68\xde\x9b\xf7\xde\xec\xca\xe9\xe2\xa8\x2b\xc2\xae\xcb\xbf\xb0\xa9\x42\xad\x7d\x8c\x00\xdc\x38\xeb\x05\x33\x48\x52\x43\xb2\x38\x88\xb8\x14\x92\x54\xb8\xa1\x14\x20\x49\x2b\x96\x43\xd8\xe5\x85\x6d\x16\x95\xdd\x85\xfd\x5e\xd7\x76\xe1\xec\x19\xf4\x62\x2f\x04\x2e\x5f\x6d\x9e\x74\xcd\xa5\x16\xba\x00\xb8\x63\xd5\x68\x36\x8b\xb6\xb1\x56\x0e\x0b\xa7\x2b\x36\x55\x0a\x0a\x40\x7e\xb9\xb3\xe5\xa5\x76\x2c\xba\xe6\xdf\x54\xc6\x88\xad\xf8\x50\x08\x76\x90\xdc\xde\xe0\xf8\xf4\xba\xf9\x7a\x7d\x7b\x83\x0f\xdf\x5b\x6b\xae\x52\x2e\x53\x2c\x77\xe7\xf7\x03\x24\x4b\x4f\x5a\xa8\x7c\x27\xd8\x47\xcb\xef\xb9\xa1\x09\x58\x4c\xad\x01\x3f\x96\xdf\xb4\xf4\xbc\xb5\x2b\x5f\xe3\x85\xa9\x35\xf0\xc6\x72\xe0\x45\x80\x7d\x30\x05\x66\x5d\x97\x7f\x64\xdf\xca\x27\x12\x21\x1f\xe3\xb3\x30\x0a\xdf\xb3\x29\x33\x8f\xb3\x7e\xf7\xf9\x67\xfa\x11\xa8\x15\x85\xe4\xbd\xf5\x7d\x44\x4f\x12\xbc\x41\xc3\xf5\x9b\xa6\x7e\x1d\x57\x9d\xc9\x4f\x9c\x39\xeb\xf2\xa5\x35\x86\x0a\x61\x6b\x14\x66\xb3\xe9\x24\xf2\x0f\xbd\x4e\x3b\x1f\xf4\xd4\x5f\x82\x4f\x88\x3b\x7a\x1c\x40\x99\x9a\x4f\x36\x5e\x38\x97\x55\x1d\xbc\xae\x63\xc4\xcd\xf6\xd2\xcd\x7f\xda\x9e\x26\x28\xbc\xb7\x2b\x5d\x51\x76\xbe\x08\x5a\xac\x1f\x12\xac\xa6\x52\xe1\x6c\xb8\x23\xfd\x27\xea\x3d\x17\xd6\x08\x19\xc1\xab\x6b\x6c\xf4\x91\xb2\xcd\x96\x8d\x90\xdf\xeb\x82\xba\x38\xc7\x9a\xcc\x33\x71\xa5\x20\xd9\x5b\x8f\x3c\xc7\x53\xcf\xf3\xda\x0c\xbf\xc7\x85\x45\x48\xa6\xe9\x1b\xde\xe2\x35\x9e\x20\x89\xf0\xb4\xa5\xd1\xc7\x1d\x3d\xfe\xeb\x78\x8e\x23\x49\x41\x84\x3f\x01\x00\x00\xff\xff\xe8\x4b\xec\x5f\x7c\x03\x00\x00"),
		},
		"/domain/domain_handler.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "domain_handler.go.tmpl",
			modTime:          time.Date(2019, 4, 5, 17, 41, 57, 910329883, time.UTC),
			uncompressedSize: 1766,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x94\x51\x6f\xd3\x30\x10\xc7\x9f\xe3\x4f\x61\xf2\x30\x39\x28\x24\xef\x48\x7b\x60\x2d\x1d\x20\xd4\x45\x2b\xd5\x78\x35\xcd\xcd\xb1\xd6\xd8\xe6\x7c\x51\x35\xa2\x7c\x77\xe4\xc4\x50\xb2\x22\xa4\x49\x30\x04\x8f\xc9\x9d\xef\xfe\xbf\xbf\xef\xec\xe4\xee\x4e\x2a\xe0\x7d\x5f\x6c\xb4\x51\xdd\x5e\xe2\x30\x30\xa6\x5b\x67\x91\xb8\x60\x49\x6a\x80\xca\x86\xc8\xa5\x8c\x25\xa9\xd2\xd4\x74\x9f\x8a\x9d\x6d\x4b\x65\x5f\xec\x1a\x5d\x22\x98\x1a\x30\x9d\xc7\xdc\x9d\x6a\xa5\x36\xa5\x6f\xad\xa5\xe6\x97\xc1\xd2\x49\xa5\x8d\x4a\x59\xc6\xd8\x6d\x67\x76\xbc\xb2\x9e\xc4\x81\x87\x96\xc5\x35\x78\x67\x8d\x87\x1b\xd4\x04\x98\x73\xe4\xcf\xe3\xff\xcf\x1d\x78\xca\x78\xcf\x92\x99\x72\xfe\xf2\x9c\x1b\x38\x88\xbe\x2f\x16\xd2\x69\x92\x7b\xfd\x05\xea\x61\xc8\x58\x02\x88\x21\x3a\xe9\x2d\x2e\xb4\xa9\x05\xe6\x73\xee\x8c\x25\xfa\x96\x87\xc4\x67\xe7\xdc\xe8\x7d\x28\x9f\x4c\x32\x8b\xeb\xf1\xdc\xbb\xcd\xd5\xfa\x35\xa2\x45\x71\xc8\x39\xe6\x21\x37\x63\x49\x82\x40\x1d\x1a\x96\x0c\x6c\xea\x73\xce\x17\x08\x92\x40\xcc\xca\xe7\x3c\xd6\xba\x04\xfa\xf0\x51\x60\xf6\x5b\xfa\x9d\xe4\xdf\x68\x6a\x36\x24\xa9\xf3\xf1\xd0\x68\xd9\xf4\x67\x92\x55\x9f\x70\x0f\xd1\xfc\x4b\x78\x94\xf7\x4e\x2a\x18\x45\x05\x67\x57\xda\xd4\x95\x54\x20\xa6\x1b\x2d\xd6\x70\xa8\x24\xca\xd6\x0b\xcc\x9e\x08\x3d\xa6\x06\x59\x33\xa6\x2b\x03\x8f\xc1\xd2\xf5\xd8\xce\x07\xaa\xa3\xee\xed\xf6\xed\x72\x85\xb6\xad\x24\x35\x61\x74\x52\x5d\xa7\xdf\x31\xfc\x37\x8e\xb3\xb3\xf1\xb3\x78\x23\xfd\x2b\x73\x2f\xb2\x53\xb0\x0b\x59\xc7\x66\x47\x32\xff\x10\xad\xef\x8b\x95\x46\x4f\xef\x81\x08\xc6\xd9\xf9\xc1\xe5\x80\x13\x34\x3e\xa5\xa7\x0f\x05\x1d\xfd\xad\xba\xff\x60\x5f\xb7\xae\xfe\x5b\xfb\x7a\x34\xf8\x67\x1b\xb9\x84\x3d\xd0\xbf\x36\xbd\xf1\xee\x96\xe0\x09\xed\xfd\x1f\x9b\xd5\x38\x1b\xf1\xb1\x9b\xbf\x74\x6b\xbb\xb0\x86\xc0\x50\xb0\xf2\x6b\x00\x00\x00\xff\xff\x5b\xca\x59\xc8\xe6\x06\x00\x00"),
		},
		"/domain/domain_repo.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "domain_repo.go.tmpl",
			modTime:          time.Date(2019, 4, 5, 17, 49, 34, 110990098, time.UTC),
			uncompressedSize: 1388,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x41\x6f\xd3\x30\x14\xc7\xcf\xf1\xa7\x78\xe4\x80\x92\x69\x72\xc4\x15\x31\xa1\xa9\xdb\x24\x2e\x50\x36\x2a\x0e\x88\x83\x9b\xbc\xa4\x56\x1d\xdb\x7d\xb6\xd5\x8e\x28\xdf\x1d\xa5\x6d\xba\x36\x4d\x01\x4d\x9c\x9f\xfd\xff\xbd\xdf\xb3\x9f\x15\xf9\x52\x54\x08\x4d\xc3\x9f\xa4\xae\x82\x12\xd4\xb6\x8c\xc9\xda\x1a\xf2\x90\xb0\x28\x2e\x84\x17\x73\xe1\x30\x73\x2b\x15\x33\x16\xc5\x95\xf4\x8b\x30\xe7\xb9\xa9\xb3\xca\xcc\x43\x59\x0a\x65\x32\x6b\x6c\x7c\xa9\x16\x82\x2c\x06\x45\xbb\xac\x32\x24\x32\xe4\xce\x0b\xb5\x90\x3a\x73\xb5\x31\x7e\x91\x59\x51\x49\x5d\xc5\x2c\x65\xac\x0c\x3a\x87\x09\xa1\xf0\x98\x9c\x74\x0b\x57\x4d\xc3\x27\xc2\x4a\x2f\x94\xfc\x85\x45\xdb\x5e\x83\xdf\xc0\x95\x35\x96\x4f\x8c\xd6\x98\x7b\x69\x74\x0a\x5b\x1e\x34\x2c\x22\xf4\x81\x34\xf8\x0d\x1f\x8b\x4b\x59\xbb\x87\x3d\x48\x5d\x7c\xd1\x98\xc8\x02\x3a\x05\x3e\x9b\x7d\xba\x1b\x8f\x4e\x46\x5a\xd8\xe2\xd2\x8e\x77\xda\xed\xfb\x1b\xd0\xb8\x4e\x86\x17\x52\x16\x21\x51\x57\xf5\x1b\xde\x91\x4f\xbb\xba\x06\x59\xa4\x2c\x92\x65\x97\x0b\x6f\x6e\x40\x4b\xd5\x65\xf7\x32\x5a\xaa\x2d\x92\x45\x2d\x3b\x18\x0e\x12\xb4\x54\x27\x6e\x53\x51\x61\x62\xe1\x6a\x37\x64\x3e\x15\x24\x6a\x77\x49\xf0\x70\xa8\xc2\x81\xdb\x54\x05\x12\xea\x92\x59\x5f\x4d\x19\x8b\x56\x7b\xbd\x69\x97\x25\x3c\x3e\x90\xa9\x77\xd4\xc4\xbe\xf8\xaf\xf8\xad\x52\xc9\x51\xf0\x6b\xbc\xfb\xbb\xfc\x9b\xd9\x6a\xae\x7a\xa8\xa1\xf4\x7c\x12\x1d\x70\x54\xfb\xc7\xcf\xbf\x3c\xec\xb1\xfc\xf9\xe1\xa6\x3d\x7e\xd5\x8e\xf2\xf6\x3f\x79\xed\x8a\xbd\xc3\xcc\x16\xaf\x5c\x8b\x6b\xc0\x4d\xae\x42\x81\x13\xa3\x42\xad\x1d\x70\xce\x9d\x27\xa9\xab\xd1\x85\x19\x03\x0d\x23\x38\xe7\x2f\x3b\x74\x87\xce\x93\x79\xfe\x87\x1d\x3a\xd0\xfa\x5f\xf2\x28\xd6\x5f\x03\xd2\x73\x12\x17\xa8\xd0\x23\x94\x64\xea\xe3\x19\xc0\x7a\x81\x84\x20\x0b\xb8\x81\x8f\xf1\x7e\x43\xc8\xac\xdd\x6d\x59\x62\xee\xb1\xd8\x0e\x69\xf7\xa5\xee\x37\x98\x7f\x97\x7e\x31\x31\x41\xfb\xe4\x4f\x93\x3f\x0c\x5d\x96\x70\x1c\x06\x1f\xe0\xdd\xe0\x9c\x21\xc7\xbb\xd0\x27\x2f\xf2\x65\xe2\x56\x8a\xdf\x13\x7d\x36\x8f\x66\xed\xd2\x93\x87\xdb\xfd\xb7\xdf\x01\x00\x00\xff\xff\x23\xba\x39\x8a\x6c\x05\x00\x00"),
		},
		"/migration": &vfsgen۰DirInfo{
			name:    "migration",
			modTime: time.Date(2019, 4, 3, 15, 20, 42, 364966851, time.UTC),
		},
		"/migration/down.sql.tmpl": &vfsgen۰FileInfo{
			name:    "down.sql.tmpl",
			modTime: time.Date(2019, 4, 1, 21, 47, 19, 0, time.UTC),
			content: []byte("\x64\x72\x6f\x70\x20\x74\x61\x62\x6c\x65\x20\x7b\x7b\x2e\x50\x6c\x75\x72\x61\x6c\x7d\x7d\x3b\x0a"),
		},
		"/migration/up.sql.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "up.sql.tmpl",
			modTime:          time.Date(2019, 4, 1, 21, 46, 56, 0, time.UTC),
			uncompressedSize: 200,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8c\x4b\x0a\x02\x41\x0c\x44\xf7\x7d\x8a\x5a\x2a\x88\x17\xf0\x12\xde\x40\xa2\x1d\x30\x98\xfe\x10\x13\x64\x1c\xe6\xee\x82\x8d\x22\xee\xe6\xed\xaa\xe0\xbd\x8b\x31\x39\xc3\xe9\xac\x8c\x79\xde\x1f\x35\x8c\x74\x59\xd2\x26\x01\x80\x64\x7c\x88\xf8\x19\x7f\xd4\xe6\xa8\xa1\x8a\x6e\x52\xc8\x26\xdc\x78\xda\xbd\x03\xa3\x9f\x4f\xe4\x70\x29\x7c\x77\x2a\x1d\x0f\xf1\x6b\x8b\xf1\xe0\xd9\x2a\x7f\x03\x43\x8a\x9e\xd7\x48\x69\x7b\x48\xaf\x00\x00\x00\xff\xff\xb9\x08\x4a\xb0\xc8\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app"].(os.FileInfo),
		fs["/domain"].(os.FileInfo),
		fs["/migration"].(os.FileInfo),
	}
	fs["/app"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/.gitignore.tmpl"].(os.FileInfo),
		fs["/app/cmd"].(os.FileInfo),
		fs["/app/config"].(os.FileInfo),
		fs["/app/docker-compose.yml.tmpl"].(os.FileInfo),
		fs["/app/main.go.tmpl"].(os.FileInfo),
		fs["/app/refresh.yml.tmpl"].(os.FileInfo),
		fs["/app/resources"].(os.FileInfo),
		fs["/app/root"].(os.FileInfo),
		fs["/app/router"].(os.FileInfo),
		fs["/app/store"].(os.FileInfo),
	}
	fs["/app/cmd"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/cmd/migrate.go.tmpl"].(os.FileInfo),
		fs["/app/cmd/root.go.tmpl"].(os.FileInfo),
		fs["/app/cmd/serve.go.tmpl"].(os.FileInfo),
	}
	fs["/app/config"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/config/env.go.tmpl"].(os.FileInfo),
		fs["/app/config/logger.go.tmpl"].(os.FileInfo),
	}
	fs["/app/resources"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/resources/config"].(os.FileInfo),
		fs["/app/resources/migrations"].(os.FileInfo),
		fs["/app/resources/resources.go.tmpl"].(os.FileInfo),
		fs["/app/resources/templates"].(os.FileInfo),
	}
	fs["/app/resources/config"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/resources/config/database.yml.tmpl"].(os.FileInfo),
	}
	fs["/app/root"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/root/root_handler.go.tmpl"].(os.FileInfo),
	}
	fs["/app/router"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/router/middleware.go.tmpl"].(os.FileInfo),
		fs["/app/router/router.go.tmpl"].(os.FileInfo),
		fs["/app/router/routes.go.tmpl"].(os.FileInfo),
	}
	fs["/app/store"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/store/connection.go.tmpl"].(os.FileInfo),
		fs["/app/store/migrator.go.tmpl"].(os.FileInfo),
	}
	fs["/domain"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/domain/domain.go.tmpl"].(os.FileInfo),
		fs["/domain/domain_handler.go.tmpl"].(os.FileInfo),
		fs["/domain/domain_repo.go.tmpl"].(os.FileInfo),
	}
	fs["/migration"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/migration/down.sql.tmpl"].(os.FileInfo),
		fs["/migration/up.sql.tmpl"].(os.FileInfo),
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
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
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

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
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