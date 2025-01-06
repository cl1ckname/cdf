package app

import (
	"github.com/cl1ckname/cdf/internal/cli"
	"github.com/cl1ckname/cdf/internal/handler"
	"github.com/cl1ckname/cdf/internal/store"
	"github.com/cl1ckname/cdf/internal/store/catalog"
	"github.com/cl1ckname/cdf/internal/store/filesystem"
)

func Run(arguments ...string) error {
	cdfCatalog := catalog.New("/home/clickname/.config/cdf", filesystem.FS)
	storage := store.New(cdfCatalog)
	if err := storage.Init(); err != nil {
		return err
	}

	call, err := cli.ParseCall(arguments)
	if err != nil {
		return err
	}
	marksHandler := handler.NewHandler(storage)
	if err := marksHandler.Permorm(*call); err != nil {
		return err
	}
	return nil
}
