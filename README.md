Templar
=======

Command line templating system written in Go. Though the initial idea was written in BASH. And that was cool; but that version had serious limitations as well. Go to the rescue!

Features
--------

* POSIX like variable exansion of `$VAR` and `${VAR}` type references in template
* Multiple data sources available and planned
	* [x] System environment variables
	* [x] `.env` file variables
		* [x] Automatic loading of local `.env` file variables
	* [x] INI file data
	* [ ] JSON file data
	* [ ] YAML file data
	* [ ] TOML file data
	* [ ] XML file data?
	* [ ] CSV file data?
	* [ ] SQLite database query?


Usage
-----

```bash
$ templar
```


Example
-------


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


Installation
------------

1. Clone this repo to somewhere you like. Maybe `~/repos/templar`
2. `cd ~/bin` or where ever makes sense to access the app in your system for your user.
3. `ln -s ~/repos/templar.bash templar`
4. Verify the install with `templar` in another directory. If you see the Templar help, your done!


Test Suite
----------

Test uses Go's normal unit testing tools at this time.


Fitness for Use
---------------

I give no garauntees for this tools fitness for use on any specific system. I suspect it is "safe enough" for use on any untested systems. I've only used this on my own iMac with macOS Mojave on it. I share this code with the world so as to be easy for me to reference personally and in the hope that others will find it useful. Though I strive for portability and usefulness I make no claim that it will always work as expected in any given situation on any system tested.

