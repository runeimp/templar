package templar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/subosito/gotenv"
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
		if checkDotEnv {
			gotenv.OverLoad()
		}

		// for _, base := range os.Environ() {
		// 	pair := strings.SplitN(base, "=", 2)
		// 	k := pair[0]
		// 	v := pair[1]
		// 	// fmt.Printf("templar.InitData() | %s = %s\n", k, v)
		// 	dataProvider[k] = v
		// }
		parseEnvironment()
		initialized = true
	}

	for _, file := range files {
		parseFileData(file)
	}
	// parseFileData(files)

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
		fmt.Printf("templar.parseFileData() | ext = %q\n", ext)
		switch strings.ToUpper(ext) {
		case ".ENV":
			fmt.Printf("templar.parseFileData() | Adding ENV data from %q\n", file)
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
func Test() {
	tmpl, _ := mustache.ParseString("Hello, {{c}}!\n")
	var buf bytes.Buffer
	for i := 0; i < 10; i++ {
		tmpl.FRender(&buf, map[string]string{"c": "world"})
	}
	fmt.Println(buf.String())
}
