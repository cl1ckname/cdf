package app

import (
	"os"
	"path/filepath"

	"github.com/cl1ckname/cdf/internal/cli"
	"github.com/cl1ckname/cdf/internal/handler"
	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/fabrics"
	"github.com/cl1ckname/cdf/internal/store"
	"github.com/cl1ckname/cdf/internal/store/catalog"
	"github.com/cl1ckname/cdf/internal/store/filesystem"
)

func Run(arguments ...string) error {
	defaultFolder, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	defaultFolder = filepath.Join(defaultFolder, "cdf")

	cdfCatalog := catalog.New(defaultFolder, filesystem.FS)
	storage := store.New(filesystem.FS, defaultFolder)
	if err = store.Init(cdfCatalog); err != nil {
		return err
	}
	marksFabric := fabrics.NewMarks(filesystem.FS)

	call, err := cli.ParseCall(arguments)
	if err != nil {
		return err
	}

	addCommand := commands.NewAdd(storage, marksFabric)

	marksHandler := handler.NewMarks(addCommand)
	if err := marksHandler.Permorm(*call); err != nil {
		return err
	}
	return nil
}
