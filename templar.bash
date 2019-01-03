#!/usr/bin/env bash
###################
# Templar
#
# @see https://stackoverflow.com/questions/2914220/bash-templating-how-to-build-configuration-files-from-templates-with-bash#answer-42314773
#
#####
# ChangeLog:
# ----------
# 2019-01-02_18:00  v1.1.0      Added template option, updated env file priorities and docs
# 2018-12-12_02:06  v1.0.2      Cleaned up help and updated README.md
# 2018-12-12_02:00  v1.0.1      Updated help with debug reference
# 2018-12-11_19:57  v1.0.0      Initial script creation
#
#


#
# APP INFO
#
APP_AUTHOR="RuneImp <runeimp@gmail.com>"
APP_DESC=$(cat <<-DESC
	Command line tool to parse a file with \${VAR} style references with their
	environment variable or .env type file counterparts.
DESC
)
APP_LICENSES="http://opensource.org/licenses/MIT"
APP_NAME="Templar"
APP_VERSION="1.1.0"
CLI_NAME="templar"


#
# CONSTANTS
#
declare -r APP_LABEL="$APP_NAME v$APP_VERSION
License(s): $APP_LICENSES"


#
# VARIABLES
#
declare -a ARGS
declare -a ENV_FILES
declare -a VARS
declare -i debug_output=1 # boolean
declare -i no_env=1 # boolean
declare output_file=''
declare template=''


#
# FUNCTIONS
#

debug()
{
	local label="$1"
	local msg="$2"

	if [[ $debug_output -eq 0 ]]; then
		printf "DEBUG: %10.10s %s\n" "$label" "$msg" 1>&2
	fi
}

export_cli_vars()
{
	# echo "export_cli_vars():" "${VARS[@]} (${#VARS[@]})"
	for env in "${VARS[@]}"; do
		debug "CLI env:" "$env"
		export "$env"
	done
}

export_file_vars()
{
	local line=""
	local quote_re='^(.+)=(['"'"'"])(.+)(['"'"'"])'

	debug "FILES:" "\${ENV_FILES[@]} = ${ENV_FILES[@]}"
	for env in ${ENV_FILES[@]}; do
		debug "FILE  env:" "$env"

		while read -r line || [ -n "$line" ]; do
			if [[ ${line:0:1} != '#' ]]; then
				debug "FILE line:" "$line (original)"
				# debug "FILE quote_re:" "$quote_re"
				if [[ $line =~ $quote_re ]]; then
					# debug "FILE line:" "\${BASH_REMATCH[0]} = ${BASH_REMATCH[0]}"
					# if [[ ${#BASH_REMATCH[1]} -gt 0 ]]; then
					# 	debug "FILE line:" "\${BASH_REMATCH[1]} = '${BASH_REMATCH[1]}'"
					# fi
					# if [[ ${#BASH_REMATCH[2]} -gt 0 ]]; then
					# 	debug "FILE line:" "\${BASH_REMATCH[2]} = '${BASH_REMATCH[2]}'"
					# fi
					# if [[ ${#BASH_REMATCH[3]} -gt 0 ]]; then
					# 	debug "FILE line:" "\${BASH_REMATCH[3]} = '${BASH_REMATCH[3]}'"
					# fi
					# if [[ ${#BASH_REMATCH[4]} -gt 0 ]]; then
					# 	debug "FILE line:" "\${BASH_REMATCH[4]} = '${BASH_REMATCH[4]}'"
					# fi
					# if [[ ${#BASH_REMATCH[5]} -gt 0 ]]; then
					# 	debug "FILE line:" "\${BASH_REMATCH[5]} = '${BASH_REMATCH[5]}'"
					# fi
					line="${BASH_REMATCH[1]}=${BASH_REMATCH[3]}"
					debug "FILE line:" "$line (parsed)"
				fi
				export "$line"
			fi
		done < $env
	done

	# cut -d= -f1
}

