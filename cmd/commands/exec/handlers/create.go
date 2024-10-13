package execHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/cmd/shared"
	"github.com/arthurbcp/kuma/internal/domain"
	"github.com/arthurbcp/kuma/internal/handlers"
	"github.com/arthurbcp/kuma/pkg/filesystem"
	"github.com/arthurbcp/kuma/pkg/style"
	"github.com/spf13/afero"
)

func HandleCreate(module string, data map[string]interface{}, vars map[string]interface{}) {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	modulePath := shared.KumaFilesPath + "/" + module + "/" + shared.KumaFilesPath
	// Initialize a new Builder with the provided configurations.
	builder, err := domain.NewBuilder(fs, domain.NewConfig(".", modulePath))
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	from, err := execBuilders.BuildStringValue("from", data, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	err = builder.SetBuilderDataFromFile(modulePath+"/"+from, vars)
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
