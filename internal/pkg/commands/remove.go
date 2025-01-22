package commands

import (
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Remover interface {
	List() ([]domain.Mark, error)
	Replace(marks []domain.Mark) error
}

type Remove struct {
	remover Store
}

func NewRemove(remover Store) Remove {
	return Remove{
		remover: remover,
	}
}

func (r Remove) Execute(alias string) error {
	marks, err := r.remover.Load()
	if err != nil {
		return err
	}
	removed := marks.Remove(alias)
	if !removed {
		return domain.ErrNotFound
	}
	return nil
}

func remove[T any](arr []T, at int) []T {
	arr[at] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}
