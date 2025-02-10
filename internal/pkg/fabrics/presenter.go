package fabrics

import (
	"os"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/pkg/presenters"
)

type Presenter struct{}

func (p Presenter) Build(format domain.Format, opts presenters.Opts) commands.Presenter {
	switch format {
	case domain.JSONFormat:
		return presenters.NewJSON(os.Stdout, opts)
	case domain.AliasFormat:
		return presenters.NewAlias(os.Stdout, opts)
	default:
		return presenters.NewList(os.Stdout, opts)
	}
}

var PresenterInstance = Presenter{}
