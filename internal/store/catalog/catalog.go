package catalog

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
)

type FS interface {
	fs.StatFS
	Mkdir(name string, perm fs.FileMode) error
	Touch(name string, perm fs.FileMode) error
}

type Catalog struct {
	sys   FS
	root  string
	marks string
}

const (
	MarksFilename = "marks"
	Perm          = 0755
)

func New(root string, sys FS) Catalog {
	marks := filepath.Join(root, MarksFilename)
	return Catalog{
		root:  root,
		marks: marks,
		sys:   sys,
	}
}

func (c Catalog) Marks() string {
	return c.marks
}

func (c Catalog) Root() string {
	return c.root
}

func (c Catalog) EnsureRoot() error {
	exists, err := c.rootExists()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return c.sys.Mkdir(c.root, Perm)
}

func (c Catalog) rootExists() (bool, error) {
	info, err := fs.Stat(c.sys, c.root)
	if err != nil {
		return false, assertNotFound(err)
	}
	if !info.IsDir() {
		return false, fmt.Errorf("%s should be a dir, file found", c.root)
	}
	return true, nil
}

func (c Catalog) EnsureMarks() error {
	exists, err := c.marksExists()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return c.sys.Touch(c.marks, Perm)
}

func (c Catalog) FindRecord(prefix string) (string, error) {
	file, err := c.sys.Open(c.marks)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		record := scanner.Text()
		if strings.HasPrefix(record, prefix) {
			return record, nil
		}
	}
	return "", commands.ErrNotFound
}

func (c Catalog) marksExists() (bool, error) {
	if exists, err := c.rootExists(); err != nil || !exists {
		return exists, err
	}
	info, err := fs.Stat(c.sys, c.marks)
	if err != nil {
		return false, assertNotFound(err)
	}
	if info.IsDir() {
		return false, fmt.Errorf("%s should be a file, dir found", c.marks)
	}
	return true, nil
}

func assertNotFound(err error) error {
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}
