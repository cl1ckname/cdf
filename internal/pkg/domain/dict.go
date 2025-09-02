package domain

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type Dict map[string]Mark

func (d Dict) Get(alias string) (Mark, error) {
	m, ok := d[alias]
	if !ok {
		return Mark{}, fmt.Errorf("mark %s: %w", alias, ErrNotFound)
	}
	return m, nil
}

func (d Dict) Set(m Mark) {
	d[m.Alias] = m
}

func (d Dict) Remove(alias string) error {
	if _, ok := d[alias]; !ok {
		return fmt.Errorf("mark %s: %w", alias, ErrNotFound)
	}
	delete(d, alias)
	return nil
}

func (d Dict) Slice() []Mark {
	marks := make([]Mark, 0, len(d))
	for _, mark := range d {
		marks = append(marks, mark)
	}
	return marks
}
