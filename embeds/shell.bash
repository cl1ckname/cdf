function f() {
	local tmp="$(mktemp -t "cdf-cwd.XXXXXX")" cwd
	cdf move "$@" --cwd-file="$tmp"
	if cwd="$(command cat -- "$tmp")" && [ -n "$cwd" ] && [ "$cwd" != "$PWD" ]; then
		builtin cd -- "$cwd"
	fi
	rm -f -- "$tmp"
}


# cdf move completion for bash
_cdf_move_completion() {
    # Define the marks file location
    local marks_file="$HOME/.config/cdf/marks"
    if [[ "${COMP_WORDS[1]}" == "remove" || "${COMP_WORDS[1]}" == "move" || "${COMP_WORDS[0]}" == "f" ]]; then
		# Check if the marks file exists
		if [[ -f "$marks_file" ]]; then
			# Read the lines from the marks file
			COMPREPLY=($(compgen -W "$(cat $marks_file | cut -d = -f 1)" -- "${COMP_WORDS[COMP_CWORD]}"))
		fi
	elif [[ "${COMP_WORDS[0]}" == "cdf" && "${COMP_WORDS[1]}" == "help" ]]; then
		COMPREPLY=($(compgen -W "help f add move remove list shell" -- "${COMP_WORDS[COMP_CWORD]}"))
	elif [[ "${COMP_WORDS[0]}" == "cdf" ]]; then
		COMPREPLY=($(compgen -W "help add move remove list shell" -- "${COMP_WORDS[COMP_CWORD]}"))
	fi
}

# Attach the completion to the `cdf move` command
complete -F _cdf_move_completion cdf move
complete -F _cdf_move_completion f
