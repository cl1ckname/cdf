package commands_test

import (
	"errors"
	"testing"

	"github.com/cl1ckname/cdf/internal/collection/dict"
	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/test/mock"
)

func TestAddSuccess(t *testing.T) {
	alias, path := "home", "/home/user"
	newMark := domain.Mark{
		Alias: alias,
		Path:  path,
	}

	store := new(mock.Store)
	store.OldData = dict.Dict{}
	f := new(fabric)
	f.mark = &newMark

	cmd := commands.NewAdd(store, f)

	err := cmd.Execute(alias, path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, ok := store.NewData.Get(alias)
	if !ok {
		t.Fatalf("mark %s not set", alias)
	}
	var counter int
	for range store.NewData.Iterate() {
		counter++
	}
	if counter != 1 {
		t.Fatalf("too much adds: %d", counter)
	}

	if got != newMark {
		t.Fatalf("wrong saved mark: %v vs %v", got, newMark)
	}
}

func TestAddAlreadyExists(t *testing.T) {
	alias, path := "home", "/home/user"
	newMark := domain.Mark{
		Alias: alias,
		Path:  path,
	}

	st := new(mock.Store)
	dt := dict.Dict{}
	dt.Set(newMark)
	st.OldData = dt

	f := new(fabric)
	f.mark = &newMark

	cmd := commands.NewAdd(st, f)

	err := cmd.Execute(alias, path)
	if err == nil {
		t.Fatalf("expected error: %v", err)
	}
	if !errors.Is(err, domain.ErrAlreadyExists) {
		t.Fatalf("expected ErrAlreadyExists, go %v", err)
	}
}

type fabric struct {
	mark *domain.Mark
}

func (f fabric) Build(_, _ string) (*domain.Mark, error) {
	return f.mark, nil
}
