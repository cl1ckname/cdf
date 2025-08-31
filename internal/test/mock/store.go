package mock

import "github.com/cl1ckname/cdf/internal/pkg/dict"

type Store struct {
	OldData dict.Dict
	NewData dict.Dict
}

func (s *Store) Load() (dict.Dict, error) {
	return s.OldData, nil
}

func (s *Store) Save(c dict.Dict) error {
	s.NewData = c
	return nil
}
