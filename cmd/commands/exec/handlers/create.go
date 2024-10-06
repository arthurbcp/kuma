package execHandlers

import (
	"os"

	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/internal/handlers"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	"github.com/spf13/afero"
)

func HandleCreate(data map[string]interface{}, vars map[string]interface{}) {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	helpers := helpers.NewHelpers()
	// Initialize a new Builder with the provided configurations.
	builder, err := domain.NewBuilder(fs, helpers, domain.NewConfig(".", shared.KumaTemplatesPath))
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	from, ok := data["from"].(string)
	if !ok {
		style.ErrorPrint("from is required")
		os.Exit(1)
	}
	from, err = helpers.ReplaceVars(from, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing from error: " + err.Error())
		os.Exit(1)
	}
	err = builder.SetBuilderDataFromFile(from, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	// Execute the build process using the BuilderHandler.
	if err = handlers.NewBuilderHandler(builder).Build(); err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
}
