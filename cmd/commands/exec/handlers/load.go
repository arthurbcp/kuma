package execHandlers

import (
	"net/url"
	"os"
	"strings"

	execBuilders "github.com/arthurbcp/kuma/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/internal/helpers"
	"github.com/arthurbcp/kuma/pkg/filesystem"
	"github.com/arthurbcp/kuma/pkg/style"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/afero"
)

func HandleLoad(load map[string]interface{}, vars map[string]interface{}) {
	var err error
	data := vars["data"].(map[string]interface{})
	fs := filesystem.NewFileSystem(afero.NewOsFs())

	from, err := execBuilders.BuildStringValue("from", load, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	out, err := execBuilders.BuildStringValue("out", load, vars, true)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	var fileVars map[string]interface{}
	parsedURI, err := url.ParseRequestURI(from)
	if err != nil {
		fileVars, err = helpers.UnmarshalFile(from, fs)
		if err != nil {
			style.ErrorPrint("parsing file error: " + err.Error())
			os.Exit(1)
		}
	} else {
		err = spinner.New().
			Title("Downloading variables file").
			Action(func() {
				varsContent, err := fs.ReadFileFromURL(from)
				if err != nil {
					style.ErrorPrint("reading file error: " + err.Error())
					os.Exit(1)
				}
				splitURIPath := strings.Split(parsedURI.Path, "/")
				fileVars, err = helpers.UnmarshalByExt(splitURIPath[len(splitURIPath)-1], []byte(varsContent))
				if err != nil {
					style.ErrorPrint("parsing file error: " + err.Error())
					os.Exit(1)
				}
			}).
			Run()

		if err != nil {
			style.ErrorPrint("downloading variables file error: " + err.Error())
			os.Exit(1)
		}
	}
	data[out] = fileVars
}
