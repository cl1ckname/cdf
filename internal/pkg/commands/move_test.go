package commands_test

import (
	"testing"
	"time"

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
	d := domain.Dict{}
	d.Set(mark)
	store := new(mock.Store)
	log := new(mock.Logger)
	base := commands.NewBase(store, log)
	store.OldData = d

	dt := time.Now()
	now := makeNow(dt)

	cmd := commands.NewMove(base, now)

	err := cmd.Execute(mark.Alias)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if store.NewData == nil {
		t.Fatalf("wasn't updated")
	}
	saved, err := store.NewData.Get(mark.Alias)
	if err != nil {
		t.Fatalf("move mark not found")
	}
	if expected := mark.TimesUsed + 1; saved.TimesUsed != expected {
		t.Fatalf("used %d times insted of %d", saved.TimesUsed, expected)
	}
	if saved.LastUsed != dt {
		t.Fatalf("last used not updated")
	}
}

func makeNow(t time.Time) func() time.Time {
	return func() time.Time {
		return t
	}
}
