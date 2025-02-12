package commands

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Remover interface {
	List() ([]domain.Mark, error)
	Replace(marks []domain.Mark) error
}

type Remove struct {
	Base
}

func NewRemove(base Base) Remove {
	return Remove{
		Base: base,
	}
}

func (r Remove) Execute(alias string) error {
	marks, err := r.store.Load()
	if err != nil {
		return err
	}
	removed := marks.Remove(alias)
	if !removed {
		return fmt.Errorf("mark %s %w", alias, domain.ErrNotFound)
	}
	if err := r.store.Save(marks); err != nil {
		return err
	}
	r.log.Info("mark", alias, "removed")
	return nil
}
