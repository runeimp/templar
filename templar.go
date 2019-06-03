package templar

import (
	"bytes"
	"fmt"
	"os"

	"github.com/cbroglie/mustache"
)

const (
	// Name denotes the library name
	Name = "Templar"

	// Version denotes the library version
	Version = "0.1.0"
)

var args struct {
	IniFile  string
	JSONFile string
	Template string
}

func dataMapper(key string) string {
	var result string
	switch key {
	case "TEMPLAR_VERSION":
		result = Version
	default:
		if len(args.IniFile) > 0 {
			// Do INI file data check
		}
		if len(args.JSONFile) > 0 {
			// Do JSON file data check
		}
		result = os.Getenv(key)
	}
	return result
}

// parseTemplate expands variables in the template using dataMapper
func parseTemplate() {
	os.Expand(args.Template, dataMapper)
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
