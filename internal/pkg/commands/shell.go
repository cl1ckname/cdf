package commands

import (
	"fmt"
	"io"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Wraps = map[domain.Shell][]byte

type Shell struct {
	out   io.Writer
	wraps Wraps
}

func NewShell(out io.Writer, wraps Wraps) Shell {
	return Shell{
		out:   out,
		wraps: wraps,
	}
}

func (c Shell) Execute(shell domain.Shell) error {
	wrap, ok := c.wraps[shell]
	if !ok {
		return fmt.Errorf("unsupported shell %s: %w", shell, domain.ErrUnsupportedShell)
	}
	if _, err := c.out.Write(wrap); err != nil {
		return err
	}
	return nil
}
