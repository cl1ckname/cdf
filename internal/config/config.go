// Package config contains struct for persistent state data storing
package config

import "github.com/cl1ckname/cdf/internal/pkg/domain"

type Config struct {
	Marks map[string]domain.Mark `json:"marks"`
}
