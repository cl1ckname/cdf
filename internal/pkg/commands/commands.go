// Package commands contains command handlers
package commands

import (
	"github.com/cl1ckname/cdf/internal/config"
	"github.com/cl1ckname/cdf/internal/logger"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Store interface {
	Load() (*config.Config, error)
	Save(config.Config) error
	Cwd() (string, error)
}

type Base struct {
	log   logger.Logger
	store Store
}

func (b Base) Load() (domain.Dict, error) {
	cfg, err := b.store.Load()
	if err != nil {
		return nil, err
	}
	return cfg.Marks, nil
}

func (b Base) Save(d domain.Dict) error {
	cfg, err := b.store.Load()
	if err != nil {
		return err
	}
	cfg.Marks = d
	return b.store.Save(*cfg)
}

func NewBase(store Store, log logger.Logger) Base {
	return Base{
		store: store,
		log:   log,
	}
}
