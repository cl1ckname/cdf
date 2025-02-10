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
	store     Store
	presenter PresenterFabric
}

func NewList(l Store, f PresenterFabric) List {
	return List{
		store:     l,
		presenter: f,
	}
}

func (l List) Execute(format domain.Format, opts presenters.Opts) error {
	coll, err := l.store.Load()
	if err != nil {
		return err
	}
	var marks []domain.Mark
	for mark := range coll.Iterate() {
		marks = append(marks, mark)
	}
	presenter := l.presenter.Build(format, opts)
	return presenter.Present(marks)
}
