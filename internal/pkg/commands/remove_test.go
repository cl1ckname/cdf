package commands_test

import (
	"errors"
	"maps"
	"testing"

	"github.com/cl1ckname/cdf/internal/collection/dict"
	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/test/mock"
)

var stub = struct{}{}

func TestRemove(t *testing.T) {
	tests := []struct {
		name string
		arr  map[string]domain.Mark
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
			st := new(mock.Store)
			log := new(mock.Logger)
			base := commands.NewBase(st, log)
			st.OldData = dict.Dict(test.arr)
			cmd := commands.NewRemove(base)
			err := cmd.Execute(test.arg)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if st.NewData == nil {
				t.Fatalf("no data saved")
			}
			rest := map[string]struct{}{}
			for mark := range st.NewData.Iterate() {
				rest[mark.Alias] = stub
			}
			if !maps.Equal(test.rest, rest) {
				t.Fatalf("wrong key set, expected %v, got %v\n", test.rest, rest)
			}
		})
	}
}

func TestNotFount(t *testing.T) {
	st := new(mock.Store)
	log := new(mock.Logger)
	base := commands.NewBase(st, log)
	st.OldData = dict.Dict{}
	cmd := commands.NewRemove(base)
	err := cmd.Execute("b")
	if err == nil {
		t.Fatalf("expected error, got nil\n")
	}
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected %v, got %v", domain.ErrNotFound, err)
	}
}

func genset(str ...string) map[string]domain.Mark {
	mrks := make(map[string]domain.Mark, len(str))
	for _, alias := range str {
		mrks[alias] = domain.Mark{Alias: alias}
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
