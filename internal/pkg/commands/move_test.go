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
	d := dict.Dict{}
	d.Set(mark)
	store := new(mock.Store)
	log := new(mock.Logger)
	base := commands.NewBase(store, log)
	store.OldData = d
	clock := mock.Clock{
		Time: time.Now(),
	}

	cmd := commands.NewMove(base, clock)

	err := cmd.Execute(mark.Alias)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
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
