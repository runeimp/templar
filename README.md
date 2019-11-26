Templar
=======

Command line templating system written in Go. Though the initial idea was written in BASH. And that was cool; but that version had serious limitations as well. Go to the rescue!

Features
--------

* [Mustache][] template system
* POSIX style variable exansion of `$VAR` and `${VAR}` type references in template via the prior Bash based system. See `extra/templar.bash`.
* Multiple data sources available and planned
	* [x] System environment variables
	* [x] `.env` file variables
		* [x] Automatic loading of local `.env` file variables
		* [x] Option to exclude automatic loading of a local `.env` file
	* [x] INI file data
	* [x] JSON file data
	* [ ] YAML file data
	* [ ] TOML file data
	* [ ] XML file data?
	* [ ] CSV file data?
	* [ ] SQLite database query?


Usage
-----

```bash
$ templar template.mustache
```


Example
-------


### INI Files

* Global variables are treated like normal ENV file variables.
* Whitespace before and after keys and values is ignored.
* Quoted values have their quotes removed but leading and trailing whitespace is preserved within the bounds of the quotes. Whitespace outside the quotes is ignored.
* You can separate key and value with a equal sign (standard) or a colon (used in some implementations).
* Global and DEFAULT section key/value pairs can be accessed directly `{{key}}` or via `{{DEFAULT.key}}` references.
* To access section variables within a template the section is prepended to the sections keys.
	Example:

	```ini
	[MySection]
	key = value
	```

	can be accessed in a template as `{{MySection.key}}`. Sections and keys are case sensitive. Though that is counter to the spec it is the norm on POSIX systems.


Installation
------------

1. Download a release from `https://github.com/runeimp/templar/releases` and copy the binary to a directory in your PATH


Test Suite
----------

Test uses Go's normal unit testing tools at this time.


Fitness for Use
---------------

I give no guarantees for this tools fitness for use on any specific system. I suspect it is "safe enough" for use on any untested systems. I've only used this on my own iMac with macOS Mojave on it. I share this code with the world so as to be easy for me to reference personally and in the hope that others will find it useful. Though I strive for portability and usefulness I make no claim that it will always work as expected in any given situation on any system tested.




[Mustache]: https://mustache.github.io/

