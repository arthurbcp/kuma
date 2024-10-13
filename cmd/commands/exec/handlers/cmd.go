package execHandlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/arthurbcp/kuma/cmd/shared"
	"github.com/arthurbcp/kuma/internal/helpers"
	"github.com/arthurbcp/kuma/pkg/style"
)

func HandleCommand(cmdStr string, vars map[string]interface{}) {
	var err error

	cmdStr, err = helpers.ReplaceVars(cmdStr, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing command error: " + err.Error())
		os.Exit(1)
	}

	style.LogPrint(fmt.Sprintf("running: %s", cmdStr))

	cmdArgs := strings.Split(cmdStr, " ")
	if err := shared.RunCommand(cmdArgs[0], cmdArgs[1:]...); err != nil {
		style.ErrorPrint("command error: " + err.Error())
		os.Exit(1)
	}
}
