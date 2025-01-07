package commands_test

import (
	"testing"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

func TestMove(t *testing.T) {
	t.Parallel()

	m := new(mover)
	cmd := commands.NewMove(m)
	mark := domain.Mark{
		Alias: "home",
		Path:  "/home/user",
	}
	m.mark = mark
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
	mark   domain.Mark
	writes []string
}

func (m *mover) Find(_ string) (domain.Mark, error) {
	return m.mark, nil
}

func (m *mover) WriteTo(file, value string) error {
	m.writes = append(m.writes, file+value)
	return nil
}
