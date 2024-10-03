// generate.go
//
// Package generate defines the 'generate' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package generate

import (
	"os"

	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/internal/handlers"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	// ProjectPath defines the directory where the project will be generated.
	ProjectPath string

	// KumaConfigFilePath specifies the path to the Kuma configuration file.
	KumaConfigFilePath string

	// KumaTemplatesPath defines the path to the directory containing Kuma templates.
	KumaTemplatesPath string

	//VariableFilePath specifies the path to the variables file.
	VariableFilePath string
)

// GenerateCmd represents the 'generate' subcommand.
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a scaffold for a project based on Go Templates",
	Run: func(cmd *cobra.Command, args []string) {
		helpers := helpers.NewHelpers()
		if VariableFilePath != "" {
			vars, err := helpers.UnmarshalFile(VariableFilePath)
			if err != nil {
				helpers.ErrorPrint("parsing file error: " + err.Error())
				os.Exit(1)
			}
			shared.TemplateVariables = vars
			build()
		}
	},
}

// build initializes the Builder and triggers the build process.
// It reads the Kuma configuration file and applies templates to generate the project structure.
func build() {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	helpers := helpers.NewHelpers()
	// Initialize a new Builder with the provided configurations.
	builder, err := domain.NewBuilder(fs, helpers, KumaConfigFilePath, shared.TemplateVariables, domain.NewConfig(ProjectPath, KumaTemplatesPath))
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

// init sets up flags for the 'generate' subcommand and binds them to variables.
func init() {
	// Target file directory
	GenerateCmd.Flags().StringVarP(&VariableFilePath, "variables-file", "v", "", "path to the variables file")
	GenerateCmd.Flags().StringVarP(&KumaConfigFilePath, "config", "c", "kuma-config.yaml", "Path to the Kuma config file")
	GenerateCmd.Flags().StringVarP(&ProjectPath, "project-path", "p", "kuma-generated", "Path to the project you want to generate")
	GenerateCmd.Flags().StringVarP(&KumaTemplatesPath, "templates-path", "t", "kuma-templates", "Path to the Kuma templates")
}
