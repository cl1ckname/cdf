package commands

import (
	"errors"
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/dict"
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

func (c Add) Execute(alias string, path *string) error {
	markPath, err := c.store.Cwd()
	if err != nil {
		return fmt.Errorf("get cwd: %w", err)
	}

	if path != nil {
		markPath = *path
	}
	mark, err := c.builder.Build(alias, markPath)
	if err != nil {
		return err
	}

	col, err := c.store.Load()
	if err != nil {
		return err
	}

	rec, err := col.Get(mark.Alias)
	if err == nil {
		return fmt.Errorf("this alias already in use (%s): %w", rec.Alias, domain.ErrAlreadyExists)
	}
	if !errors.Is(err, dict.ErrNotFound) {
		return err
	}
	col.Set(*mark)

	if err := c.store.Save(col); err != nil {
		return err
	}
	c.log.Info("new mark added:", mark.Path, "as", mark.Alias)
	return nil
}
