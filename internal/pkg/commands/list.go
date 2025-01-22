package commands

import "github.com/cl1ckname/cdf/internal/pkg/domain"

type Presenter interface {
	Present(marks []domain.Mark) error
}

type PresenterFabric interface {
	Build(format domain.Format) Presenter
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

func (l List) Execute(format domain.Format) error {
	coll, err := l.store.Load()
	if err != nil {
		return err
	}
	var marks []domain.Mark
	for mark := range coll.Iterate() {
		marks = append(marks, mark)
	}
	presenter := l.presenter.Build(format)
	return presenter.Present(marks)
}
