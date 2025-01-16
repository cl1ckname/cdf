package store

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"

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

func (f Filestore) Load() ([]string, error) {
	file, err := readOpen(f.marks())
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	var marks []string
	for {
		path, err := reader.ReadString('\n')
		if err == nil {
			marks = append(marks, path)
			continue
		}
		if errors.Is(err, io.EOF) {
			break
		}
		return nil, err
	}
	return marks, nil
}

func (f Filestore) Append(mark domain.Mark) error {
	file, err := appendOpen(f.marks())
	if err != nil {
		return err
	}
	rec := NewRecord(mark)
	if err = rec.Write(file); err != nil {
		return err
	}
	return nil
}

func (f Filestore) Find(alias string) (domain.Mark, error) {
	record, err := f.findRecord(alias)
	if err != nil {
		return domain.Mark{}, err
	}

	return NewMark(*record), nil
}

func (f Filestore) findRecord(prefix string) (*Record, error) {
	file, err := readOpen(f.marks())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := json.NewDecoder(file)
	var rec Record
	for scanner.More() {
		if err := scanner.Decode(&rec); err != nil {
			return nil, err
		}
		if rec.Alias == prefix {
			return &rec, nil
		}
	}
	return nil, domain.ErrNotFound
}

func (f Filestore) List() ([]domain.Mark, error) {
	file, err := readOpen(f.marks())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := json.NewDecoder(file)
	var marks []domain.Mark
	var rec Record
	for scanner.More() {
		if err := scanner.Decode(&rec); err != nil {
			return nil, err
		}
		mark := NewMark(rec)
		marks = append(marks, mark)
	}

	return marks, nil
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

func (f Filestore) Replace(marks []domain.Mark) error {
	dst, err := replaceOpen(f.marks())
	if err != nil {
		return err
	}
	defer dst.Close()
	for _, mark := range marks {
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
