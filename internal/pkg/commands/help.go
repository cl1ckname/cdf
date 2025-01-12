package commands

import (
	"io"
	"strings"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type Help struct {
	out     io.Writer
	version string
}

func NewHelp(version string, out io.Writer) Help {
	return Help{
		out:     out,
		version: version,
	}
}

func (h Help) Execute(code *domain.Command) error {
	if code == nil {
		return h.writeGeneralMessage()
	}
	return h.writeCommandMessage(*code)
}

func (h Help) writeGeneralMessage() error {
	return h.writeWithHeader(HelpMessage)
}

func (h Help) writeCommandMessage(cmd domain.Command) error {
	descriptions := map[domain.Command]string{
		domain.Command("f"): HelpMessageF,
		domain.CommandAdd:   HelpMessageAdd,
		domain.CommandHelp:  HelpMessageList,
		domain.CommandShell: HelpMessageShell,
		domain.CommandMove:  HelpMessageMove,
	}
	message, ok := descriptions[cmd]
	if !ok {
		return domain.ErrUnknownCommand
	}
	return h.writeWithHeader(message)
}

func (h Help) writeWithHeader(b string) error {
	msg := "CDF version " + h.version + "\n\n"
	if err := h.write(msg); err != nil {
		return err
	}
	return h.write(b)
}

func (h Help) write(msg string) error {
	msg = strings.TrimPrefix(msg, "\n")
	_, err := h.out.Write([]byte(msg))
	return err
}

const HelpMessage = `
usage: [cdf | f] [-h | --help] <command> [<args>]

To get information about specific command use

cdf help <command>

description:
	CDF is a cli path bookmark manager. It consists of 2 commands - cdf and f.

There are different commands of CDF.

move to bookmark:
	f       Moves cwd to path behind specified alias. See "cdf help f" for more information.

manage bookmarks:
	help    Print this message or get information about command
	add     Add new path bookmark
	list    Get list of added marks

service commands for f usage:
	shell   Prints shell commands for f command using and autocomplitions
	move    Provides path of alias to special file
`

const HelpMessageAdd = `
usage: cdf add <alias> <directory>

summary: adds new mark at specified directory

options:
	<alias>
		The name of bookmark that will be used for future moves. Should contain  only a-z A-Z 0-9 _ - / symbols.
	<directory>
		The absolute or relative path to exsisting directory. Future relocations will take place within this directory.

examples:
	Use with absolute path
	cdf add home /home/username

	Use with relative path
	cdf add projects ./projects
`

const HelpMessageList = `
usage: cdf list

summary: prints list of added marks

examples:
	cdf list
`

const HelpMessageF = `
usege: f <alias>

summary: move cwd to path with provided alias

examples:
	f home
`

const HelpMessageShell = `
usage cdf shell [bash | fish | zsh]

summary: prints commands for specified shell for autocompletions and f command using

description:
	This command should be used in file that controls your shell. As part of 
	installing process you should add corresponding line to end of file.
	
	Bash:
		eval "$(cdf shell bash)"
	Fish:
		cdf shell fish | source
	Zas:
		coming soon

exammples:
	cdf shell fish
`

const HelpMessageMove = `
usege cdf move <alias> --cwf-file=<path>

summary: writes path of alias to cwd-file

description:
	In Unix subprocess couldn't change state of parent process, so any cd
	command of command couldn't change cwd of shell. But we can do a trick
	if we define a shell function that performs cd. Exactly this way are using
	in CDF. This command writes path to --cwd-file and f() function of your shell
	reads it and performs cd to it. 
	YOU SHOULDN'T CALL IT BY HANDS, LET f() TO DO IT.

examples:
	cdf move home --cwd-file=/tmp/cdf-cwd-251267
`
