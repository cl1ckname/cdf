package commands_test

import (
	"testing"
	"time"

	"github.com/cl1ckname/cdf/internal/collection/dict"
	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/test/mock"
)

func TestMove(t *testing.T) {
	t.Parallel()

	mark := domain.Mark{
		Alias:     "home",
		Path:      "/home/user",
		TimesUsed: 10,
	}
	m := new(mover)
	d := dict.Dict{}
	d.Set(mark)
	store := new(mock.Store)
	log := new(mock.Logger)
	base := commands.NewBase(store, log)
	store.OldData = d
	clock := mock.Clock{
		Time: time.Now(),
	}

	cmd := commands.NewMove(base, m, clock)
	to := "/tmp/cdf-2112"

	err := cmd.Execute(mark.Alias, to)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if writes := len(m.writes); writes != 1 {
		t.Fatalf("expected only one write, actually %d", writes)
	}
	write := m.writes[0]
	if expected := to + mark.Path; expected != write {
		t.Fatalf("expected write %s, got %s", expected, write)
	}

	if store.NewData == nil {
		t.Fatalf("wasn't updated")
	}
	saved, ok := store.NewData.Get(mark.Alias)
	if !ok {
		t.Fatalf("move mark not found")
	}
	if expected := mark.TimesUsed + 1; saved.TimesUsed != expected {
		t.Fatalf("used %d times insted of %d", saved.TimesUsed, expected)
	}
	if saved.LastUsed != clock.Time {
		t.Fatalf("last used not updated")
	}
}

type mover struct {
	writes []string
}

func (m *mover) WriteTo(file, value string) error {
	m.writes = append(m.writes, file+value)
	return nil
}
