package execHandlers

import (
	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/arthurbcp/kuma/v2/cmd/shared"
	"github.com/arthurbcp/kuma/v2/internal/domain"
	"github.com/arthurbcp/kuma/v2/internal/handlers"
	"github.com/arthurbcp/kuma/v2/pkg/filesystem"
	"github.com/spf13/afero"
)

func HandleCreate(module string, data map[string]interface{}, vars map[string]interface{}) error {
	path := shared.KumaFilesPath
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	if module != "" {
		path = shared.KumaFilesPath + "/" + module + "/" + shared.KumaFilesPath
	}
	builder, err := domain.NewBuilder(fs, domain.NewConfig(".", path))
	if err != nil {
		return err
	}
	from, err := execBuilders.BuildStringValue("from", data, vars, true, constants.CreateHandler)
	if err != nil {
		return err
	}
	err = builder.SetBuilderDataFromFile(path+"/"+from, vars)
	if err != nil {
		return err
	}

	if err = handlers.NewBuilderHandler(builder).Build(); err != nil {
		return err
	}
	return nil
}
