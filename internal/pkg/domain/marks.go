package domain

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

const AliasRE = `^[\w-]+`

var aliasRe = regexp.MustCompile(AliasRE)

var ErrInvalidAlias = errors.New("invalid path alias")
var ErrInvalidPath = errors.New("invalid target path")
var ErrAlreadyExists = errors.New("bookmark with this alias already exist")
var ErrNotFound = errors.New("not found")

type Mark struct {
	Alias     string
	Path      string
	TimesUsed int
	Created   time.Time
	LastUsed  time.Time
}

func NewMark(alias, path string, now time.Time) (m Mark, err error) {
	if !aliasRe.MatchString(alias) {
		return m, fmt.Errorf("alias should contain only a-z A-Z 0-9 _ - symbols: %w", ErrInvalidAlias)
	}
	m.Alias = alias
	m.Path = path
	m.Created = now
	return
}

func (m *Mark) Use(now time.Time) {
	m.LastUsed = now
	m.TimesUsed++
}

type Collection interface {
	Get(alias string) (Mark, bool)
	Set(m Mark)
	Iterate() <-chan Mark
	Remove(alias string) bool
}
