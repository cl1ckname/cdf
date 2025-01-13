package domain

import "errors"

type Format string

const (
	DefaultFormat Format = "default"
	JSONFormat    Format = "json"
)

var ErrUnknownFormat = errors.New("unknown format")

func ParseFormat(s *string) (Format, bool) {
	if s == nil {
		return DefaultFormat, true
	}
	formatMap := map[string]Format{
		string(JSONFormat):    JSONFormat,
		string(DefaultFormat): DefaultFormat,
		"":                    DefaultFormat,
	}
	format, ok := formatMap[*s]
	return format, ok
}
