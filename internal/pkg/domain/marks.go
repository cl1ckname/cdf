package domain

import (
	"encoding/json"
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
	Alias   string    `json:"alias"`
	Path    string    `json:"path"`
	Created time.Time `json:"created"`
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

func (m Mark) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}

func ParseMark(record string) (m Mark, err error) {
	err = json.Unmarshal([]byte(record), &m)
	return
}
