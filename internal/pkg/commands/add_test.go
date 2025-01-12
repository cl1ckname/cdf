package commands_test

import (
	"errors"
	"testing"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

func TestAddSuccess(t *testing.T) {
	alias, path := "home", "/home/user"
	newMark := domain.Mark{
		Alias: alias,
		Path:  path,
	}

	ap := new(appender)
	ap.findErr = domain.ErrNotFound

	f := new(fabric)
	f.mark = &newMark

	cmd := commands.NewAdd(ap, f)

	err := cmd.Execute(alias, path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if aps := len(ap.appends); aps != 1 {
		t.Fatalf("should be only one append, got %d", aps)
	}
	got := ap.appends[0]
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

	ap := new(appender)
	ap.find = newMark

	f := new(fabric)
	f.mark = &newMark

	cmd := commands.NewAdd(ap, f)

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

type appender struct {
	appends []domain.Mark
	find    domain.Mark
	findErr error
}

func (a *appender) Append(rec domain.Mark) error {
	a.appends = append(a.appends, rec)
	return nil
}

func (a *appender) Find(alias string) (domain.Mark, error) {
	return a.find, a.findErr
}
