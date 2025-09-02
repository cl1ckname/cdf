package mock

import (
	"github.com/cl1ckname/cdf/internal/config"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Store struct {
	OldData domain.Dict
	NewData domain.Dict
	Wd      string
}

func (s *Store) Load() (*config.Config, error) {
	return &config.Config{
		Marks: s.OldData,
	}, nil
}

func (s *Store) Save(c config.Config) error {
	s.NewData = c.Marks
	return nil
}

func (s Store) Cwd() (string, error) {
	return s.Wd, nil
}
