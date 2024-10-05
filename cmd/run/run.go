// run.go
//
// Package run defines the 'run' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package run

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/gookit/color"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	Run       string
	Variables map[interface{}]interface{}
)

// RunCmd represents the 'run' subcommand.
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a specific pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		ExecRun(Run)
	},
}

func ExecRun(name string) {
	helpers := helpers.NewHelpers()
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	data, err := helpers.UnmarshalFile(shared.KumaRunsPath, fs)
	if err != nil {
		helpers.ErrorPrint("parsing file error: " + err.Error())
		os.Exit(0)
	}
	run, ok := data[name]
	if !ok {
		helpers.ErrorPrint("run not found: " + name)
		os.Exit(0)
	}
	for _, step := range run.([]interface{}) {
		step := step.(map[interface{}]interface{})
		for key, value := range step {
			color.Gray.Printf("executing: %s %s\n", key, value)
			if key == "cmd" {
				handleCommand(value.(string))
			} else if key == "input" {
				handleInput(value.(map[interface{}]interface{}))
			}
		}
	}
}

func handleInput(input map[interface{}]interface{}) {
	var helpers = helpers.NewHelpers()
	msg, ok := input["msg"].(string)
	if !ok {
		helpers.ErrorPrint("msg is required for input")
		os.Exit(0)
	}
	out, ok := input["out"].(string)
	if !ok {
		helpers.ErrorPrint("out is required for input")
		os.Exit(0)
	}

	// Create a reader to read input from stdin (standard input)
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for input
	fmt.Print(msg)

	// Read the full line of input (until the user presses enter)
	outValue, _ := reader.ReadString('\n')

	Variables[out] = strings.TrimSpace(outValue)
	fmt.Println(outValue)
}

func handleCommand(cmdStr string) {
	helpers := helpers.NewHelpers()
	cmdArgs := strings.Split(cmdStr, " ")
	execCmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	// Set the command's standard output to the console
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	// Execute the command
	err := execCmd.Run()
	if err != nil {
		helpers.ErrorPrint("command error: " + err.Error())
		os.Exit(0)
	}
}

// init sets up flags for the 'run' subcommand and binds them to variables.
func init() {
	// Repository name
	RunCmd.Flags().StringVarP(&Run, "name", "n", "", "run to use")
	RunCmd.MarkFlagRequired("name")
}
