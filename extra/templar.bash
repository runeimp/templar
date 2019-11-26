#!/usr/bin/env bash
###################
# Templar
#
# @see https://stackoverflow.com/questions/2914220/bash-templating-how-to-build-configuration-files-from-templates-with-bash#answer-42314773
#
#####
# ChangeLog:
# ----------
# 2019-01-03_22:45  v1.2.0      Added INI data file support
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
	environment variable, .env type file counterparts, or INI files.
DESC
)
APP_LICENSES="http://opensource.org/licenses/MIT"
APP_NAME="Templar"
APP_VERSION="1.2.0"
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
declare -a FILES_DATA
declare -a FILES_ENV
declare -a FILES_INI
declare -a FILES_JSON
declare -a VARS
declare -i debug_output=1 # boolean false
declare -i keep_lines=1 # boolean false
declare -i no_env=1 # boolean false
declare output_file=''
declare template=''


#
# FUNCTIONS
#

# Routes data files based on extesion
data_files_router()
{
	local data_file=""
	local file_ext=""

	debug "DATA file:" "data_files_router() | \${#FILES_DATA[@]} = ${#FILES_DATA[@]}"

	shopt -s nocasematch

	for data_file in ${FILES_DATA[@]}; do
		# Parse filename
		file_ext="${data_file#*.}"
		debug "DATA file:" "data_files_router() | \$data_file = $data_file | \${data_file#*.} = $file_ext"
		
		case "$file_ext" in
			ENV)
				# ENV File
				debug "ENV_FILE $data_file"
				FILES_ENV=( ${FILES_ENV[@]} "$data_file" )
				;;
			INI)
				# INI File
				debug "INI_FILE $data_file"
				FILES_INI=( ${FILES_INI[@]} "$data_file" )
				;;
			JSON)
				# JSON File
				debug "JSON_FILE $data_file"
				FILES_JSON=( ${FILES_JSON[@]} "$data_file" )
				;;
			*)
				# Unhandled File
				echo "Unhandled file extension '$file_ext'" 1>&2
				echo "Skipping '$data_file'" 1>&2
				;;
		esac
	done

	shopt -u nocasematch
}

