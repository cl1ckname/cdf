package commands_test

import (
	"errors"
	"maps"
	"testing"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

var stub = struct{}{}

func TestRemove(t *testing.T) {
	tests := []struct {
		name string
		arr  []domain.Mark
		rest map[string]struct{}
		arg  string
	}{
		{
			name: "only mark",
			arr:  genset("a"),
			rest: genrest(),
			arg:  "a",
		},
		{
			name: "first mark",
			arr:  genset("a", "b", "c"),
			rest: genrest("b", "c"),
			arg:  "a",
		},
		{
			name: "middle mark",
			arr:  genset("a", "b", "c"),
			rest: genrest("a", "c"),
			arg:  "b",
		},
		{
			name: "last mark",
			arr:  genset("a", "b", "c"),
			rest: genrest("a", "b"),
			arg:  "c",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := new(remover)
			r.marks = test.arr
			cmd := commands.NewRemove(r)
			err := cmd.Execute(test.arg)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}
			rest := map[string]struct{}{}
			for _, mark := range r.res {
				rest[mark.Alias] = stub
			}
			if !maps.Equal(test.rest, rest) {
				t.Fatalf("wrong key set, expected %v, got %v\n", test.rest, rest)
			}
		})
	}
}

func TestNotFount(t *testing.T) {
	r := new(remover)
	cmd := commands.NewRemove(r)
	err := cmd.Execute("b")
	if err == nil {
		t.Fatalf("expected error, got nil\n")
	}
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected %v, got %v", domain.ErrNotFound, err)
	}
}

func genset(str ...string) []domain.Mark {
	mrks := make([]domain.Mark, len(str))
	for i, alias := range str {
		mrks[i] = domain.Mark{Alias: alias}
	}
	return mrks
}

func genrest(str ...string) map[string]struct{} {
	mrks := make(map[string]struct{}, len(str))
	for _, alias := range str {
		mrks[alias] = stub
	}
	return mrks
}

type remover struct {
	marks []domain.Mark
	res   []domain.Mark
}

func (r *remover) List() ([]domain.Mark, error) {
	return r.marks, nil
}

func (r *remover) Replace(marks []domain.Mark) error {
	r.res = marks
	return nil
}
