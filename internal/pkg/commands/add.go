package commands

import (
	"errors"
	"fmt"
	"io/fs"
	"regexp"
)

type Appender interface {
	Append(record string) error
	Find(alias string) (string, error)
	Stat(path string) (fs.FileInfo, error)
	Abs(path string) (string, error)
}

// RecordSeparator is a sybol that separates alias and corresponding cd path
const RecordSeparator = "="

const AliasRE = `^[\w-]+`

var aliasRe = regexp.MustCompile(AliasRE)

var ErrInvalidAlias = errors.New("invalid path alias")
var ErrInvalidPath = errors.New("invalid target path")
var ErrAlreadyExists = errors.New("bookmark with this alias already exist")
var ErrNotFound = errors.New("not found")

type Add struct {
	appender Appender
}

func NewAdd(a Appender) Add {
	return Add{
		appender: a,
	}
}

func (c Add) Execute(alias, path string) error {
	if !aliasRe.MatchString(alias) {
		return fmt.Errorf("alias should contain only a-z A-Z 0-9 _ - symbols: %w", ErrInvalidAlias)
	}
	if rec, err := c.appender.Find(alias); err == nil {
		return fmt.Errorf("this alias already in use (%s): %w", rec, ErrNotFound)
	}

	info, err := c.appender.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("path sould be and folder, not file: %w", ErrInvalidPath)
	}
	absPath, err := c.appender.Abs(path)
	if err != nil {
		return err
	}
	record := formatRecord(alias, absPath)
	return c.appender.Append(record)
}

func formatRecord(alias, path string) string {
	return alias + RecordSeparator + path
}
