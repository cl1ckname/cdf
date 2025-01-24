package filesystem

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Filesystem struct{}

func (f Filesystem) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (f Filesystem) Mkdir(name string, perm fs.FileMode) error {
	return os.Mkdir(name, perm)
}

func (f Filesystem) Touch(name string, perm fs.FileMode) error {
	return os.WriteFile(name, nil, perm)
}

func (f Filesystem) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (f Filesystem) Abs(name string) (string, error) {
	return filepath.Abs(name)
}

func (f Filesystem) OpenFile(path string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(path, flag, perm)
}

var FS Filesystem
