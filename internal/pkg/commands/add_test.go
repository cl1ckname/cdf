package commands_test

import (
	"errors"
	"testing"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/test/mock"
)

const (
	alias = "home"
	path  = "/home/user"
)

func TestAddSuccess(t *testing.T) {
	newMark := domain.Mark{
		Alias: alias,
		Path:  path,
	}

	store := new(mock.Store)
	log := new(mock.Logger)
	base := commands.NewBase(store, log)
	store.OldData = domain.Dict{}
	f := new(fabric)
	f.mark = &newMark

	cmd := commands.NewAdd(base, f)

	path := path
	err := cmd.Execute(alias, &path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := store.NewData.Get(alias)
	if err != nil {
		t.Fatalf("mark %s not set", alias)
	}
	if count := len(store.NewData); count != 1 {
		t.Fatalf("too much adds: %d", count)
	}

	if got != newMark {
		t.Fatalf("wrong saved mark: %v vs %v", got, newMark)
	}
}

func TestAddAlreadyExists(t *testing.T) {
	newMark := domain.Mark{
		Alias: alias,
		Path:  path,
	}

	store := new(mock.Store)
	log := new(mock.Logger)
	base := commands.NewBase(store, log)
	dt := domain.Dict{}
	dt.Set(newMark)
	store.OldData = dt

	f := new(fabric)
	f.mark = &newMark

	cmd := commands.NewAdd(base, f)

	path := path
	err := cmd.Execute(alias, &path)
	if err == nil {
		t.Fatalf("expected error: %v", err)
	}
	if !errors.Is(err, domain.ErrAlreadyExists) {
		t.Fatalf("expected ErrAlreadyExists, go %v", err)
	}
}

func TestAddCwd(t *testing.T) {
	alias, path := "home", "/home/user"
	newMark := domain.Mark{
		Alias: alias,
		Path:  path,
	}

	store := new(mock.Store)
	log := new(mock.Logger)
	base := commands.NewBase(store, log)
	store.OldData = domain.Dict{}
	store.Wd = path
	f := new(fabric)
	f.mark = &newMark

	cmd := commands.NewAdd(base, f)

	err := cmd.Execute(alias, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := store.NewData.Get(alias)
	if err != nil {
		t.Fatalf("mark %s not set", alias)
	}
	if count := len(store.NewData); count != 1 {
		t.Fatalf("too much adds: %d", count)
	}

	if got != newMark {
		t.Fatalf("wrong saved mark: %v vs %v", got, newMark)
	}
}

type fabric struct {
	mark *domain.Mark
}

func (f fabric) Build(_, _ string) (*domain.Mark, error) {
	return f.mark, nil
}
