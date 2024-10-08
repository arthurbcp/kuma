// create.go
//
// Package create defines the 'create' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package create

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/internal/handlers"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	// ProjectPath defines the directory where the project will be created.
	ProjectPath string

	//VariablesFile specifies the path to the variables file.
	VariablesFile string

	//FromFile specifies the path to the YAML file with the structure and templates.
	FromFile string

	// TemplateVariables holds the variables for template replacement during the generate process.
	TemplateVariables map[string]interface{}
)

// CreateCmd represents the 'create' subcommand.
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a scaffold for a project based on Go Templates",
	Run: func(cmd *cobra.Command, args []string) {
		Create()
	},
}

func Create() {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	if VariablesFile != "" {
		var vars interface{}
		_, err := url.ParseRequestURI(VariablesFile)
		if err != nil {
			vars, err = helpers.UnmarshalFile(VariablesFile, fs)
			if err != nil {
				style.ErrorPrint("parsing file error: " + err.Error())
				os.Exit(1)
			}
		} else {
			style.LogPrint("downloading variables file")
			varsContent, err := readFileFromURL(VariablesFile)
			if err != nil {
				style.ErrorPrint("reading file error: " + err.Error())
				os.Exit(1)
			}
			splitURL := strings.Split(VariablesFile, "/")
			vars, err = helpers.UnmarshalByExt(splitURL[len(splitURL)-1], []byte(varsContent))
			if err != nil {
				style.ErrorPrint("parsing file error: " + err.Error())
				os.Exit(1)
			}
		}
		TemplateVariables = vars.(map[string]interface{})
		build()
	}
}

func readFileFromURL(url string) (string, error) {
	// Send the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if request succeeded
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Read the body into a byte slice
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	return string(bodyBytes), nil
}

// build initializes the Builder and triggers the build process.
// It reads the Kuma configuration file and applies templates to create the project structure.
func build() {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	// Initialize a new Builder with the provided configurations.
	builder, err := domain.NewBuilder(fs, domain.NewConfig(ProjectPath, shared.KumaFilesPath))
	builder.SetBuilderDataFromFile(shared.KumaFilesPath+"/"+FromFile, TemplateVariables)
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

// init sets up flags for the 'create' subcommand and binds them to variables.
func init() {
	// Target file directory
	CreateCmd.Flags().StringVarP(&VariablesFile, "variables", "v", "", "path or URL to the variables file")
	CreateCmd.Flags().StringVarP(&ProjectPath, "project", "p", ".", "Path to the project you want to create")
	CreateCmd.Flags().StringVarP(&FromFile, "from", "f", ".", "Path to the YAML file with the structure and templates")
}
