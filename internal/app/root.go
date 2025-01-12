package app

import (
	"os"
	"path/filepath"

	"github.com/cl1ckname/cdf/embeds"
	"github.com/cl1ckname/cdf/internal/cli"
	"github.com/cl1ckname/cdf/internal/handler"
	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/pkg/fabrics"
	"github.com/cl1ckname/cdf/internal/pkg/presenters"
	"github.com/cl1ckname/cdf/internal/store"
	"github.com/cl1ckname/cdf/internal/store/catalog"
	"github.com/cl1ckname/cdf/internal/store/filesystem"
)

func Run(version string, arguments ...string) error {
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

	helpCommand := commands.NewHelp(version, os.Stdout)

	marksFabric := fabrics.NewMarks(filesystem.FS)
	addCommand := commands.NewAdd(storage, marksFabric)

	presenter := presenters.NewList(os.Stdout)
	listCommand := commands.NewList(storage, presenter)

	removeCommand := commands.NewRemove(storage)

	moveCommand := commands.NewMove(storage)

	shellCommand := commands.NewShell(os.Stdout, commands.Wraps{
		domain.FishShell: embeds.FishShell,
		domain.BashShell: embeds.BashShell,
	})

	marksHandler := handler.NewMarks(
		helpCommand,
		addCommand,
		listCommand,
		removeCommand,
		moveCommand,
		shellCommand,
	)

	call, err := cli.ParseCall(arguments)
	if err != nil {
		return err
	}
	if err := marksHandler.Permorm(*call); err != nil {
		return err
	}
	return nil
}
