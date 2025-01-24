package commands

import "github.com/cl1ckname/cdf/internal/pkg/domain"

type Store interface {
	Load() (domain.Collection, error)
	Save(domain.Collection) error
}
