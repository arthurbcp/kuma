package execHandlers

import (
	"os"

	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/arthurbcp/kuma/v2/cmd/shared"
	"github.com/arthurbcp/kuma/v2/internal/domain"
	"github.com/arthurbcp/kuma/v2/internal/handlers"
	"github.com/arthurbcp/kuma/v2/pkg/filesystem"
	"github.com/arthurbcp/kuma/v2/pkg/style"
	"github.com/spf13/afero"
)

func HandleCreate(module string, data map[string]interface{}, vars map[string]interface{}) {
	path := shared.KumaFilesPath
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	if module != "" {
		path = shared.KumaFilesPath + "/" + module + "/" + shared.KumaFilesPath
	}
	builder, err := domain.NewBuilder(fs, domain.NewConfig(".", path))
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	from, err := execBuilders.BuildStringValue("from", data, vars, true, constants.CreateHandler)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	err = builder.SetBuilderDataFromFile(path+"/"+from, vars)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	if err = handlers.NewBuilderHandler(builder).Build(); err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
}
