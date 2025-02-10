package store

import (
	"encoding/json"
	"io"
	"time"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Record struct {
	Alias     string `json:"alias"`
	Path      string `json:"path"`
	TimesUsed int    `json:"timesUsed"`
	LastUsed  int64  `json:"lastUsed,omitempty"`
	Created   int64  `json:"created"`
}

func (r Record) Write(w io.Writer) error {
	b, _ := json.Marshal(r)
	_, err := w.Write(b)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\n"))
	return err
}

func NewRecord(m domain.Mark) Record {
	return Record{
		Alias:     m.Alias,
		Path:      m.Path,
		TimesUsed: m.TimesUsed,
		LastUsed:  m.LastUsed.Unix(),
		Created:   m.Created.Unix(),
	}
}

func NewMark(r Record) domain.Mark {
	return domain.Mark{
		Alias:     r.Alias,
		Path:      r.Path,
		TimesUsed: r.TimesUsed,
		Created:   time.Unix(r.Created, 0),
		LastUsed:  time.Unix(r.LastUsed, 0),
	}
}

func ParseRecord(row []byte) (r Record, err error) {
	err = json.Unmarshal(row, &r)
	return
}
