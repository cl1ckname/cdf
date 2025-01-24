package mock

import "github.com/cl1ckname/cdf/internal/pkg/domain"

type Store struct {
	OldData domain.Collection
	NewData domain.Collection
}

func (s *Store) Load() (domain.Collection, error) {
	return s.OldData, nil
}

func (s *Store) Save(c domain.Collection) error {
	s.NewData = c
	return nil
}
