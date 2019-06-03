#!/usr/bin/env bats


#
# BASE TESTS
#
@test "BASE stdout hello Horatio" {
	result="$(templar example.tmpl | grep -F --color=never Hello)"
	[ "$result" = "Hello Horatio!" ]
}

@test "BASE stdout ENV_FILE_COMMENT" {
	result="$(templar example.tmpl | grep -F --color=never ENV_FILE_COMMENT)"
	[ "$result" = "ENV_FILE_COMMENT == ''" ]
}

@test "BASE stdout ENV_FILE_VAR" {
	result="$(templar example.tmpl | grep -F --color=never ENV_FILE_VAR)"
	[ "$result" = "    ENV_FILE_VAR == '.env Ninja!'" ]
}

@test "BASE stdout CLI_ENV_VAR" {
	result="$(templar example.tmpl | grep -F --color=never CLI_ENV_VAR)"
	[ "$result" = "     CLI_ENV_VAR == ''" ]
}

@test "BASE stdout CLI_VAR" {
	result="$(templar example.tmpl | grep -F --color=never CLI_VAR)"
	[ "$result" = "         CLI_VAR == ''" ]
}


#
# NO .ENV TESTS
#
@test "NO DOTENV stdout hello \$USER ($USER)" {
	result="$(templar example.tmpl --no-dotenv | grep -F --color=never Hello)"
	echo "\$USER = $USER"
	[ "$result" = "Hello $USER!" ]
}

@test "NO DOTENV stdout ENV_FILE_COMMENT" {
	result="$(templar example.tmpl --no-dotenv | grep -F --color=never ENV_FILE_COMMENT)"
	[ "$result" = "ENV_FILE_COMMENT == ''" ]
}

@test "NO DOTENV stdout ENV_FILE_VAR" {
	result="$(templar example.tmpl --no-dotenv | grep -F --color=never ENV_FILE_VAR)"
	[ "$result" = "    ENV_FILE_VAR == ''" ]
}

@test "NO DOTENV stdout CLI_ENV_VAR" {
	result="$(templar example.tmpl --no-dotenv | grep -F --color=never CLI_ENV_VAR)"
	[ "$result" = "     CLI_ENV_VAR == ''" ]
}

@test "NO DOTENV stdout CLI_VAR" {
	result="$(templar example.tmpl --no-dotenv | grep -F --color=never CLI_VAR)"
	[ "$result" = "         CLI_VAR == ''" ]
}



#
# CLI VAR TESTS
#
@test "CLI VARS stdout CLI_ENV_VAR" {
	result="$(CLI_ENV_VAR="Sound and fury" templar example.tmpl CLI_VAR="As you like it" | grep -F --color=never CLI_ENV_VAR)"
	[ "$result" = "     CLI_ENV_VAR == 'Sound and fury'" ]
}

@test "CLI VARS stdout CLI_VAR" {
	result="$(CLI_ENV_VAR="Sound and fury" templar example.tmpl CLI_VAR="As you like it" | grep -F --color=never CLI_VAR)"
	[ "$result" = "         CLI_VAR == 'As you like it'" ]
}


#
# WITH ENV FILE
#

@test "WITH ENV FILE stdout hello Testiclese with example.env" {
	result="$(templar example.tmpl --env-file example.env | grep -F --color=never Hello)"
	[ "$result" = "Hello Testiclese!" ]
}

@test "WITH ENV FILE stdout ENV_FILE_COMMENT with example.env" {
	result="$(templar example.tmpl --env-file example.env | grep -F --color=never ENV_FILE_COMMENT)"
	[ "$result" = "ENV_FILE_COMMENT == ''" ]
}

@test "WITH ENV FILE stdout ENV_FILE_VAR" {
	result="$(templar example.tmpl --env-file example.env | grep -F --color=never ENV_FILE_VAR)"
	[ "$result" = "    ENV_FILE_VAR == 'The Bard'" ]
}

@test "WITH ENV FILE and NO DOTENV stdout ENV_FILE_VAR" {
	result="$(templar example.tmpl --env-file example.env --no-dotenv | grep -F --color=never ENV_FILE_VAR)"
	[ "$result" = "    ENV_FILE_VAR == 'The Bard'" ]
}


#
# WITH INI FILE
#

@test "WITH INI FILE stdout hello Testiclese with example.env" {
	result="$(templar example.tmpl --ini-file example.ini | grep -F --color=never Hello)"
	[ "$result" = "Hello Hamlet!" ]
}

@test "WITH INI FILE stdout ENV_FILE_COMMENT with example.env" {
	result="$(templar example.tmpl --ini-file example.ini | grep -F --color=never ENV_FILE_COMMENT)"
	[ "$result" = "ENV_FILE_COMMENT == ''" ]
}

@test "WITH INI FILE stdout ENV_FILE_VAR" {
	result="$(templar example.tmpl --ini-file example.ini | grep -F --color=never ENV_FILE_VAR)"
	[ "$result" = "    ENV_FILE_VAR == '.env Ninja!'" ]
}

@test "WITH INI FILE and NO DOTENV stdout ENV_FILE_VAR" {
	result="$(templar example.tmpl --ini-file example.ini --no-dotenv | grep -F --color=never ENV_FILE_VAR)"
	[ "$result" = "    ENV_FILE_VAR == ''" ]
}


@test "WITH INI FILE global variable" {
	result="$(templar example-ini.tmpl --ini-file example.ini | grep -F --color=never global_ini | cut -d'|' -f1)"
	[ "$result" = "          global_ini == true                         " ]
}

@test "WITH INI FILE section variable" {
	result="$(templar example-ini.tmpl --ini-file example.ini | grep -F --color=never numbers_one | cut -d'|' -f3)"
	[ "$result" = " numbers_one == 1" ]
}

@test "WITH INI FILE section default value" {
	result="$(templar example-ini.tmpl --ini-file example.ini | grep -F --color=never numbers_four | cut -d'|' -f3)"
	[ "$result" = ' numbers_four == "FOUR"' ]
}

@test "WITH INI FILE trimming unquoted value" {
	result="$(templar example-ini.tmpl --ini-file example.ini | grep -F --color=never words_all | cut -d'|' -f1)"
	[ "$result" = '           words_all == "your base are belong to us" ' ]
}

@test "WITH INI FILE preserving quoted value with leading and trailing space" {
	result="$(templar example-ini.tmpl --ini-file example.ini | grep -F --color=never words_system_default | cut -d'|' -f1)"
	[ "$result" = 'words_system_default == " Chrome OS "                ' ]
}















