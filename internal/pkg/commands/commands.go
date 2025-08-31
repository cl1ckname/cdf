package commands

import (
	"github.com/cl1ckname/cdf/internal/collection/dict"
)

type Store interface {
	Load() (dict.Dict, error)
	Save(dict.Dict) error
}
