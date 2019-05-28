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

	arg "github.com/alexflint/go-arg"
	// "github.com/wrfly/ecp"
)

/*
 * CONSTANTS
 */
const (
	AppDesc    AppMetaData = "Command line templating system using BASH style variable expansion referencing with their environment variable, .env type file counterparts, or INI files."
	AppName    AppMetaData = "Templar"
	AppVersion AppMetaData = "0.1.0"
	CLIName    AppMetaData = "templar"
)

/*
 * DERIVED CONSTANTS
 */
var (
	AppLabel = AppMetaData(fmt.Sprintf("%s v%s", string(AppName), string(AppVersion)))
)

/*
 * TYPES
 */
type (
	// AppMetaData defines meta-data about an application
	AppMetaData string

	appArgs struct {
		DataFile   string `arg:"-f" help:"Use the specified DATA_FILE to populate the template environment. The filenamed will be parsed to determine the file type."`
		Debug      bool   `help:"Show debug info on stderr"`
		EnvFile    string `arg:"-e" help:"Use the specified ENV_FILE to populate the template environment."`
		JSONFile   string `arg:"-j help:"Use the specified JSON_FILE to populate the template environment."`
		IniFile    string `arg:"-i" help:"Use the specified INI_FILE to populate the template environment."`
		KeepLines  bool   `arg:"-l" help:"Force empty lines at end of template to be preserved."`
		NoDotEnv   bool   `arg:"-n" help:"Do not load a local .env file if present."`
		OutputFile string `arg:"-o" help:"Output to the specified file."`
		Template   string `arg:"-t" help:"Specify the template file to render."`
		// ____       string `arg:"-_" help:"____"`
		// ____       string `arg:"-_" help:"____"`
	}
)

func (appArgs) Description() string {
	return string(AppDesc)
}

func (appArgs) Version() string {
	return string(AppLabel)
}

/*
 * VARIABLES
 */
var (
	args appArgs
)

/*
 * FUNCTIONS
 */

func dataMapper(key string) {
	var result string
	switch key {
	case "TEMPLAR_VERSION":
		return AppVersion
	default:
		if len(args.IniFile) > 0 {
			// Do INI file data check
		}
		if len(args.JSONFile) > 0 {
			// Do JSON file data check
		}
		os.Getenv(key)
	}
}

func init() {
	// ecp.Default(&args)
	arg.MustParse(&args)
}

// parseTemplate expands variables in the template using dataMapper
func parseTemplate() {
	os.Expand(args.Template, dataMapper)
}

/*
 * MAIN ENTRYPOINT
 */
func main() {
	// log.Printf("%s", AppLabel)
}
