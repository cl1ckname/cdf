package fabrics

import (
	"fmt"
	"io/fs"
	"time"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type FS interface {
	Stat(path string) (fs.FileInfo, error)
	Abs(path string) (string, error)
}

type Clock interface {
	Now() time.Time
}

type Marks struct {
	fs    FS
	clock Clock
}

func NewMarks(fs FS, clock Clock) Marks {
	return Marks{
		fs:    fs,
		clock: clock,
	}
}

func (b Marks) Build(alias, path string) (*domain.Mark, error) {
	info, err := b.fs.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("path could be and folder, not file: %w", domain.ErrInvalidPath)
	}
	absPath, err := b.fs.Abs(path)
	if err != nil {
		return nil, err
	}
	mark, err := domain.NewMark(alias, absPath, b.clock.Now())
	if err != nil {
		return nil, err
	}
	return &mark, nil
}
