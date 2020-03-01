package filesystem

import (
	"net/http"
	"path/filepath"
)

// FileSystem .
type FileSystem struct {
	http.FileSystem
	root      string
	routeSize int
	http.Handler
}

// New .
func New(httpfs http.FileSystem) *FileSystem {
	fs := &FileSystem{
		FileSystem: httpfs,
	}
	fs.Handler = http.FileServer(fs)
	return fs
}

// SetRoute .
func (fs *FileSystem) SetRoute(route string) *FileSystem {
	fs.routeSize = len(route)
	return fs
}

// SetRoot .
func (fs *FileSystem) SetRoot(root string) *FileSystem {
	fs.root = root
	return fs
}

// HTTPHandler .
func (fs *FileSystem) HTTPHandler() http.Handler {
	return fs.Handler
}

// Sub .
func (fs *FileSystem) Sub(root, route string) *FileSystem {
	subfs := &FileSystem{
		FileSystem: fs.FileSystem,
		root:       filepath.Join(fs.root, root),
		routeSize:  len(route),
	}
	subfs.Handler = http.FileServer(fs)
	return subfs
}

// Open .
func (fs *FileSystem) Open(name string) (http.File, error) {
	return fs.FileSystem.Open(filepath.Join("/", fs.root, name[fs.routeSize:]))
}
