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
	AppDesc    = "Command line templating system based on Mustache template engine and data supplied by environment variables, ENV, INI, and JSON files. And soon YAML, and TOML files as well."
	AppName    = "Templar"
	AppVersion = "2.0.0"
	CLIName    = "templar"
)

const (
	ErrorENVParsing = iota + 50
	ErrorINIParsing
	ErrorJSONParsing
	ErrorTemplateMissing
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
	DataENV    []string `short:"e" help:"Use the specified ENV_FILE (regardless of it's file extension) to populate the template environment." placeholder:"ENV_FILE" type:"existingfile"`
	DataFile   []string `short:"f" help:"Use the specified DATA_FILE to populate the template environment. File type determined by the extension on the file name." placeholder:"DATA_FILE" type:"existingfile"`
	DataINI    []string `short:"i" help:"Use the specified INI_FILE (regardless of it's file extension) to populate the template environment." placeholder:"INI_FILE" type:"existingfile"`
	DataJSON   []string `short:"j" help:"Use the specified JSON_FILE (regardless of it's file extension) to populate the template environment." placeholder:"JSON_FILE" type:"existingfile"`
	Debug      bool     `help:"Show debug info on stderr." hidden`
	NoDotenv   bool     `short:"n" help:"Do not load a local .env file if present."`
	OutputFile string   `short:"o" help:"Output to the specified file." placeholder:"FILE" sep:' ' type:"existingfile"`
	Template   string   `arg optional help:"Specify the template file to render." type:"existingfile"`
	// ____       string `arg:"-_" help:"____"`
	// ____       string `arg:"-_" help:"____"`
}

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
	if cli.Debug {
		fmt.Printf("%s\n", AppLabel)
		fmt.Printf("%s Library v%s\n", templar.Name, templar.Version)
	}

	if len(cli.Template) == 0 {
		if len(cli.DataENV) == 0 && len(cli.DataFile) == 0 && len(cli.DataINI) == 0 && len(cli.DataJSON) == 0 && len(cli.OutputFile) == 0 && cli.NoDotenv == false {
			ctx.PrintUsage(false)
			os.Exit(0)
		}

		fmt.Fprintf(os.Stderr, "Usage Error: %s\n", "Template Missing")
		os.Exit(ErrorTemplateMissing)
	}

	// fmt.Printf("templar.main() | cli.DataFile = %q\n", cli.DataFile)
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

	for _, file := range cli.DataENV {
		envFiles = append(envFiles, file)
	}
	for _, file := range cli.DataINI {
		iniFiles = append(iniFiles, file)
	}
	for _, file := range cli.DataJSON {
		jsonFiles = append(jsonFiles, file)
	}

	// templar.Data.Template = cli.Template
	checkDotEnv := !cli.NoDotenv

	// fmt.Printf("templar.main() | cli.Template = %q\n", cli.Template)
	// fmt.Printf("templar.main() | templar.Data = %#v\n", templar.Data)
	if cli.Debug {
		fmt.Printf("templar.main() |    envFiles = %#v\n", envFiles)
		fmt.Printf("templar.main() |    iniFiles = %#v\n", iniFiles)
		fmt.Printf("templar.main() |   jsonFiles = %#v\n", jsonFiles)
		fmt.Printf("templar.main() |    template = %q\n", cli.Template)
		fmt.Printf("templar.main() | checkDotEnv = %t\n", checkDotEnv)
	}

	err := templar.InitData(checkDotEnv, envFiles...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ENV File Parsing Error: %s\n", err.Error())
		os.Exit(ErrorJSONParsing)
	}

	err = templar.InitData(checkDotEnv, iniFiles...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "INI File Parsing Error: %s\n", err.Error())
		os.Exit(ErrorJSONParsing)
	}

	err = templar.InitData(checkDotEnv, jsonFiles...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "JSON File Parsing Error: %s\n", err.Error())
		os.Exit(ErrorJSONParsing)
	}

	output, err := templar.Render(cli.Template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Template Rendering Error: %s\n", err.Error())
		os.Exit(ErrorTemplateRendering)
	}
	fmt.Printf("%s", output)
	kong.UsageOnError()
	// kong.Help(kong.DefaultHelpPrinter(kong.HelpOptions, ctx))
}