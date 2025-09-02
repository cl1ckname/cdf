package store

import (
	"encoding/json"
	"io/fs"
	"os"

	"github.com/cl1ckname/cdf/internal/config"
	"github.com/cl1ckname/cdf/internal/logger"
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

func (f Filestore) Load() (*config.Config, error) {
	file, err := f.readOpen(f.file)
	if err != nil {
		return nil, err
	}
	var cfg config.Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	if err := file.Close(); err != nil {
		return nil, err
	}
	if cfg.Marks == nil {
		cfg.Marks = domain.Dict{}
	}
	return &cfg, nil
}

func (f Filestore) Save(cfg config.Config) error {
	dst, err := f.replaceOpen(f.file)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(dst).Encode(cfg); err != nil {
		return err
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
