package commands

import (
	"errors"
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Appender interface {
	Append(record domain.Mark) error
	Find(alias string) (domain.Mark, error)
}

type MarkFabric interface {
	Build(alias, path string) (*domain.Mark, error)
}

var ErrAlreadyExists = errors.New("bookmark with this alias already exist")
var ErrNotFound = errors.New("not found")

type Add struct {
	appender Appender
	builder  MarkFabric
}

func NewAdd(a Appender, f MarkFabric) Add {
	return Add{
		appender: a,
		builder:  f,
	}
}

func (c Add) Execute(alias, path string) error {
	mark, err := c.builder.Build(alias, path)
	if err != nil {
		return err
	}

	rec, err := c.appender.Find(mark.Alias)
	if err == nil {
		return fmt.Errorf("this alias already in use (%s): %w", rec, ErrAlreadyExists)
	}
	if !errors.Is(err, ErrNotFound) {
		return fmt.Errorf("find error: %w", err)
	}

	return c.appender.Append(*mark)
}
