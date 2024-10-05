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
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	// ProjectPath defines the directory where the project will be created.
	ProjectPath string

	//VariablesFile specifies the path to the variables file.
	VariablesFile string
)

// CreateCmd represents the 'create' subcommand.
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a scaffold for a project based on Go Templates",
	Run: func(cmd *cobra.Command, args []string) {
		create()
	},
}

func create() {
	helpers := helpers.NewHelpers()
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	if VariablesFile != "" {
		var vars interface{}
		_, err := url.ParseRequestURI(VariablesFile)
		if err != nil {
			vars, err = helpers.UnmarshalFile(VariablesFile, fs)
			if err != nil {
				helpers.ErrorPrint("parsing file error: " + err.Error())
				os.Exit(1)
			}
		} else {
			helpers.TitlePrint("downloading variables file")
			varsContent, err := readFileFromURL(VariablesFile)
			if err != nil {
				helpers.ErrorPrint("reading file error: " + err.Error())
				os.Exit(1)
			}
			splitURL := strings.Split(VariablesFile, "/")
			vars, err = helpers.UnmarshalByExt(splitURL[len(splitURL)-1], []byte(varsContent))
			if err != nil {
				helpers.ErrorPrint("parsing file error: " + err.Error())
				os.Exit(1)
			}
		}
		shared.TemplateVariables = vars.(map[string]interface{})
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
	helpers := helpers.NewHelpers()
	// Initialize a new Builder with the provided configurations.
	builder, err := domain.NewBuilder(fs, helpers, shared.KumaConfigFilePath, shared.TemplateVariables, domain.NewConfig(ProjectPath, shared.KumaTemplatesPath))
	if err != nil {
		helpers.ErrorPrint(err.Error())
		os.Exit(1)
	}

	// Execute the build process using the BuilderHandler.
	if err = handlers.NewBuilderHandler(builder).Build(); err != nil {
		helpers.ErrorPrint(err.Error())
		os.Exit(1)
	}
}

// init sets up flags for the 'create' subcommand and binds them to variables.
func init() {
	// Target file directory
	CreateCmd.Flags().StringVarP(&VariablesFile, "variables-file", "v", "", "path or URL to the variables file")
	CreateCmd.Flags().StringVarP(&ProjectPath, "project-path", "p", ".", "Path to the project you want to create")
}
