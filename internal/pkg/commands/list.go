package commands

import (
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/pkg/presenters"
)

type Presenter interface {
	Present(marks []domain.Mark) error
}

type PresenterFabric interface {
	Build(format domain.Format, opts presenters.Opts) Presenter
}

type List struct {
	Base
	presenter PresenterFabric
}

func NewList(l Base, f PresenterFabric) List {
	return List{
		Base:      l,
		presenter: f,
	}
}

func (l List) Execute(format domain.Format, opts presenters.Opts) error {
	coll, err := l.store.Load()
	if err != nil {
		return err
	}
	marks := coll.Slice()
	presenter := l.presenter.Build(format, opts)
	return presenter.Present(marks)
}
