Templar
=======

Command line templating system written in BASH


Usage
-----

```bash
$ templar
Templar v1.2.0
License(s): http://opensource.org/licenses/MIT

Command line tool to parse a file with ${VAR} style references with their
environment variable, .env type file counterparts, or INI files.

templar [OPTIONS] TEMPLATE_FILE

OPTIONS:
  -d | -debug | --debug       Show debug info on stderr
  -e | -env   | --env-file ENV_FILE
      Use the specified ENV_FILE to populate the template environment.
  -f | -file  | --data-file DATA_FILE
      Use the specified DATA_FILE to populate the template environment. The
      filenamed will be parsed to determine the file type.
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
Options may appear anywhere on the command line before and/or after the named
template file. Also data file options (--data-file, --env-file, --ini-file) can
be specified more than once and variables
in files specified later on the command line take higher priority.

```


Example
-------

### `example.tmpl`

```bash
Hello ${USER}!
How do you like it in "${PWD}"?

ENV_FILE_COMMENT == '${ENV_FILE_COMMENT}'
    ENV_FILE_VAR == '${ENV_FILE_VAR}'
     CLI_ENV_VAR == '${CLI_ENV_VAR}'
         CLI_VAR == '${CLI_VAR}'

```

### `example-ini.tmpl`

```bash

Hello ${USER}!
How do you like it in "${PWD}"?

Environmentesque Data:

ENV_FILE_COMMENT == '${ENV_FILE_COMMENT}'
    ENV_FILE_VAR == '${ENV_FILE_VAR}'
     CLI_ENV_VAR == '${CLI_ENV_VAR}'
         CLI_VAR == '${CLI_VAR}'

INI Data:

               POSIX == ${POSIX}                     || numbers_one == ${numbers_one}
          global_ini == ${global_ini}                         || numbers_two == ${numbers_two}
           words_all == "${words_all}" || numbers_three == ${numbers_three}
words_system_default == "${words_system_default}"                || numbers_four == "${numbers_four}"

```

### `.env`

```env
# ENV_FILE_COMMENT="Not Commented?"
ENV_FILE_VAR=".env Ninja!"
USER="Horatio"
```

### `example.env`

```env
# ENV_FILE_COMMENT="Not Commented?"
ENV_FILE_VAR="The Bard"
USER=Testiclese
```

### `example.ini`

```ini
; Standard INI comment

POSIX = Awesome!
# POSIX = comment

USER = Hamlet
global_ini: true

[DEFAULT]
system_default : ' Chrome OS '
four =   "FOUR"  

[numbers]
 one= 1
  two = 2.1
   three=3

[words]
all =   your base are belong to us   

```

### Command Line Examples to stdout

```bash
$ templar example.tmpl
Hello Horatio!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == '.env Ninja!'
     CLI_ENV_VAR == ''
         CLI_VAR == ''
```

```bash
$ templar example.tmpl -file example.env
Hello Testiclese!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == 'The Bard'
     CLI_ENV_VAR == ''
         CLI_VAR == ''
```

```bash
$ templar example.tmpl -file example.env -f .env
Hello Horatio!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == '.env Ninja!'
     CLI_ENV_VAR == ''
         CLI_VAR == ''
```

```bash
$ CLI_ENV_VAR="Sound and fury" templar example.tmpl -file example.env CLI_VAR="As you like it"
Hello Testiclese!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == 'The Bard'
     CLI_ENV_VAR == 'Sound and fury'
         CLI_VAR == 'As you like it'
```

```cli
$ CLI_ENV_VAR="Sound and fury" templar example-ini.tmpl -e example.env -i example.ini

Hello Hamlet!
How do you like it in "/Users/runeimp/Dropbox/UserENV/Profile/dev/apps/templar"?

Environmentesque Data:

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == 'The Bard'
     CLI_ENV_VAR == 'Sound and fury'
         CLI_VAR == ''

INI Data:

               POSIX == Awesome!                     || numbers_one == 1
          global_ini == true                         || numbers_two == 2.1
           words_all == "your base are belong to us" || numbers_three == 3
words_system_default == " Chrome OS "                || numbers_four == "FOUR"
```

