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

const (
	// Name denotes the library name
	Name = "Templar"

	// Version denotes the library version
	Version = "0.1.0"
)

var (
	dataProvider map[string]interface{}
	initialized  = false
)

var Data struct {
	INIFile  []string
	JSONFile []string
	Template string
}

func init() {
	dataProvider = make(map[string]interface{})

}

func InitData(checkDotEnv bool, files ...string) (err error) {
	if initialized == false {
		// fmt.Printf("templar.InitData() | checkDotEnv = %t\n", checkDotEnv)
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
		// fmt.Printf("templar.InitData() | %s = %s\n", k, v)
		dataProvider[k] = v
	}
}

func parseFileData(file string) (err error) {
	if len(file) > 0 {
		ext := path.Ext(file)
		// fmt.Printf("templar.parseFileData() | ext = %q\n", ext)
		switch strings.ToUpper(ext) {
		case ".ENV":
			// fmt.Printf("templar.parseFileData() | Adding ENV data from %q\n", file)
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
		// fmt.Printf("templar.ParseINI() | section.Name() = %q | section.KeyStrings() = %q\n", section.Name(), section.KeyStrings())
		for _, key := range section.Keys() {
			// fmt.Printf("templar.ParseINI() | %s.%s = %v (%T / %T)\n", section.Name(), key.Name(), key.Value(), key.Value(), key.String())
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

	// fmt.Printf("templar.ParseINI() |           Keys = %#v\n", iniData.Section("").Keys())
	// fmt.Printf("templar.ParseINI() |     KeyStrings = %q\n", iniData.Section("").KeyStrings())

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
	// fmt.Printf("templar.ParseJSON() |     jsonData = %s\n", string(jsonData))
	// fmt.Printf("templar.ParseJSON() | dataProvider = %#v\n", dataProvider)

	return err
}

func Render(template string) (output string, err error) {
	output, err = mustache.RenderFile(template, dataProvider)
	return output, err
}

// Test Mustache template system
// func Test() {
// 	tmpl, _ := mustache.ParseString("Hello, {{c}}!\n")
// 	var buf bytes.Buffer
// 	for i := 0; i < 10; i++ {
// 		tmpl.FRender(&buf, map[string]string{"c": "world"})
// 	}
// 	fmt.Println(buf.String())
// }
