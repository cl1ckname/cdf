package dict

import "github.com/cl1ckname/cdf/internal/pkg/domain"

type Dict map[string]domain.Mark

func (d Dict) Get(alias string) (domain.Mark, bool) {
	m, ok := d[alias]
	return m, ok
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

func (d Dict) Remove(alias string) bool {
	if _, ok := d[alias]; !ok {
		return false
	}
	delete(d, alias)
	return true
}
