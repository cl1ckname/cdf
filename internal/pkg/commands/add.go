package commands

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type MarkFabric interface {
	Build(alias, path string) (*domain.Mark, error)
}

type Add struct {
	appender Store
	builder  MarkFabric
}

func NewAdd(a Store, f MarkFabric) Add {
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

	col, err := c.appender.Load()
	if err != nil {
		return err
	}

	rec, exists := col.Get(mark.Alias)
	if exists {
		return fmt.Errorf("this alias already in use (%s): %w", rec.Alias, domain.ErrAlreadyExists)
	}
	col.Set(*mark)

	return c.appender.Save(col)
}
