package execHandlers

import (
	"os"
	"strings"

	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/internal/helpers"
	"github.com/arthurbcp/kuma/v2/pkg/filesystem"
	"github.com/arthurbcp/kuma/v2/pkg/style"
	"github.com/spf13/afero"
)

func HandleModify(data map[string]interface{}, vars map[string]interface{}) {

	fs := filesystem.NewFileSystem(afero.NewOsFs())

	file, err := execBuilders.BuildStringValue("file", data, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	fileContent, err := fs.ReadFile(file)
	if err != nil {
		style.ErrorPrint("reading file error: " + err.Error())
		os.Exit(1)
	}
	template, err := execBuilders.BuildStringValue("template", data, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	codeMark, err := execBuilders.BuildStringValue("mark", data, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	replace, err := execBuilders.BuildBoolValue("replace", data, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	templateContent, err := fs.ReadFile(template)
	if err != nil {
		style.ErrorPrint("reading template file error: " + err.Error())
		os.Exit(1)
	}
	templateContent, err = helpers.ReplaceVars(templateContent, map[string]interface{}{"data": data}, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing template file error: " + err.Error())
		os.Exit(1)
	}
	if replace {
		fileContent = strings.ReplaceAll(fileContent, codeMark, templateContent)
	} else {
		fileContent = strings.ReplaceAll(fileContent, codeMark, templateContent+"\n"+codeMark)
	}
	err = fs.WriteFile(file, fileContent)
	if err != nil {
		style.ErrorPrint("writing file error: " + err.Error())
		os.Exit(1)
	}
}
