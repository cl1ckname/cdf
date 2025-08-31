package store

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"os"

	"github.com/cl1ckname/cdf/internal/logger"
	"github.com/cl1ckname/cdf/internal/pkg/dict"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

const (
	ReplaceFlag = os.O_TRUNC | os.O_WRONLY
	AppendFlag  = os.O_APPEND | os.O_WRONLY
	ReadFlag    = os.O_RDONLY
	Perm        = 0o666
)

type FS interface {
	Stat(path string) (fs.FileInfo, error)
	Abs(path string) (string, error)
	OpenFile(path string, flag int, perm fs.FileMode) (*os.File, error)
	Cwd() (string, error)
}

type Filestore struct {
	FS
	file string
	log  logger.Logger
}

func New(sys FS, file string, log logger.Logger) Filestore {
	return Filestore{
		FS:   sys,
		file: file,
		log:  log,
	}
}

func (f Filestore) Load() (dict.Dict, error) {
	file, err := f.readOpen(f.file)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	marks := make(map[string]domain.Mark)
	for i := 1; true; i++ {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
		}
		rec, err := ParseRecord(line)
		if err != nil {
			f.log.Warning("error while reading line", i, ":", err.Error())
			continue
		}
		marks[rec.Alias] = NewMark(rec)
	}
	return dict.Dict(marks), nil
}

func (f Filestore) Save(marks dict.Dict) error {
	dst, err := f.replaceOpen(f.file)
	if err != nil {
		return err
	}
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

func (f Filestore) Cwd() (string, error) {
	return f.FS.Cwd()
}

func (f Filestore) readOpen(path string) (*os.File, error) {
	return f.openWithFlag(path, ReadFlag)
}

func (f Filestore) replaceOpen(path string) (*os.File, error) {
	return f.openWithFlag(path, ReplaceFlag)
}

func (f Filestore) openWithFlag(path string, flag int) (*os.File, error) {
	file, err := f.OpenFile(path, flag, Perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}