Note the change in `ENV_FILE_VAR` values between the 1st and 2nd examples as `example.env` is loaded in the 2nd only. ENV files specified on the command line will have it's variables take precidence over variables from a local `.env` file. To have the local .env file take precidence just specify it last on the command line as in the 3rd example.


### INI Files

* Global variables are treated like normal ENV file variables.
* Whitespace before and after keys and values is ignored.
* Quoted values have their quotes removed but leading and trailing whitespace is preserved within the bounds of the quotes. Whitespace outside the quotes is ignored.
* You can separate key and value with a equal sign (standard) or a colon (used in some implementations).
* If a `[DEFAULT]` (or `[default]`) section is present then it's key/value pairs will be populated in all other sections. But will not take precedence over section key/value pairs that are already set.
* To access section variables within a template the section is prepended to the sections keys.
	Example:

	```ini
	[MySection]
	key = value
	```

	can be accessed in a template as `${MySection_key}`. Sections and keys are case sensitive. Though that is counter to the spec it is the norm on POSIX systems.


#### Command Line Example to files

```cli
$ CLI_ENV_VAR="Sound and fury" templar --output-file output.txt example.tmpl --env-file example.env CLI_VAR="As you like it"
$ cat output.txt
Hello Testiclese!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == 'The Bard'
     CLI_ENV_VAR == 'Sound and fury'
         CLI_VAR == 'As you like it'
```

```cli
$ CLI_ENV_VAR="Sound and fury" templar example.tmpl -f example.env CLI_VAR="As you like it" > output.txt
$ cat output.txt
Hello Testiclese!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == 'The Bard'
     CLI_ENV_VAR == 'Sound and fury'
         CLI_VAR == 'As you like it'
```

Installation
------------

1. Clone this repo to somewhere you like. Maybe `~/repos/templar`
2. `cd ~/bin` or where ever makes sense to access the app in your system.
3. `ln -s ~/repos/templar.bash templar`
4. Verify the install with `templar` in another directory. If you see the Templar help, your done!


Test Suite
----------

A small test suite is now part of the Templar repo. This uses [BATS][] to help ensure that Templar is working properly and continues to do so. This has only been tested on macOS High Sierra thus far. Please install [BATS][] and test before use on your system and report issue to <https://github.com/runeimp/templar/issues>.

Running the test suite should look something like this:

```cli
$ bats test_suite.bats
 ✓ BASE stdout hello Horatio
 ✓ BASE stdout ENV_FILE_COMMENT
 ✓ BASE stdout ENV_FILE_VAR
 ✓ BASE stdout CLI_ENV_VAR
 ✓ BASE stdout CLI_VAR
 ✓ NO DOTENV stdout hello $USER (runeimp)
 ✓ NO DOTENV stdout ENV_FILE_COMMENT
 ✓ NO DOTENV stdout ENV_FILE_VAR
 ✓ NO DOTENV stdout CLI_ENV_VAR
 ✓ NO DOTENV stdout CLI_VAR
 ✓ CLI VARS stdout CLI_ENV_VAR
 ✓ CLI VARS stdout CLI_VAR
 ✓ WITH ENV FILE stdout hello Testiclese with example.env
 ✓ WITH ENV FILE stdout ENV_FILE_COMMENT with example.env
 ✓ WITH ENV FILE stdout ENV_FILE_VAR
 ✓ WITH ENV FILE and NO DOTENV stdout ENV_FILE_VAR

16 tests, 0 failures
```

Now up to 25 tests.


TODO
----

* [x] Support INI files as a data source
* [ ] Support JSON files as a data source
* [ ] Support YAML files as a data source
* [ ] Support TOML files as a data source


Fitness for Use
---------------

I give no garauntees for this scripts fitness for use on any specific computer. I suspect it is "safe enough" to not mangle any untested systems. I've only used this on my own iMac with macOS High Sierra on it. Do not blame me if you try this and your system suffers for it. I share this code with the world so as to be easy for me to reference personally and in the hope that others will find it useful.



[BATS]: https://github.com/sstephenson/bats

