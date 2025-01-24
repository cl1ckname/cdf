package mock

import "github.com/cl1ckname/cdf/internal/pkg/domain"

type Store struct {
	oldData domain.Collection
	newData domain.Collection
}

func (s *Store) Load() (domain.Collection, error) {
	return s.oldData, nil
}

func (s *Store) Save(c domain.Collection) error {
	s.newData = c
	return nil
}
