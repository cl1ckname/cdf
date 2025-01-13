package presenters

import (
	"encoding/json"
	"io"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type JSON struct {
	out io.Writer
}

func NewJSON(out io.Writer) JSON {
	return JSON{out: out}
}

func (j JSON) Present(marks []domain.Mark) error {
	markMap := make(map[string]string, len(marks))
	for _, mark := range marks {
		markMap[mark.Alias] = mark.Path
	}

	encoder := json.NewEncoder(j.out)
	encoder.SetIndent("", "  ")
	return encoder.Encode(markMap)
}
