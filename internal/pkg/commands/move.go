package commands

import "github.com/cl1ckname/cdf/internal/pkg/domain"

type Mover interface {
	Find(alias string) (domain.Mark, error)
	WriteTo(file, value string) error
}

type Move struct {
	mover Mover
}

func NewMove(mover Mover) Move {
	return Move{
		mover: mover,
	}
}

func (c Move) Execute(alias string, resTo string) (string, error) {
	mark, err := c.mover.Find(alias)
	if err != nil {
		return "", err
	}
	return mark.Path, c.mover.WriteTo(resTo, mark.Path)
}
