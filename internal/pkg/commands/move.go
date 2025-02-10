package commands

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/clock"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Mover interface {
	WriteTo(file, value string) error
}

type Move struct {
	store Store
	mover Mover
	clock clock.Clock
}

func NewMove(store Store, mover Mover, cl clock.Clock) Move {
	return Move{
		store: store,
		mover: mover,
		clock: cl,
	}
}

func (c Move) Execute(alias string, resTo string) (string, error) {
	coll, err := c.store.Load()
	if err != nil {
		return "", err
	}
	mark, ok := coll.Get(alias)
	if !ok {
		return "", fmt.Errorf("mark %s: %w", alias, domain.ErrNotFound)
	}
	mark.Use(c.clock.Now())
	coll.Set(mark)
	if err := c.store.Save(coll); err != nil {
		return "", err
	}

	return mark.Path, c.mover.WriteTo(resTo, mark.Path)
}
