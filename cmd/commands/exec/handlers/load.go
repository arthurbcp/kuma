package execHandlers

import (
	"net/url"
	"os"
	"strings"

	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	"github.com/spf13/afero"
)

func HandleLoad(load map[string]interface{}, vars map[string]interface{}) {
	var err error
	data := vars["data"].(map[string]interface{})
	helpers := helpers.NewHelpers()
	fs := filesystem.NewFileSystem(afero.NewOsFs())

	from, ok := load["from"].(string)
	if !ok {
		style.ErrorPrint("from is required")
		os.Exit(1)
	}
	from, err = helpers.ReplaceVars(from, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing from error: " + err.Error())
		os.Exit(1)
	}

	out, ok := load["out"].(string)
	if !ok {
		style.ErrorPrint("out is required")
		os.Exit(1)
	}
	out, err = helpers.ReplaceVars(out, vars, helpers.GetFuncMap())
	if err != nil {
		style.ErrorPrint("parsing from error: " + err.Error())
		os.Exit(1)
	}

	var fileVars map[string]interface{}
	_, err = url.ParseRequestURI(from)
	if err != nil {
		fileVars, err = helpers.UnmarshalFile(from, fs)
		if err != nil {
			style.ErrorPrint("parsing file error: " + err.Error())
			os.Exit(1)
		}
	} else {
		style.LogPrint("downloading variables file")
		varsContent, err := fs.ReadFileFromURL(from)
		if err != nil {
			style.ErrorPrint("reading file error: " + err.Error())
			os.Exit(1)
		}
		splitURL := strings.Split(from, "/")
		fileVars, err = helpers.UnmarshalByExt(splitURL[len(splitURL)-1], []byte(varsContent))
		if err != nil {
			style.ErrorPrint("parsing file error: " + err.Error())
			os.Exit(1)
		}
	}
	data[out] = fileVars
}
