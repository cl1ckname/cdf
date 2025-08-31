// Package dict contains Dict - collection of marks
package dict

import (
	"errors"
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

var ErrNotFound = errors.New("not found")

type Dict map[string]domain.Mark

func (d Dict) Get(alias string) (domain.Mark, error) {
	m, ok := d[alias]
	if !ok {
		return domain.Mark{}, fmt.Errorf("mark %s: %w", alias, ErrNotFound)
	}
	return m, nil
}

func (d Dict) Set(m domain.Mark) {
	d[m.Alias] = m
}

func (d Dict) Iterate() <-chan domain.Mark {
	c := make(chan domain.Mark)
	go func() {
		defer close(c)
		for _, mark := range d {
			c <- mark
		}
	}()
	return c
}

func (d Dict) Remove(alias string) error {
	if _, ok := d[alias]; !ok {
		return fmt.Errorf("mark %s: %w", alias, ErrNotFound)
	}
	delete(d, alias)
	return nil
}
