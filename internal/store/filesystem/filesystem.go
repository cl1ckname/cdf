// Package filesystem contains store.FS interface implementation
package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var ErrInvalidPath = errors.New("invalid path")

type Filesystem struct {
	safePath string
}

func (f Filesystem) Open(name string) (fs.File, error) {
	return os.Open(filepath.Clean(name))
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
	path = filepath.Clean(path)
	if !strings.HasPrefix(path, f.safePath) {
		return nil, fmt.Errorf("go to %s: %w", path, ErrInvalidPath)
	}
	return os.OpenFile(path, flag, perm)
}

func New(safePath string) Filesystem {
	return Filesystem{
		safePath: safePath,
	}
}
