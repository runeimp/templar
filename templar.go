package templar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/cbroglie/mustache"
	// "github.com/subosito/gotenv"
	"github.com/runeimp/gotenv"
	ini "gopkg.in/ini.v1"
)

/*
 * CONSTANTS
 */
const (
	// Name denotes the library name
	Name = "Templar"

	// Version denotes the library version
	Version = "0.2.0"
)

/*
 * CONSTANTS
 */
const (
	DebugOff   = 0
	DebugLog   = 1
	DebugInfo  = 2
	DebugWarn  = 3
	DebugError = 4
)

/*
 * VARIABLES
 */
var (
	// envBackupData map[string]string
	dataProvider map[string]interface{}
	Debug        = DebugWarn
	initialized  = false
)

/*
 * TYPES
 */

// Data is a collection for all external data
var Data struct {
	INIFile  []string
	JSONFile []string
	Template string
}

/*
 * METHODS
 */
// func envBackup() {
// 	if Debug <= DebugWarn {
// 		fmt.Fprintln(os.Stderr, "templar.envBackup()")
// 	}

// 	envBackupData = make(map[string]string)

// 	for _, envar := range os.Environ() {
// 		kv := strings.Split(envar, "=")
// 		k := kv[0]
// 		v := kv[1]
// 		if Debug <= DebugLog {
// 			fmt.Fprintf(os.Stderr, "templar.envBackup() | envBackupData[%q] = %q\n", k, v)
// 		}

// 		envBackupData[k] = v
// 	}
// }

// func envReset() {
// 	if Debug <= DebugWarn {
// 		fmt.Fprintln(os.Stderr, "templar.envReset()")
// 	}
// 	os.Clearenv()
// 	for k, v := range envBackupData {
// 		if Debug <= DebugLog {
// 			fmt.Fprintf(os.Stderr, "templar.envReset() | os.Setenv(%q, %q)\n", k, v)
// 		}
// 		os.Setenv(k, v)
// 	}
// }

func init() {
	dataProvider = make(map[string]interface{})
	// envBackup()
}

// InitData initializes the template environment with external data
func InitData(checkDotEnv bool, files ...string) (err error) {
	if Debug <= DebugInfo {
		fmt.Fprintf(os.Stderr, "templar.InitData() | checkDotEnv = %t | initialized = %t\n", checkDotEnv, initialized)
	}
	if initialized == false {
		if checkDotEnv {
			gotenv.OverLoad()
		}
		parseEnvironment()
		initialized = true
	}

	for _, file := range files {
		parseFileData(file)
	}

	return err
}

func parseEnvironment() {
	if Debug <= DebugInfo {
		fmt.Fprintf(os.Stderr, "templar.parseEnvironment() | initialized = %t\n", initialized)
	}
	for _, base := range os.Environ() {
		pair := strings.SplitN(base, "=", 2)
		k := pair[0]
		v := pair[1]
		dataProvider[k] = v
	}
}

func parseFileData(file string) (err error) {
	if Debug <= DebugInfo {
		fmt.Fprintf(os.Stderr, "templar.parseFileData() | file = %q\n", file)
	}
	if len(file) > 0 {
		ext := path.Ext(file)
		switch strings.ToUpper(ext) {
		case ".ENV":
			if Debug <= DebugInfo {
				fmt.Fprintf(os.Stderr, "templar.parseFileData() | .ENV | gotenv.OverLoad(%q)\n", file)
			}
			err = gotenv.OverLoad(file)
			parseEnvironment()
		case ".INI":
			err = ParseINI(file)
			if err != nil {
				return err
			}
		case ".JSON":
			err = ParseJSON(file)
			if err != nil {
				return err
			}
		default:
			fmt.Errorf("Unknown data file type: %q", ext)
		}
	}

	return err
}

// ParseINI loads INI file data into the dataProvider
func ParseINI(file string) (err error) {
	var iniData *ini.File
	iniData, err = ini.Load(file)

	for _, section := range iniData.Sections() {
		for _, key := range section.Keys() {
			if section.Name() == ini.DEFAULT_SECTION {
				if _, ok := dataProvider[ini.DEFAULT_SECTION]; ok == false {
					dataProvider[ini.DEFAULT_SECTION] = make(map[string]string)
				}
				dataProvider[ini.DEFAULT_SECTION].((map[string]string))[key.Name()] = key.Value()
				dataProvider[key.Name()] = key.Value()
			} else {
				if _, ok := dataProvider[section.Name()]; ok == false {
					dataProvider[section.Name()] = make(map[string]string)
				}
				dataProvider[section.Name()].((map[string]string))[key.Name()] = key.Value()
			}
		}
	}

	return err
}

// ParseJSON loads JSON file data into the dataProvider
func ParseJSON(file string) (err error) {
	var jsonData []byte

	if len(file) > 0 {
		jsonData, err = ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		err = json.Unmarshal(jsonData, &dataProvider)
	}

	return err
}

// Reinitialize resets the dataProvider
func Reinitialize(debug int) {
	Debug = debug
	if Debug <= DebugInfo {
		fmt.Fprintf(os.Stderr, "templar.Reinitialize() | debug = %d | initialized = %t\n", debug, initialized)
	}
	dataProvider = make(map[string]interface{})
	gotenv.Reset()
	initialized = false
}

// Render handles template rendering
func Render(template string) (output string, err error) {
	output, err = mustache.RenderFile(template, dataProvider)
	return output, err
}

// RenderToFile handles rendering templates to file
func RenderToFile(filename, template string) (output string, err error) {
	output, err = mustache.RenderFile(template, dataProvider)
	if err != nil {
		return output, err
	}
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	return output, err
}
