package noindexfs

import (
	"net/http"
	"os"
)

type File struct {
	http.File
}

func (f File) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

type Filesystem struct {
	fs http.FileSystem
}

func New(fs http.FileSystem) Filesystem {
	return Filesystem{fs: fs}
}

func (fs Filesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return File{f}, nil
}
