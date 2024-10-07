package execHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma-cli/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/internal/handlers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	"github.com/spf13/afero"
)

func HandleCreate(data map[string]interface{}, vars map[string]interface{}) {
	fs := filesystem.NewFileSystem(afero.NewOsFs())

	// Initialize a new Builder with the provided configurations.
	builder, err := domain.NewBuilder(fs, domain.NewConfig(".", shared.KumaTemplatesPath))
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	from, err := execBuilders.BuildStringValue("from", data, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
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
