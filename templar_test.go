package templar

import (
	"os"
	"testing"
)

var noDotEnvExpectation = `Hello Oedipus!
How do you like it in "/Users/runeimp/dev/apps/templar"?

  ENV_FILE_COMMENT == ''
      ENV_FILE_VAR == 'The Bard'
       CLI_ENV_VAR == 'Sound and fury'
           CLI_VAR == 'As you like it'
           boolean == 
     one.two.three == 
       numbers.two == 
               all ==
         words.all == 
             POSIX == 
DEFAULT.global_ini == 

`

var dotEnvExpectation = `Hello Horatio!
How do you like it in "/Users/runeimp/dev/apps/templar"?

  ENV_FILE_COMMENT == ''
      ENV_FILE_VAR == '.env Ninja!'
       CLI_ENV_VAR == 'Sound and fury'
           CLI_VAR == 'As you like it'
           boolean == 
     one.two.three == 
       numbers.two == 
               all ==
         words.all == 
             POSIX == 
DEFAULT.global_ini == 

`

var dotEnvAndINIExpectation = `Hello Hamlet!
How do you like it in "/Users/runeimp/dev/apps/templar"?

  ENV_FILE_COMMENT == ''
      ENV_FILE_VAR == '.env Ninja!'
       CLI_ENV_VAR == 'Sound and fury'
           CLI_VAR == 'As you like it'
           boolean == 
     one.two.three == 
       numbers.two == 2.1
               all ==
         words.all == your base are belong to us
             POSIX == Awesome!
DEFAULT.global_ini == true

`

var dotEnvAndJSONExpectation = `Hello Horatio!
How do you like it in "/Users/runeimp/dev/apps/templar"?

  ENV_FILE_COMMENT == ''
      ENV_FILE_VAR == '.env Ninja!'
       CLI_ENV_VAR == 'Sound and fury'
           CLI_VAR == 'As you like it'
           boolean == false
     one.two.three == 4
       numbers.two == 
               all == your base are belong to us
         words.all == 
             POSIX == 
DEFAULT.global_ini == 

`

func TestTemplar(t *testing.T) {
	tests := []struct {
		name        string
		checkDotEnv bool
		data        []string
		template    string
		want        string
	}{
		{name: ".env", template: "example.tmpl", checkDotEnv: true, want: dotEnvExpectation},
		{name: ".env and example.ini", template: "example.tmpl", checkDotEnv: true, data: []string{"example.ini"}, want: dotEnvAndINIExpectation},
		{name: ".env and example.json", template: "example.tmpl", checkDotEnv: true, data: []string{"example.json"}, want: dotEnvAndJSONExpectation},
		{name: "example.env", template: "example.tmpl", checkDotEnv: false, data: []string{"example.env"}, want: noDotEnvExpectation},
	}

	for _, tc := range tests {
		Reinitialize()
		os.Setenv("CLI_ENV_VAR", "Sound and fury")
		os.Setenv("CLI_VAR", "As you like it")
		if len(tc.data) == 0 {
			InitData(tc.checkDotEnv)
		} else {
			for _, file := range tc.data {
				InitData(tc.checkDotEnv, file)
			}
		}
		got, _ := Render(tc.template)
		if tc.want != got {
			t.Fatalf(`"%s": expected: %v, got: %v`, tc.name, tc.want, got)
		}
	}
}

// func TestDotEnv(t *testing.T) {
// 	Reinitialize()
// 	os.Setenv("CLI_ENV_VAR", "Sound and fury")
// 	os.Setenv("CLI_VAR", "As you like it")
// 	checkDotEnv := true
// 	want := dotEnvExpectation
// 	InitData(checkDotEnv)
// 	got, _ := Render("example.tmpl")
// 	if got != want {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }

// func TestDotEnvAndINI(t *testing.T) {
// 	Reinitialize()
// 	os.Setenv("CLI_ENV_VAR", "Sound and fury")
// 	os.Setenv("CLI_VAR", "As you like it")
// 	checkDotEnv := true
// 	want := dotEnvAndINIExpectation
// 	InitData(checkDotEnv, "example.ini")
// 	got, _ := Render("example.tmpl")
// 	if got != want {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }

// func TestDotEnvAndJSON(t *testing.T) {
// 	Reinitialize()
// 	os.Setenv("CLI_ENV_VAR", "Sound and fury")
// 	os.Setenv("CLI_VAR", "As you like it")
// 	checkDotEnv := true
// 	want := dotEnvAndJSONExpectation
// 	InitData(checkDotEnv, "example.json")
// 	got, _ := Render("example.tmpl")
// 	if got != want {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }

// func TestNoDotEnv(t *testing.T) {
// 	Reinitialize()
// 	os.Setenv("CLI_ENV_VAR", "Sound and fury")
// 	os.Setenv("CLI_VAR", "As you like it")
// 	checkDotEnv := false
// 	want := noDotEnvExpectation
// 	InitData(checkDotEnv, "example.env")
// 	got, _ := Render("example.tmpl")
// 	if got != want {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }
