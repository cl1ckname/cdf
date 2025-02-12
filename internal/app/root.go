package app

import (
	"io"
	"os"
	"path/filepath"

	"github.com/cl1ckname/cdf/embeds"
	"github.com/cl1ckname/cdf/internal/cli"
	"github.com/cl1ckname/cdf/internal/clock"
	"github.com/cl1ckname/cdf/internal/handler"
	"github.com/cl1ckname/cdf/internal/logger"
	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/pkg/fabrics"
	"github.com/cl1ckname/cdf/internal/store"
	"github.com/cl1ckname/cdf/internal/store/catalog"
	"github.com/cl1ckname/cdf/internal/store/filesystem"
	"github.com/cl1ckname/cdf/internal/store/mover"
)

const (
	MarkFilepathArg = "usefile"
	VerboseArg      = "verbose"
)

type System struct {
	Stdout, Stderr io.Writer
	Args           []string
	Version        string
}

func Run(sys System) error {
	args, kwargs, err := cli.ParseFlags(sys.Args)
	if err != nil {
		return err
	}

	_, verbose := kwargs[VerboseArg]
	log := buildLogger(sys, verbose)

	filepath, err := marksFile(log, filesystem.FS, kwargs)
	if err != nil {
		return err
	}
	storage := store.New(filesystem.FS, filepath, log)

	marksHandler := buildHandler(sys, storage, log)

	if verbose {
		_, err = sys.Stdout.Write([]byte{'\n'})
		if err != nil {
			return err
		}
	}

	call, err := cli.NewCall(args, kwargs)
	if err != nil {
		return err
	}
	if err := marksHandler.Permorm(*call); err != nil {
		return err
	}
	return nil
}

func buildLogger(sys System, verbose bool) logger.Logger {
	loglevel := logger.Error
	if verbose {
		loglevel = logger.Info
	}
	return logger.New(sys.Stdout, sys.Stderr, loglevel)
}

func buildHandler(sys System, storage commands.Store, log logger.Logger) handler.Marks {
	base := commands.NewBase(storage, log)
	helpCommand := commands.NewHelp(sys.Version, sys.Stdout)
	marksFabric := fabrics.NewMarks(filesystem.FS, clock.Time)
	addCommand := commands.NewAdd(base, marksFabric)

	listCommand := commands.NewList(base, fabrics.PresenterInstance)

	removeCommand := commands.NewRemove(base)

	mvr := mover.NewMover(filesystem.FS)
	moveCommand := commands.NewMove(base, mvr, clock.Time)

	shellCommand := commands.NewShell(sys.Stdout, commands.Wraps{
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
	return marksHandler
}

func marksFile(log logger.Logger, fs catalog.FS, kwargs handler.Kwargs) (string, error) {
	path, ok := kwargs[MarkFilepathArg]
	if ok {
		log.Info("marks file specifies, lookup at", path)
		return path, catalog.EnsureFile(log, path, fs)
	}
	log.Info("marks file not specified, use default")
	return defaultMarksPath(log, fs)
}

func defaultMarksPath(log logger.Logger, fs catalog.FS) (string, error) {
	defaultFolder, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	defaultFolder = filepath.Join(defaultFolder, "cdf")

	filepath, err := catalog.InitInFolder(log, defaultFolder, fs)
	if err != nil {
		return "", err
	}
	return filepath, nil
}
