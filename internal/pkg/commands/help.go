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
		domain.Command("f"):  HelpMessageF,
		domain.CommandAdd:    HelpMessageAdd,
		domain.CommandList:   HelpMessageList,
		domain.CommandHelp:   HelpMessage,
		domain.CommandRemove: HelpMessageRemove,
		domain.CommandShell:  HelpMessageShell,
		domain.CommandMove:   HelpMessageMove,
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
usage: [cdf | f] [-h | --help] [--usefile] [--verbose] <command> [<args>]

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
	remove  Remove mark by alias

service commands for f usage:
	shell   Prints shell commands for f command using and autocomplitions
	move    Provides path of alias to special file

options:
	--usefile=<filepath>
		Customise path to file with marks saves. Creates it if not exists. Default is ~/.local/share/cdf.

	--verbose
		Shows more information about command execution. Information messages are writing into stdout,
		warnings and errors into stderr.
`

const HelpMessageAdd = `
usage: cdf add <alias> [directory]

summary: adds new mark at specified directory

options:
	<alias>
		The name of bookmark that will be used for future moves. Should contain  only a-z A-Z 0-9 _ - / symbols.
	<directory>
		The absolute or relative path to exsisting directory. If not presented cwd is used. Future relocations will take place within this
		directory.

examples:
	Use with absolute path
	cdf add home /home/username

	Use with relative path
	cdf add projects ./projects
`

const HelpMessageList = `
usage: cdf list [--format] [--long | -l]

summary: prints list of added marks.

options:
	--format=<FORMAT>
		Determines the form of outputs. Should be one of the next valus:
			
			• default
				The output is TAB separated alias - value pairs.

			• json
				The output will be a single JSON object where alias are keys and marks are values.
	
			• alias
				The output will contain only list of all aliases,

		Otherwise it will be default.

	-l, --long
		Add more field (e.g. usage status) to outputs. Each format use it in its own way. 

examples:
	cdf list

	cdf list --format=json

	cdf list -l
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

const HelpMessageRemove = `
usege: cdf remove <alias>

summary: removes mark from list by alias

exaples:
	cdf remove home
`

const HelpMessageMove = `
usege: cdf move <alias>

summary: prints path of alias

description:
	In Unix subprocess couldn't change state of parent process, so any cd
	command of command couldn't change cwd of shell. But we can do a trick
	if we define a shell function that performs cd. Exactly this way are using
	in CDF. This command writes path to stdout and f() function of your shell
	reads it and performs cd to it. 
	YOU SHOULDN'T CALL IT BY HANDS, LET f() TO DO IT.

examples:
	cdf move home
`