renderer()
{
	local line=""
	local output="cat << EOF"

	while IFS= read -r line; do
		debug "TMPL line:" "$line"
	    output="${output}\n$line"
	done < /dev/stdin
	output="${output}\nEOF"
	output="$(printf "$output" | bash)"

	if [[ ${#output_file} -gt 0 ]]; then
		echo "$output" > "$output_file"
	else
		echo "$output"
	fi
		
}

show_help()
{
	cat <<-EOH
	$APP_LABEL

	$APP_DESC

	$CLI_NAME [OPTIONS] TEMPLATE_FILE

	OPTIONS:
	  -d | -debug | --debug      Show debug info on stderr
	  -f | -file  | --env-file ENV_FILE
	      Use the specified ENV_FILE to populate the template environment.
	      Can be used more than once. Variables in files specified later on
	      the command line take higher priority.
	  -h | -help  | --help       Display this help info.
	  -n | -dot   | --no-dotenv  Do not load a local .env file.
	  -o | -out   | --output-file OUTPUT_FILE
	      Output to the specified file.
	  -t | -temp | --template TEMPLATE_FILE
	      Specify the template file to render.
	  -v | -ver   | --version    Display app version info.

	NOTE:
	Options may appear anywhere on the command line before or/and after the named
	template file.

EOH
}


#
# CONFIG
#
# if [[ ${#XDG_CONFIG_HOME} -gt 0 ]] && [[ -d "${XDG_CONFIG_HOME}" ]]; then
# 	CONFIG_PATH="${XDG_CONFIG_HOME}/${CLINAME}.env"
# elif [[ -d "~/.config" ]]; then
# 	CONFIG_PATH="~/.config/${CLINAME}.env"
# elif [[ ${#XDG_DATA_HOME} -gt 0 ]] && [[ -d "${XDG_DATA_HOME}" ]]; then
# 	CONFIG_PATH="${XDG_DATA_HOME}/${CLINAME}.env"
# elif [[ -d "~/.local" ]]; then
# 	if [[ -d "~/.local/share" ]]; then
# 		CONFIG_PATH="~/.local/share/${CLINAME}.env"
# 	else
# 		CONFIG_PATH="~/.local/${CLINAME}.env"
# 	fi
# else
# 	CONFIG_PATH="~/.${CLINAME}"
# fi
# if [[ ! -d "${CONFIG_PATH}" ]]; then
# 	printf "Creating $APP_NAME configuration directory at:\n    $CONFIG_PATH"
# 	mkdir -p "${CONFIG_PATH}"
# fi


#
# OPTION PARSING
#
if [[ $# -eq 0 ]]; then
	show_help
else
	until [[ $# -eq 0 ]]; do
		case $1 in
		-d | -debug | --debug)
			debug_output=0
			;;
		-f | -file | --env-file)
			ENV_FILES=( ${ENV_FILES[@]} "$2" )
			shift
			;;
		-h | -help | --help)
			show_help
			exit 0
			;;
		-n | -dot | --no-dotenv)
			no_env=0
			;;
		-o | -out | --output-file)
			output_file="$2"
			shift
			;;
		-t | -temp | --template)
			template="$2"
			shift
			;;
		-v | -ver | --version)
			echo "$APP_LABEL"
			exit 0
			;;
		*=*)
			VARS=( "${VARS[@]}" "$1" )
			;;
		*)
			ARGS=( ${ARGS[@]} "$1" )
			;;
		esac

		shift
	done
fi


#
# MAIN ENTRYPOINT
#

if [[ ${#template} -gt 0 ]] || [[ ${#ARGS[@]} -gt 0 ]]; then
	if [[ $no_env -ne 0 ]] && [[ -e ".env" ]]; then
		if [[ -r ".env" ]]; then
			ENV_FILES=( ".env" ${ENV_FILES[@]} )
		else
			echo ".env found but not readable."
		fi
	fi

	export_file_vars
	export_cli_vars
	if [[ ${#template} -eq 0 ]] && [[ ${#ARGS[@]} -gt 0 ]]; then
		template="${ARGS[0]}"
	fi
	cat "$template" | renderer
fi

