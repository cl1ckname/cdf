package commands

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/clock"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Move struct {
	Base
	clock clock.Clock
}

func NewMove(base Base, cl clock.Clock) Move {
	return Move{
		Base:  base,
		clock: cl,
	}
}

func (c Move) Execute(alias string) error {
	coll, err := c.store.Load()
	if err != nil {
		return err
	}
	mark, ok := coll.Get(alias)
	if !ok {
		return fmt.Errorf("mark %s: %w", alias, domain.ErrNotFound)
	}
	mark.Use(c.clock.Now())
	coll.Set(mark)
	if err := c.store.Save(coll); err != nil {
		return err
	}

	fmt.Printf(mark.Path)
	return nil
}
