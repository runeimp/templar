Templar
=======

Command line templating system written in BASH


Usage
-----

```bash
$ templar
Templar v1.0.2
License(s): http://opensource.org/licenses/MIT

Command line tool to parse a file with ${VAR} style references with their
environment variable or .env type file counterparts.

templar [OPTIONS] ...

OPTIONS:
  -d | -debug | --debug        Show debug info on stderr
  -f | -file  | --env-file     Use the specified file to populate the
                               template environment.
  -h | -help  | --help         Display this help info.
  -n | -dot   | --no-dotenv    Do not automatically load a local .env file.
  -o | -out   | --output-file  Output to the specified file.
  -v | -ver   | --version      Display app version info.

```


Example
-------

### `example.tmpl`

```text
Hello ${USER}!
How do you like it in "${PWD}"?

ENV_FILE_COMMENT == '${ENV_FILE_COMMENT}'
    ENV_FILE_VAR == '${ENV_FILE_VAR}'
     CLI_ENV_VAR == '${CLI_ENV_VAR}'
         CLI_VAR == '${CLI_VAR}'

```

### `example.env`

```bash
# ENV_FILE_COMMENT="Not Commented?"
ENV_FILE_VAR="The Bard"
USER=Testiclese
```

### `.env`

```bash
# ENV_FILE_COMMENT="Not Commented?"
ENV_FILE_VAR=".env Ninja!"
USER="Horatio"
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

Note the change in `ENV_FILE_VAR` values between the 1st and 2nd examples as `example.env` is loaded in the 2nd only. ENV files specified on the command line will have it's variables take precidence over variables from a local `.env` file. To have the local .env file take precidence just specify it last on the command line as in the 3rd example.


#### Command Line Example to files

```bash
$ CLI_ENV_VAR="Sound and fury" templar --output-file output.txt example.tmpl --env-file example.env CLI_VAR="As you like it"
$ cat output.txt
Hello Testiclese!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == 'The Bard'
     CLI_ENV_VAR == 'Sound and fury'
         CLI_VAR == 'As you like it'
```

```bash
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

TODO
----

* [ ] Support INI files as a data source
* [ ] Support JSON files as a data source
* [ ] Support YAML files as a data source
* [ ] Support TOML files as a data source


Fitness for Use
---------------

I give no garauntees for this scripts fitness for use on any specific computer. I suspect it is "safe enough" to not mangle any untested systems. I've only used this on my own iMac with macOS High Sierra on it. Do not blame me if you try this and your system suffers for it. I share this code with the world so as to be easy for me to reference personally and in the hope that others will find it useful.



[BATS]: https://github.com/sstephenson/bats

