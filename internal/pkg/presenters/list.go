package presenters

import (
	"io"
	"strconv"
	"strings"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

const layout = "15:04:05 2006-01-02"

type List struct {
	out  io.Writer
	opts Opts
}

func NewList(out io.Writer, opts Opts) List {
	return List{
		out:  out,
		opts: opts,
	}
}

func (l List) Present(marks []domain.Mark) error {
	_, long := l.opts["l"]
	err := l.printHeader(long)
	if err != nil {
		return err
	}
	for _, mark := range marks {
		if err := l.print(mark, long); err != nil {
			return err
		}
	}
	return nil
}

func (l List) printHeader(long bool) error {
	if !long {
		return nil
	}
	headerLine := header()
	_, err := l.out.Write(append([]byte(headerLine), '\n'))
	return err
}

func (l List) print(m domain.Mark, long bool) error {
	line := formatMark(m, long)
	_, err := l.out.Write(append([]byte(line), '\n'))
	return err
}

func formatMark(m domain.Mark, long bool) string {
	fields := fieldSet(m, long)
	return strings.Join(fields, "\t")
}

func header() string {
	columns := []string{"alias", "path", "used", "last usage", "created"}
	return strings.Join(columns, "\t")
}

func fieldSet(m domain.Mark, long bool) []string {
	elems := []string{m.Alias, m.Path}
	if !long {
		return elems
	}
	count := strconv.Itoa(m.TimesUsed)
	last := m.LastUsed.Format(layout)
	created := m.Created.Format(layout)
	return append(elems, count, last, created)
}
