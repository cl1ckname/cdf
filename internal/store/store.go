package store

import (
	"bufio"
	"errors"
	"io"
	"os"
)

const (
	AppendFlag = os.O_APPEND | os.O_WRONLY
	ReadFlag   = os.O_WRONLY
	Perm       = 0666
)

type Catalog interface {
	Root() string
	EnsureRoot() error
	Marks() string
	EnsureMarks() error
}

type Filestore struct {
	dir Catalog
}

func New(dir Catalog) Filestore {
	return Filestore{
		dir: dir,
	}
}

func (f Filestore) Init() error {
	if err := f.dir.EnsureRoot(); err != nil {
		return err
	}
	if err := f.dir.EnsureMarks(); err != nil {
		return err
	}
	return nil
}

func (f Filestore) Load() ([]string, error) {
	file, err := readOpenOrCreate(f.dir.Marks())
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

func (f Filestore) Append(mark string) error {
	file, err := appendOpenOrCreate(f.dir.Marks())
	if err != nil {
		return err
	}
	_, err = file.WriteString(mark + "\n")
	if err != nil {
		return err
	}
	return nil
}

func appendOpenOrCreate(path string) (*os.File, error) {
	return openOrCreateWithFlag(path, AppendFlag)
}

func readOpenOrCreate(path string) (*os.File, error) {
	return openOrCreateWithFlag(path, ReadFlag)
}

func openOrCreateWithFlag(path string, flag int) (*os.File, error) {
	file, err := os.OpenFile(path, flag, Perm)
	if err == nil {
		return file, nil
	}
	file, err = os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil

}
