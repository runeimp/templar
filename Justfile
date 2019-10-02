#
# This is a Justfile for https://github.com/casey/just and is much like a Makefile for make
# I highly recommend using just for basic automation.
#

# Like in make the first recipe is used by default.
# I like listing all the recipes by default.
# I also like wiping the terminal buffer like CLS in DOS. It makes me happy.  :-)
@_default:
	just _term-wipe
	just --list


# Run a test
@test cmd="help" +data="":
	just _term-wipe
	just test-{{cmd}} "{{data}}"

# Test with debug enabled
test-debug +data="example.env":
	@# CLI_ENV_VAR="Sound and fury" go run cmd/templar/main.go example.tmpl --data-file example.env CLI_VAR="As you like it" --debug
	CLI_ENV_VAR="Sound and fury" CLI_VAR="As you like it" go run cmd/templar/main.go example.tmpl --data-file {{data}} --debug

# Test the help system
test-help +data="example.env":
	go run cmd/templar/main.go --help

# Test example.env with .env
test-stdout +data="example.json -f example2.json":
	@# CLI_ENV_VAR="Sound and fury" go run cmd/templar/main.go example.tmpl --data-file example.env CLI_VAR="As you like it"
	CLI_ENV_VAR="Sound and fury" CLI_VAR="As you like it" go run cmd/templar/main.go example.tmpl --data-file {{data}}

# Test example.env with out .env
test-no-dotenv +data="example.env":
	#!/bin/sh
	if [ -z "{{data}}" ]; then
		# echo 'CLI_ENV_VAR="Sound and fury" CLI_VAR="As you like it" go run cmd/templar/main.go example.tmpl --no-dotenv'
		CLI_ENV_VAR="Sound and fury" CLI_VAR="As you like it" go run cmd/templar/main.go example.tmpl --no-dotenv
	else
		# echo 'CLI_ENV_VAR="Sound and fury" CLI_VAR="As you like it" go run cmd/templar/main.go example.tmpl --data-file {{data}} --no-dotenv'
		CLI_ENV_VAR="Sound and fury" CLI_VAR="As you like it" go run cmd/templar/main.go example.tmpl --data-file {{data}} --no-dotenv
	fi

	# CLI_ENV_VAR="Sound and fury" go run cmd/templar/main.go example.tmpl --data-file example.env CLI_VAR="As you like it" --no-dotenv
	

# Test creating an output file
test-with-file +data="example.env":
	rm output.txt
	@# CLI_ENV_VAR="Sound and fury" go run cmd/templar/main.go --output-file output.txt example.tmpl --env-file example.env CLI_VAR="As you like it"
	CLI_ENV_VAR="Sound and fury" CLI_VAR="As you like it" go run cmd/templar/main.go --output-file output.txt example.tmpl --env-file {{data}}
	cat output.txt


# Wipes the terminal buffer for a clean start
_term-wipe:
	#!/bin/sh
	if [[ ${#VISUAL_STUDIO_CODE} -gt 0 ]]; then
		clear
	elif [[ ${KITTY_WINDOW_ID} -gt 0 ]] || [[ ${#TMUX} -gt 0 ]] || [[ "${TERM_PROGRAM}" = 'vscode' ]]; then
		printf '\033c'
	elif [[ "$(uname)" == 'Darwin' ]] || [[ "${TERM_PROGRAM}" = 'Apple_Terminal' ]] || [[ "${TERM_PROGRAM}" = 'iTerm.app' ]]; then
		osascript -e 'tell application "System Events" to keystroke "k" using command down'
	elif [[ -x "$(which tput)" ]]; then
		tput reset
	elif [[ -x "$(which reset)" ]]; then
		reset
	else
		clear
	fi

