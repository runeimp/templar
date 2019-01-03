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


