package commands

import (
	"time"
)

type Move struct {
	Base
	now func() time.Time
}

func NewMove(base Base, now func() time.Time) Move {
	return Move{
		Base: base,
		now:  now,
	}
}

func (c Move) Execute(alias string) error {
	coll, err := c.Load()
	if err != nil {
		return err
	}
	mark, err := coll.Get(alias)
	if err != nil {
		return err
	}
	mark.Use(c.now())
	coll.Set(mark)
	if err := c.Save(coll); err != nil {
		return err
	}

	c.log.Print(mark.Path)
	return nil
}
