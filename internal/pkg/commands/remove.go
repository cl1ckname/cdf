package commands

import (
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Remover interface {
	List() ([]domain.Mark, error)
	Replace(marks []domain.Mark) error
}

type Remove struct {
	remover Remover
}

func NewRemove(remover Remover) Remove {
	return Remove{
		remover: remover,
	}
}

func (r Remove) Execute(alias string) error {
	marks, err := r.remover.List()
	if err != nil {
		return err
	}

	var removed bool
	for i, mark := range marks {
		if mark.Alias == alias {
			removed = true
			marks = remove(marks, i)
			break
		}
	}
	if !removed {
		return domain.ErrNotFound
	}

	return r.remover.Replace(marks)
}

func remove[T any](arr []T, at int) []T {
	arr[at] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}
