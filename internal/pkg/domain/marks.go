package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const AliasRE = `^[\w-]+`

var aliasRe = regexp.MustCompile(AliasRE)

// RecordSeparator is a sybol that separates alias and corresponding cd path
const RecordSeparator = "="

var ErrInvalidAlias = errors.New("invalid path alias")
var ErrInvalidPath = errors.New("invalid target path")

type Mark struct {
	Alias string
	Path  string
}

func NewMark(alias, path string) (m Mark, err error) {
	if !aliasRe.MatchString(alias) {
		return m, fmt.Errorf("alias should contain only a-z A-Z 0-9 _ - symbols: %w", ErrInvalidAlias)
	}
	m.Alias = alias
	m.Path = path
	return
}

func (m Mark) String() string {
	return m.Alias + RecordSeparator + m.Path
}

func ParseMark(record string) (Mark, error) {
	parts := strings.Split(record, RecordSeparator)
	if len(parts) != 2 {
		return Mark{}, fmt.Errorf("invalid record: %s", record)
	}
	alias := parts[0]
	path := parts[1]
	return NewMark(alias, path)
}
