package templar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/subosito/gotenv"
	ini "gopkg.in/ini.v1"
)

/*
 * CONSTANTS
 */
const (
	// Name denotes the library name
	Name = "Templar"

	// Version denotes the library version
	Version = "0.1.1"
)

/*
 * VARIABLES
 */
var (
	dataProvider map[string]interface{}
	initialized  = false
)

/*
 * TYPES
 */
var Data struct {
	INIFile  []string
	JSONFile []string
	Template string
}

/*
 * METHODS
 */
func init() {
	dataProvider = make(map[string]interface{})
}

func InitData(checkDotEnv bool, files ...string) (err error) {
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
	for _, base := range os.Environ() {
		pair := strings.SplitN(base, "=", 2)
		k := pair[0]
		v := pair[1]
		dataProvider[k] = v
	}
}

func parseFileData(file string) (err error) {
	if len(file) > 0 {
		ext := path.Ext(file)
		switch strings.ToUpper(ext) {
		case ".ENV":
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
			fmt.Errorf("Unknown data file type: %q\n", ext)
		}
	}

	return err
}

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

func Reinitialize() {
	dataProvider = make(map[string]interface{})
	initialized = false
}

func Render(template string) (output string, err error) {
	output, err = mustache.RenderFile(template, dataProvider)
	return output, err
}

func RenderToFile(filename, template string) (output string, err error) {
	output, err = mustache.RenderFile(template, dataProvider)
	if err != nil {
		return output, err
	}
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	return output, err
}
