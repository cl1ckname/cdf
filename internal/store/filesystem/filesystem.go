package filesystem

import (
	"io/fs"
	"os"
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

var FS Filesystem