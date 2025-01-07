package presenters_test

import (
	"bytes"
	"testing"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/pkg/presenters"
)

func TestList(t *testing.T) {
	expected := "home\t/home/user\nsystemd\t/etc/systemd\n"
	buf := bytes.NewBufferString("")
	l := presenters.NewList(buf)
	marks := []domain.Mark{
		{
			Alias: "home",
			Path:  "/home/user",
		},
		{
			Alias: "systemd",
			Path:  "/etc/systemd",
		},
	}

	err := l.Present(marks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	res := buf.String()
	if res != expected {
		t.Fatalf("expected %s, got %s", expected, res)
	}
}
