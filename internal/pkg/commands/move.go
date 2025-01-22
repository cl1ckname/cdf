package commands

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Mover interface {
	WriteTo(file, value string) error
}

type Move struct {
	store Store
	mover Mover
}

func NewMove(store Store, mover Mover) Move {
	return Move{
		store: store,
		mover: mover,
	}
}

func (c Move) Execute(alias string, resTo string) (string, error) {
	coll, err := c.store.Load()
	if err != nil {
		return "", err
	}
	mark, ok := coll.Get(alias)
	if ok {
		return "", fmt.Errorf("mark %s: %w", alias, domain.ErrNotFound)
	}
	return mark.Path, c.mover.WriteTo(resTo, mark.Path)
}
