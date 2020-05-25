package templar

import (
  "fmt"
  "os"
  "testing"
)

var exampleDotEnvExpectation = `Hello Oedipus!
How do you like it in "/Users/runeimp/Dropbox/Profile/Home/dev/apps/templar"?

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
How do you like it in "/Users/runeimp/Dropbox/Profile/Home/dev/apps/templar"?

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
How do you like it in "/Users/runeimp/Dropbox/Profile/Home/dev/apps/templar"?

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
How do you like it in "/Users/runeimp/Dropbox/Profile/Home/dev/apps/templar"?

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

var noDotEnvExpectation = `Hello runeimp!
How do you like it in "/Users/runeimp/Dropbox/Profile/Home/dev/apps/templar"?

  ENV_FILE_COMMENT == ''
      ENV_FILE_VAR == ''
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
    {name: "example.env", template: "example.tmpl", checkDotEnv: false, data: []string{"example.env"}, want: exampleDotEnvExpectation},
    {name: "no.env", template: "example.tmpl", checkDotEnv: false, want: noDotEnvExpectation},
  }

  for _, tc := range tests {
    debug := DebugWarn
    Reinitialize(debug)
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
      t.Fatalf(fmt.Sprintf("%q:\n\texpected: %v\n\tgot: %v\n\t | tc.checkDotEnv = %t\n", tc.name, tc.want, got, tc.checkDotEnv))
    }
  }
}
