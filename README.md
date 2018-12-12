Templar
=======

Command line templating system written in BASH


Usage
-----

```bash
$ templar
Templar v1.0.1
License(s): http://opensource.org/licenses/MIT

Command line tool to parse a file with ${VAR} style references with their
environment variable or .env type file counterparts.

templar [OPTIONS] ...

OPTIONS:
  -f | -file | --env-file     Use the specified file to populate the
                              template environment.
  -h | -help | --help         Display this help info.
  -n | -dot  | --no-dotenv    Do not automatically load a local .env file.
  -o | -out  | --output-file  Output to the specified file.
  -v | -ver  | --version      Display app version info.

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
```

### Command Line Examples to stdout

```bash
$ CLI_ENV_VAR="Sound and fury" templar example.tmpl -file example.env CLI_VAR="As you like it"
Hello Testiclese!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == '.env Ninja!'
     CLI_ENV_VAR == 'Sound and fury'
         CLI_VAR == 'As you like it'
```

```bash
$ CLI_ENV_VAR="Sound and fury" templar example.tmpl -file example.env CLI_VAR="As you like it" --no-dotenv
Hello Testiclese!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == 'The Bard'
     CLI_ENV_VAR == 'Sound and fury'
         CLI_VAR == 'As you like it'
```

Note the change in `ENV_FILE_VAR` values as `.env` is not loaded in the second example.


#### Command Line Example to files

```bash
$ CLI_ENV_VAR="Sound and fury" templar --output-file output.txt example.tmpl --env-file example.env CLI_VAR="As you like it"
$ cat output.txt
Hello Testiclese!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == '.env Ninja!'
     CLI_ENV_VAR == 'Sound and fury'
         CLI_VAR == 'As you like it'
```

```bash
$ CLI_ENV_VAR="Sound and fury" templar example.tmpl -f example.env CLI_VAR="As you like it" > output.txt
$ cat output.txt
Hello Testiclese!
How do you like it in "/Users/runeimp/dev/apps/templar"?

ENV_FILE_COMMENT == ''
    ENV_FILE_VAR == '.env Ninja!'
     CLI_ENV_VAR == 'Sound and fury'
         CLI_VAR == 'As you like it'
```

Installation
------------

1. Clone this repo to somewhere you like. Maybe `~/repos/templar`
2. `cd ~/bin` or where ever makes sense to access the app in your system.
3. `ln -s ~/repos/templar.bash templar`
4. Verify the install with `templar` in another directory. If you see the Templar help, your done!


TODO
----

* [ ] Support INI files as a data source
* [ ] Support JSON files as a data source
* [ ] Support YAML files as a data source
* [ ] Support TOML files as a data source


Fitness for Use
---------------

I give no garauntees for this scripts fitness for use on any specific computer. I suspect it is "safe enough" to not mangle any untested systems. I've only used this on my own iMac with macOS High Sierra on it. Do not blame me if you try this and your system suffers for it. I share this code with the world so as to be easy for me to reference personally and in the hope that others will find it useful.

