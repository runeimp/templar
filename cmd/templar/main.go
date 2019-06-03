//
// PACKAGES
//
package main

//
// IMPORTS
//
import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"github.com/runeimp/templar"
)

/*
 * CONSTANTS
 */
const (
	AppDesc    = "Command line templating system based on Mustache template engine and data suppied by environment variable, .env file counterparts, and soon INI, JSON, YAML, and TOML files."
	AppName    = "Templar"
	AppVersion = "2.0.0"
	CLIName    = "templar"
)

/*
 * DERIVED CONSTANTS
 */
var (
	AppLabel = fmt.Sprintf("%s v%s", AppName, AppVersion)
)

/*
 * TYPES
 */

// CLI defines the command line interface for this tool
var CLI struct {
	DataFile string `short:"f" help:"Use the specified DATA_FILE to populate the template environment. The file name will be parsed to determine the file type." placeholder:"DATA_FILE" type:"existingfile"`
	Debug    bool   `help:"Show debug info on stderr." hidden`
	EnvFile  string `short:"e" help:"Use the specified ENV_FILE to populate the template environment." placeholder:"ENV_FILE" type:"existingfile"`
	JSONFile string `short:"j" help:"Use the specified JSON_FILE to populate the template environment." hidden placeholder:"JSON_FILE" type:"existingfile"`
	IniFile  string `short:"i" help:"Use the specified INI_FILE to populate the template environment." hidden placeholder:"INI_FILE" type:"existingfile"`
	// KeepLines  bool   `help:"Force empty lines at end of template to be preserved."`
	NoDotEnv   bool   `short:"n" help:"Do not load a local .env file if present."`
	OutputFile string `short:"o" help:"Output to the specified file." placeholder:"FILE" sep:' ' type:"existingfile"`
	Template   string `arg optional help:"Specify the template file to render." type:"existingfile"`
	// ____       string `arg:"-_" help:"____"`
	// ____       string `arg:"-_" help:"____"`
}

// func (appArgs) Description() string {
// 	return string(AppDesc)
// }

// func (appArgs) Version() string {
// 	return string(AppLabel)
// }

/*
 * VARIABLES
 */
var (
	ctx *kong.Context
)

/*
 * FUNCTIONS
 */

func init() {
	ctx = kong.Parse(&CLI, kong.Name(CLIName), kong.Description(AppDesc))

}

/*
 * MAIN ENTRYPOINT
 */
func main() {
	log.Printf("%s\n", AppLabel)
	log.Printf("%s Library v%s\n", templar.Name, templar.Version)

	templar.Test()
	println()
	kong.UsageOnError()
	// kong.Help(kong.DefaultHelpPrinter(kong.HelpOptions, ctx))
}
