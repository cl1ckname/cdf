package catalog

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/cl1ckname/cdf/internal/logger"
)

const (
	MarksFilename = "marks"
	Perm          = 0755
)

type FS interface {
	fs.StatFS
	Touch(name string, perm fs.FileMode) error
	Mkdir(name string, perm fs.FileMode) error
}

func EnsureFile(log logger.Logger, filepath string, fs FS) error {
	exists, err := marksExists(filepath, fs)
	if err != nil {
		return err
	}
	if exists {
		log.Info("marks file found at", filepath)
		return nil
	}
	log.Info("marks file at", filepath, "does not exist, createing")
	return fs.Touch(filepath, Perm)
}

func InitInFolder(log logger.Logger, path string, fs FS) (string, error) {
	if err := ensureRoot(log, fs, path); err != nil {
		return "", err
	}
	marksFilePath := filepath.Join(path, MarksFilename)
	if err := EnsureFile(log, marksFilePath, fs); err != nil {
		return "", err
	}
	return marksFilePath, nil
}

func ensureRoot(log logger.Logger, sys FS, root string) error {
	exists, err := rootExists(sys, root)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	log.Info("cdf data folder does not exist, creating")
	return sys.Mkdir(root, Perm)
}

func rootExists(sys FS, root string) (bool, error) {
	info, err := fs.Stat(sys, root)
	if err != nil {
		return false, assertNotFound(err)
	}
	if !info.IsDir() {
		return false, fmt.Errorf("%s should be a dir, file found", root)
	}
	return true, nil
}

func marksExists(filepath string, fs FS) (bool, error) {
	info, err := fs.Stat(filepath)
	if err != nil {
		return false, assertNotFound(err)
	}
	if info.IsDir() {
		return false, fmt.Errorf("%s should be a file, dir found", filepath)
	}
	return true, nil
}

func assertNotFound(err error) error {
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}
