//
// PACKAGES
//
package main

//
// IMPORTS
//
import (
	"fmt"
	"os"
	"path"
	"strings"

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

const (
	ErrorENVParsing = iota + 50
	ErrorINIParsing
	ErrorJSONParsing
	ErrorTemplateRendering
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

// cli defines the command line interface for this tool
var cli struct {
	DataFile []string `short:"f" help:"Use the specified DATA_FILE to populate the template environment. The file name will be parsed to determine the file type." placeholder:"DATA_FILE" type:"existingfile"`
	Debug    bool     `help:"Show debug info on stderr." hidden`
	EnvFile  string   `short:"e" help:"Use the specified ENV_FILE to populate the template environment." placeholder:"ENV_FILE" type:"existingfile"`
	JSONFile string   `short:"j" help:"Use the specified JSON_FILE to populate the template environment." hidden placeholder:"JSON_FILE" type:"existingfile"`
	IniFile  string   `short:"i" help:"Use the specified INI_FILE to populate the template environment." hidden placeholder:"INI_FILE" type:"existingfile"`
	// KeepLines  bool   `help:"Force empty lines at end of template to be preserved."`
	NoDotenv   bool   `short:"n" help:"Do not load a local .env file if present."`
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
	ctx       *kong.Context
	envFiles  []string
	iniFiles  []string
	jsonFiles []string
)

/*
 * FUNCTIONS
 */

func init() {
	ctx = kong.Parse(&cli, kong.Name(CLIName), kong.Description(AppDesc))
}

/*
 * MAIN ENTRYPOINT
 */
func main() {
	fmt.Printf("%s\n", AppLabel)
	fmt.Printf("%s Library v%s\n", templar.Name, templar.Version)

	// fmt.Printf("templar.main() | cli.DataFile = %q\n", cli.DataFile)
	// if len(cli.DataFile) > 0 {
	// 	ext := path.Ext(cli.DataFile)
	// 	// fmt.Printf("templar.main() | ext = %q\n", ext)
	// 	switch strings.ToUpper(ext) {
	// 	case ".INI":
	// 		iniFiles = append(iniFiles, cli.DataFile)
	// 	case ".JSON":
	// 		jsonFiles = append(jsonFiles, cli.DataFile)
	// 	default:
	// 		fmt.Errorf("Unknown data file type: %q\n", ext)
	// 	}
	// }
	for _, file := range cli.DataFile {
		ext := path.Ext(file)
		switch strings.ToUpper(ext) {
		case ".ENV":
			envFiles = append(envFiles, file)
		case ".INI":
			iniFiles = append(iniFiles, file)
		case ".JSON":
			jsonFiles = append(jsonFiles, file)
		default:
			fmt.Errorf("Unknown data file type: %q\n", ext)
		}
	}
	// fmt.Printf("templar.main() | templar.Data = %#v\n", templar.Data)

	templar.Data.Template = cli.Template

	// fmt.Printf("templar.main() | cli.Template = %q\n", cli.Template)
	// fmt.Printf("templar.main() | templar.Data = %#v\n", templar.Data)
	fmt.Printf("templar.main() |    envFiles = %#v\n", envFiles)
	fmt.Printf("templar.main() |    iniFiles = %#v\n", iniFiles)
	fmt.Printf("templar.main() |   jsonFiles = %#v\n", jsonFiles)

	// for _, file := range templar.Data.INIFile {
	// 	err := templar.ParseINI(file)
	// 	if err != nil {
	// 		fmt.Errorf("JSON Parsing Error: %s", err)
	// 		os.Exit(ErrorJSONParsing)
	// 	}
	// }

	// for _, file := range templar.Data.JSONFile {
	// 	err := templar.ParseJSON(file)
	// 	if err != nil {
	// 		fmt.Errorf("JSON Parsing Error: %s", err)
	// 		os.Exit(ErrorJSONParsing)
	// 	}
	// }

	checkDotEnv := !cli.NoDotenv
	fmt.Printf("templar.main() | checkDotEnv = %t\n", checkDotEnv)

	err := templar.InitData(checkDotEnv, envFiles...)
	if err != nil {
		fmt.Errorf("ENV Parsing Error: %s", err)
		os.Exit(ErrorJSONParsing)
	}

	err = templar.InitData(checkDotEnv, iniFiles...)
	if err != nil {
		fmt.Errorf("INI Parsing Error: %s", err)
		os.Exit(ErrorJSONParsing)
	}

	err = templar.InitData(checkDotEnv, jsonFiles...)
	if err != nil {
		fmt.Errorf("JSON Parsing Error: %s", err)
		os.Exit(ErrorJSONParsing)
	}

	//
	output, err := templar.Render(templar.Data.Template)
	if err != nil {
		fmt.Errorf("Template Rendering Error: %s", err)
		os.Exit(ErrorTemplateRendering)
	}
	fmt.Printf("%s", output)
	// templar.Test()
	println()
	kong.UsageOnError()
	// kong.Help(kong.DefaultHelpPrinter(kong.HelpOptions, ctx))
}
