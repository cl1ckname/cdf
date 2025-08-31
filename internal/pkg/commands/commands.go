// Package commands contains command handlers
package commands

import (
	"github.com/cl1ckname/cdf/internal/logger"
	"github.com/cl1ckname/cdf/internal/pkg/dict"
)

type Store interface {
	Load() (dict.Dict, error)
	Save(dict.Dict) error
	Cwd() (string, error)
}

type Base struct {
	log   logger.Logger
	store Store
}

func NewBase(store Store, log logger.Logger) Base {
	return Base{
		store: store,
		log:   log,
	}
}
