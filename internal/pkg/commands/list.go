package commands

import "github.com/cl1ckname/cdf/internal/pkg/domain"

type Lister interface {
	List() ([]domain.Mark, error)
}

type Presenter interface {
	Present(marks []domain.Mark) error
}

type PresenterFabric interface {
	Build(format domain.Format) Presenter
}

type List struct {
	lister    Lister
	presenter PresenterFabric
}

func NewList(l Lister, f PresenterFabric) List {
	return List{
		lister:    l,
		presenter: f,
	}
}

func (l List) Execute(format domain.Format) error {
	marks, err := l.lister.List()
	if err != nil {
		return err
	}
	presenter := l.presenter.Build(format)
	return presenter.Present(marks)
}
