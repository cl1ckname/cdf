package mover

import (
	"io/fs"
	"os"
)

const Perm = 0666
const Mode = os.O_TRUNC | os.O_WRONLY

type FS interface {
	OpenFile(path string, flag int, perm fs.FileMode) (*os.File, error)
}

type Mover struct {
	fs FS
}

func NewMover(fs FS) Mover {
	return Mover{
		fs: fs,
	}
}

func (m Mover) WriteTo(path, content string) error {
	file, err := m.fs.OpenFile(path, Mode, Perm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
