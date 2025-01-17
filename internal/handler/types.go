package handler

import "github.com/cl1ckname/cdf/internal/pkg/domain"

const RecordSeparator = "="

type Kwargs = map[string]string
type Args = []string
type Handler = func(args Args, kwargs Kwargs) error

type Call struct {
	Kwargs Kwargs
	Args   Args
	Code   *domain.Command
}
