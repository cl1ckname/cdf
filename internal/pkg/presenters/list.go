package presenters

import (
	"io"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type List struct {
	out io.Writer
}

func NewList(out io.Writer) List {
	return List{
		out: out,
	}
}

func (l List) Present(marks []domain.Mark) error {
	for _, mark := range marks {
		if err := l.print(mark); err != nil {
			return err
		}
	}
	return nil
}

func (l List) print(m domain.Mark) error {
	line := formatMark(m)
	_, err := l.out.Write([]byte(line))
	return err
}

func formatMark(m domain.Mark) string {
	return m.Alias + "\t" + m.Path + "\n"
}
