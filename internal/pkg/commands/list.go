package commands

import "github.com/cl1ckname/cdf/internal/pkg/domain"

type Lister interface {
	List() ([]domain.Mark, error)
}

type Presenter interface {
	Present(marks []domain.Mark) error
}

type List struct {
	lister    Lister
	presenter Presenter
}

func NewList(l Lister, p Presenter) List {
	return List{
		lister:    l,
		presenter: p,
	}
}

func (l List) Execute() error {
	marks, err := l.lister.List()
	if err != nil {
		return err
	}
	return l.presenter.Present(marks)
}
