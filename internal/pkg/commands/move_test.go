package commands_test

import (
	"testing"

	"github.com/cl1ckname/cdf/internal/collection/dict"
	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/test/mock"
)

func TestMove(t *testing.T) {
	t.Parallel()

	mark := domain.Mark{
		Alias: "home",
		Path:  "/home/user",
	}
	m := new(mover)
	d := dict.Dict{}
	d.Set(mark)
	store := new(mock.Store)
	store.OldData = d

	cmd := commands.NewMove(store, m)
	to := "/tmp/cdf-2112"

	got, err := cmd.Execute(mark.Alias, to)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected := mark.Path; got != expected {
		t.Fatalf("expected %s, got %s", expected, got)
	}
	if writes := len(m.writes); writes != 1 {
		t.Fatalf("expected only one write, actually %d", writes)
	}
	write := m.writes[0]
	if expected := to + mark.Path; expected != write {
		t.Fatalf("expected write %s, got %s", expected, write)
	}
}

type mover struct {
	writes []string
}

func (m *mover) WriteTo(file, value string) error {
	m.writes = append(m.writes, file+value)
	return nil
}
