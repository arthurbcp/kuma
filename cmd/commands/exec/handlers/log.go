package execHandlers

import (
	"os"

	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/style"
)

func HandleLog(log string, vars map[string]interface{}) {
	var err error
	helpers := helpers.NewHelpers()

	log, err = helpers.ReplaceVars(log, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing log error: " + err.Error())
		os.Exit(1)
	}

	style.LogPrint(log)
}
