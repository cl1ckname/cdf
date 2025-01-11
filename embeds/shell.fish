function f
	set tmp (mktemp -t "cdf-cwd.XXXXXX")
	cdf move $argv --cwd-file="$tmp"
	if set cwd (command cat -- "$tmp"); and [ -n "$cwd" ]; and [ "$cwd" != "$PWD" ]
		builtin cd -- "$cwd"
	end
	rm -f -- "$tmp"
end

set arguments_file "~/.config/cdf/marks"
complete -f -c cdf -n __fish_use_subcommand -a add -d "Add mark"
complete -f -c cdf -n __fish_use_subcommand -a move -d "Move to mark"
complete -f -c cdf -n __fish_use_subcommand -a list -d "List marks"
complete -f -c cdf -n __fish_use_subcommand -a shell -d "Wrap shell with helpers"
complete -f -c cdf -n '__fish_seen_subcommand_from move' -a "(cat $arguments_file | cut -d = -f 1)"
complete -f -c cdf -n '__fish_seen_subcommand_from shell' -a "fish bash zsh"
