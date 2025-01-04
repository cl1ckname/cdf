package handler

import (
	"fmt"
)

type Store interface {
	Append(record string) error
}

type Marks struct {
	store Store
}

func NewHandler(store Store) Marks {
	return Marks{
		store: store,
	}
}

func (h Marks) Add(args Args, _ Kwargs) error {
	if len(args) < 2 {
		return fmt.Errorf("alias mark path required")
	}
	alias := args[0]
	path := args[1]

	record := formatRecord(alias, path)
	return h.store.Append(record)
}

func formatRecord(alias, path string) string {
	return alias + "=" + path
}
