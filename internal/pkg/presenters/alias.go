package presenters

import (
	"io"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Alias struct {
	out io.Writer
}

func NewAlias(out io.Writer) Alias {
	return Alias{out: out}
}

func (a Alias) Present(marks []domain.Mark) error {
	for _, mark := range marks {
		b := []byte(mark.Alias)
		b = append(b, ' ')
		_, err := a.out.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
