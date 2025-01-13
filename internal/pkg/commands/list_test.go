package commands_test

import (
	"testing"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/utils"
)

func TestList(t *testing.T) {
	p := new(presenter)
	fab := new(presenterFabric)
	fab.presenter = p
	l := new(lister)
	cmd := commands.NewList(l, fab)
	marks := []domain.Mark{
		{Alias: "h", Path: "/home/username"},
	}
	l.marks = marks

	f := domain.JSONFormat
	err := cmd.Execute(f)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if fab.format != f {
		t.Fatalf("wrong format passed: %s", f)
	}
	if !utils.ArrayEq(marks, p.marks) {
		t.Fatalf("presenter expected %v, got %v", marks, p.marks)
	}

}

type presenter struct {
	marks []domain.Mark
}

func (p *presenter) Present(marks []domain.Mark) error {
	p.marks = marks
	return nil
}

type presenterFabric struct {
	format    domain.Format
	presenter *presenter
}

func (p *presenterFabric) Build(f domain.Format) commands.Presenter {
	p.format = f
	return p.presenter
}

type lister struct {
	marks []domain.Mark
}

func (l lister) List() ([]domain.Mark, error) {
	return l.marks, nil
}
