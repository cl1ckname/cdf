package commands

import "github.com/cl1ckname/cdf/internal/logger"

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
