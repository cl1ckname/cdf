package commands

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type MarkFabric interface {
	Build(alias, path string) (*domain.Mark, error)
}

type Add struct {
	Base
	builder MarkFabric
}

func NewAdd(a Base, f MarkFabric) Add {
	return Add{
		Base:    a,
		builder: f,
	}
}

func (c Add) Execute(alias, path string) error {
	mark, err := c.builder.Build(alias, path)
	if err != nil {
		return err
	}

	col, err := c.store.Load()
	if err != nil {
		return err
	}

	rec, exists := col.Get(mark.Alias)
	if exists {
		return fmt.Errorf("this alias already in use (%s): %w", rec.Alias, domain.ErrAlreadyExists)
	}
	col.Set(*mark)

	if err := c.store.Save(col); err != nil {
		return err
	}
	c.log.Info("new mark added:", mark.Path, "as", mark.Alias)
	return nil
}
