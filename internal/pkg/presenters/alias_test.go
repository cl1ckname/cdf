package presenters_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/pkg/presenters"
)

func TestAlias(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	opts := presenters.Opts{}
	l := presenters.NewAlias(buf, opts)
	marks := []domain.Mark{
		{Alias: "home"},
		{Alias: "user"},
		{Alias: "projects"},
	}
	err := l.Present(marks)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	aliases := map[string]bool{}
	for _, mark := range marks {
		aliases[mark.Alias] = false
	}

	wrote := buf.String()
	wrote = strings.Trim(wrote, "\n")
	wroteAliases := strings.Split(wrote, " ")
	for _, alias := range wroteAliases {
		aliases[alias] = true
	}
	for alias, v := range aliases {
		if !v {
			t.Fatalf("alias %s not met", alias)
		}
	}
}