debug()
{
	local label="$1"
	local msg="$2"

	if [[ $debug_output -eq 0 ]]; then
		printf "DEBUG: %12.12s %s\n" "$label" "$msg" 1>&2
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

export_env_file_vars()
{
	local line=""
	local quote_re='^(.+)=(['"'"'"])(.+)(['"'"'"])'

	debug "FILES:" "\${FILES_ENV[@]} = ${FILES_ENV[@]}"
	for env in ${FILES_ENV[@]}; do
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

export_ini_file_vars()
{
	local -a default_section
	local last_section=""
	local line=""
	local value=""

	local keyvalue_re='^[	 ]*([a-zA-Z_]+)[	 ]*[=:][	 ]*['"'"'"]?(.+)['"'"'"]?$'
	local section_re='^\[(.+)\]$'

	debug "FILES:" "\${FILES_INI[@]} = ${FILES_INI[@]}"
	for ini in ${FILES_INI[@]}; do
		debug "FILE  ini:" "$ini"

		while read -r line || [ -n "$line" ]; do
			if [[ $line =~ $section_re ]]; then
				debug "INI line:" "$line (original)"
				# debug "INI section:" "\${BASH_REMATCH[0]} = ${BASH_REMATCH[0]}"
				# if [[ ${#BASH_REMATCH[1]} -gt 0 ]]; then
				# 	debug "INI section:" "\${BASH_REMATCH[1]} = '${BASH_REMATCH[1]}'"
				# fi
				# if [[ ${#BASH_REMATCH[2]} -gt 0 ]]; then
				# 	debug "INI section:" "\${BASH_REMATCH[2]} = '${BASH_REMATCH[2]}'"
				# fi
				# line="${BASH_REMATCH[1]}"
				last_section="${BASH_REMATCH[1]}"
				debug "INI section:" "$last_section (parsed)"
				if [[ "$last_section" != 'DEFAULT' ]] || [[ "$last_section" != 'default' ]]; then
					if [[ ${#default_section[@]} -gt 0 ]]; then
						section_defaults_add "$last_section"
					fi
				fi
			elif [[ $line =~ $keyvalue_re ]]; then
				debug "INI line:" "$line (original)"
				# debug "INI key:" "\${BASH_REMATCH[0]} = ${BASH_REMATCH[0]}"
				# [[ ${#BASH_REMATCH[1]} -gt 0 ]] && debug "INI key:" "\${BASH_REMATCH[1]} = '${BASH_REMATCH[1]}'"
				if [[ ${#BASH_REMATCH[2]} -gt 0 ]]; then
					value="$(strip_quotes "${BASH_REMATCH[2]}")"
					# debug "INI key:" "\${BASH_REMATCH[2]} = '$value'"
				fi
				# [[ ${#BASH_REMATCH[3]} -gt 0 ]] && debug "INI key:" "\${BASH_REMATCH[3]} = '${BASH_REMATCH[3]}'"

				if [[ "$last_section" = 'DEFAULT' ]] || [[ "$last_section" = 'default' ]]; then
					line="${BASH_REMATCH[1]}=$value"
					default_section=( "${default_section[@]}" "$line" )
				else
					if [[ ${#last_section} -gt 0 ]]; then
						line="${last_section}_${BASH_REMATCH[1]}=$value"
					else
						line="${BASH_REMATCH[1]}=$value"
					fi
					export "$line"
				fi
				debug "INI key:" "$line (parsed)"
			fi
		done < $ini
	done
}

fold_line()
{
	local -i cols=$1
	local -i i=$1
	local indent=""
	local indent_re="^([	 ]+).*$"
	local line="$2"
	
	if [[ ${#line} -gt $cols ]]; then
		while [[ "${line:$i:1}" != " " ]]; do
			if [[ $i -gt 0 ]]; then
				# echo "fold_line() | \$i = $i"
				let "i -= 1"
			else
				break
			fi
		done
		echo "${line:0:$i}"

		if [[ "${line:0:$i}" =~ $indent_re ]]; then
			indent="${BASH_REMATCH[1]}"
		fi
		let "i += 1"
		fold_line $cols "${indent}${line:$i}"
	else
		echo "$line"
	fi
}

fold_stdin()
{
	local -i cols=80
	local -i i=0

	if [[ -x "$(which tput)" ]]; then
		cols=$(tput cols)
	fi

	set -f

	while IFS= read -r line; do
	    if [[ ${#line} -gt $cols ]]; then
		    fold_line $cols "$line"
		else
			echo "$line"
		fi
	done < /dev/stdin

	set +f
}

renderer()
{
	local line=""
	local output="cat <<EOF"

	set -f

	while IFS= read -r line; do
		debug "TMPL line:" "$line"
	    output="${output}\n$line"
	done < /dev/stdin
	if [[ $keep_lines -eq 0 ]]; then
		output="${output} " # SP (space)
		# output="${output}\r" # CR (carriage return)
	fi
	output="${output}\nEOF"
	# output="$(printf "$output")"
	output="$(printf "${output}" | bash)"

# 	output="$(eval "cat <<EOF
# $(</dev/stdin)
# EOF
# " 2> /dev/null)"

	if [[ ${#output_file} -gt 0 ]]; then
		echo "$output" > "$output_file"
	else
		echo "$output"
	fi

	set +f
}

section_defaults_add()
{
	local section="$1"

	for default in "${default_section[@]}"; do
		debug "section_defaults_add() | \$default = $default"
		line="${section}_${default}"
		debug "section_defaults_add() | \$line = $line"
		export "$line"
	done
}

strip_quotes()
{
	local value="$1"
	local last_char=""
	local first_char=""

	# echo "strip_quotes() | \$value == \"$value\" (input)" 1>&2

	first_char="${value:0:1}"

	if [[ $first_char = "'" ]] || [[ $first_char = '"' ]]; then
		value="${value:1}"
	fi

	let "minus_one = ${#value} - 1"
	last_char="${value:$minus_one}"

	if [[ $last_char = "'" ]] || [[ $last_char = '"' ]]; then
		value="${value:0:$minus_one}"
	fi

	# echo "strip_quotes() | \$first_char == $first_char" 1>&2
	# echo "strip_quotes() | \$last_char == $last_char" 1>&2
	# echo "strip_quotes() | \$value == \"$value\" (output)" 1>&2
	echo "$value"
}

show_help()
{
	local msg=$(cat <<-EOH
	$APP_LABEL

	$APP_DESC

	$CLI_NAME [OPTIONS] TEMPLATE_FILE

	OPTIONS:
	  -d | -debug | --debug       Show debug info on stderr
	  -e | -env   | --env-file ENV_FILE
	      Use the specified ENV_FILE to populate the template environment.
	  -f | -file  | --data-file DATA_FILE
	      Use the specified DATA_FILE to populate the template environment. The filenamed will be parsed to determine the file type.
	  -h | -help  | --help        Display this help info.
	  -i | -ini   | --ini-file INI_FILE
	      Use the specified INI_FILE to populate the template environment.
	  -l | -lines | --keep-lines
	      Force empty lines at end of template to be preserved.
	  -n | -nodot | --no-dotenv   Do not load a local .env file.
	  -o | -out   | --output-file OUTPUT_FILE
	      Output to the specified file.
	  -t | -temp  | --template TEMPLATE_FILE
	      Specify the template file to render.
	  -v | -ver   | --version     Display app version info.

	NOTE:
	Options may appear anywhere on the command line before and/or after the named template file. Also data file options (--data-file, --env-file, --ini-file) can be specified more than once and variables in files specified later on the command line take higher priority.

EOH
)
	# printf "%s\n\n" "$msg"
	printf "%s\n\n" "$msg" | fold_stdin
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
		-e | -env | --env-file)
			FILES_ENV=( ${FILES_ENV[@]} "$2" )
			shift
			;;
		-f | -file | --data-file)
			FILES_DATA=( ${FILES_DATA[@]} "$2" )
			shift
			;;
		-h | -help | --help)
			show_help
			exit 0
			;;
		-i | -ini | --ini-file)
			FILES_INI=( ${FILES_INI[@]} "$2" )
			shift
			;;
		# -j | -json | --json-file)
		# 	FILES_JSON=( ${FILES_JSON[@]} "$2" )
		# 	shift
		# 	;;
		-l | -lines | --keep-lines)
			keep_lines=0
			;;
		-n | -nodot | --no-dotenv)
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
			ARGS=( "${ARGS[@]}" "$1" )
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
			FILES_ENV=( ".env" ${FILES_ENV[@]} )
		else
			echo ".env found but not readable."
		fi
	fi

	data_files_router
	export_env_file_vars
	export_ini_file_vars
	export_cli_vars
	if [[ ${#template} -eq 0 ]] && [[ ${#ARGS[@]} -gt 0 ]]; then
		template="${ARGS[0]}"
	fi
	cat "$template" | renderer
fi

