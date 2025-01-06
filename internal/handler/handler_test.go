package handler_test

import (
	"testing"

	"github.com/cl1ckname/cdf/internal/handler"
)

func TestAdd(t *testing.T) {
	s := new(store)
	h := handler.NewHandler(s)

	path := "/home"
	alias := "h"

	expectedRecord := alias + handler.RecordSeparator + path

	call := handler.Call{
		Code: handler.CodeAdd,
		Args: handler.Args{alias, path},
	}
	err := h.Permorm(call)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if appends := len(s.appends); appends != 1 {
		t.Fatalf("unexpected number of appends: %d", appends)
	}
	if record := s.appends[0]; record != expectedRecord {
		t.Fatalf("wrong record, expected %s, got %s", expectedRecord, record)
	}
}

type store struct {
	appends []string
}

func (s *store) Append(record string) error {
	s.appends = append(s.appends, record)
	return nil
}
