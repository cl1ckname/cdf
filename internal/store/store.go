package store

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/cl1ckname/cdf/internal/collection/dict"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

const (
	MarksFilename = "marks"
	ReplaceFlag   = os.O_TRUNC | os.O_WRONLY
	AppendFlag    = os.O_APPEND | os.O_WRONLY
	ReadFlag      = os.O_RDONLY
	Perm          = 0666
)

type Catalog interface {
	EnsureRoot() error
	EnsureMarks() error
}

type FS interface {
	Stat(path string) (fs.FileInfo, error)
	Abs(path string) (string, error)
}

type Filestore struct {
	FS
	base string
}

func New(sys FS, base string) Filestore {
	return Filestore{
		FS:   sys,
		base: base,
	}
}

func Init(dir Catalog) error {
	if err := dir.EnsureRoot(); err != nil {
		return err
	}
	if err := dir.EnsureMarks(); err != nil {
		return err
	}
	return nil
}

func (f Filestore) Load() (domain.Collection, error) {
	file, err := readOpen(f.marks())
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	marks := make(map[string]domain.Mark)
	var mark domain.Mark
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
		}
		if err := json.Unmarshal(line, &mark); err != nil {
			return nil, err
		}
		marks[mark.Alias] = mark
	}
	return dict.Dict(marks), nil
}

func (f Filestore) WriteTo(to, value string) error {
	dst, err := appendOpen(to)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = dst.WriteString(value)
	return err
}

func (f Filestore) Save(marks domain.Collection) error {
	dst, err := replaceOpen(f.marks())
	if err != nil {
		return err
	}
	defer dst.Close()
	for mark := range marks.Iterate() {
		rec := NewRecord(mark)
		if err := rec.Write(dst); err != nil {
			return err
		}
	}
	if err := dst.Close(); err != nil {
		return err
	}

	return nil
}

func (f Filestore) marks() string {
	return filepath.Join(f.base, MarksFilename)
}

func appendOpen(path string) (*os.File, error) {
	return openWithFlag(path, AppendFlag)
}

func readOpen(path string) (*os.File, error) {
	return openWithFlag(path, ReadFlag)
}

func replaceOpen(path string) (*os.File, error) {
	return openWithFlag(path, ReplaceFlag)
}

func openWithFlag(path string, flag int) (*os.File, error) {
	file, err := os.OpenFile(path, flag, Perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}
