package commands

import (
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
	marks, err := r.Load()
	if err != nil {
		return err
	}
	if err = marks.Remove(alias); err != nil {
		return err
	}
	if err := r.Save(marks); err != nil {
		return err
	}
	r.log.Info("mark", alias, "removed")
	return nil
}
