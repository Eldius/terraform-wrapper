package main

import (
  "os"

  "github.com/hashicorp/terraform/backend"
  "github.com/hashicorp/terraform/command"
  "github.com/hashicorp/terraform/tfdiags"
)

var code string

const stateFilename = "simple.tfstate"

/*
CommandMeta is a test 
*/
type CommandMeta struct {
	command.Meta
}

/*
func readTestTemplate(fileName string)  {
	content, err := ioutil.ReadFile("golangcode.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Convert []byte to string and print to screen
    text := string(content)
}
*/

/*
Execute executes
*/
func (c CommandMeta)Execute() {
	var be backend.Enhanced
	var beDiags tfdiags.Diagnostics
	configPath, err := command.ModulePath([]string{})
	if err != nil {
		c.Ui.Error(err.Error())
		os.Exit(1)
	}
	backendConfig, configDiags := c.loadBackendConfig(configPath)
	diags = diags.Append(configDiags)
	if configDiags.HasErrors() {
		c.showDiagnostics(diags)
		return 1
	}

	be, beDiags = c.Backend(&command.BackendOpts{
		Config: backendConfig,
	})
}

func main() {
	cmdMeta := CommandMeta{}
	cmdMeta.Execute()
}
