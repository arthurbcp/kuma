package execHandlers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/style"
)

func HandleCommand(cmdStr string, vars map[string]interface{}) {
	var err error
	helpers := helpers.NewHelpers()

	cmdStr, err = helpers.ReplaceVars(cmdStr, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing command error: " + err.Error())
		os.Exit(1)
	}

	style.LogPrint(fmt.Sprintf("running: %s", cmdStr))

	cmdArgs := strings.Split(cmdStr, " ")
	execCmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	// Set the command's standard output to the console
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	// Execute the command
	err = execCmd.Run()
	if err != nil {
		style.ErrorPrint("command error: " + err.Error())
		os.Exit(1)
	}
}
