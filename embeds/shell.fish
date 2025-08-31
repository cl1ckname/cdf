function f
	set cwd (cdf move $argv)
	if [ -n "$cwd" ]; and [ "$cwd" != "$PWD" ]
		builtin cd -- "$cwd"
	end
end

set arguments_file "~/.local/share/cdf/marks"
set aliases (cdf list --format=alias)
complete -f -c f  -a $aliases
complete -f -c cdf -n __fish_use_subcommand -a help -d "Get help"
complete -f -c cdf -n __fish_use_subcommand -a add -d "Add mark"
complete -f -c cdf -n __fish_use_subcommand -a move -d "Move to mark"
complete -f -c cdf -n __fish_use_subcommand -a remove -d "Remove mark"
complete -f -c cdf -n __fish_use_subcommand -a list -d "List marks"
complete -f -c cdf -n __fish_use_subcommand -a shell -d "Wrap shell with helpers"

complete -f -c cdf -n '__fish_seen_subcommand_from move' -a $aliases
complete -f -c cdf -n '__fish_seen_subcommand_from remove' -a $aliases
complete -f -c cdf -n '__fish_seen_subcommand_from shell' -a "fish bash zsh"
complete -f -c cdf -n '__fish_seen_subcommand_from help' -a "help f add move list shell"
