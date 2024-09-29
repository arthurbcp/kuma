// generate.go
//
// Package generate defines the 'generate' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package generate

import (
	"fmt"
	"os"

	"github.com/arthurbcp/kuma-cli/cmd/parser"
	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/internal/handlers"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/spf13/cobra"
)

var (
	// ParserToUse specifies which parser helper to utilize.
	ParserToUse string

	// ProjectPath defines the directory where the project will be generated.
	ProjectPath string

	// KumaConfigFile specifies the path to the Kuma configuration file.
	KumaConfigFile string

	// KumaTemplates defines the path to the directory containing Kuma templates.
	KumaTemplates string
)

// GenerateCmd represents the 'generate' subcommand.
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a scaffold for a project based on Go Templates",
	Run: func(cmd *cobra.Command, args []string) {
		// If a parser is specified, validate and execute parsing before building.
		if ParserToUse != "" {
			// Ensure that a parser file path is provided.
			if parser.ParserFilePath == "" {
				fmt.Println("File to parse is required")
				return
			}

			// Check if the specified parser is available.
			if !helpers.Contains(parser.AvailableParsers, ParserToUse) {
				fmt.Printf("Parser %s not found!\nAvailable parsers:\n - %s",
					ParserToUse, parser.GetAvailableParsersString())
				return
			}

			// Execute the specified parser command.
			parser.ParseCmd.Run(cmd, []string{
				ParserToUse,
				"--file",
				parser.ParserFilePath,
			})

			// Proceed to build the project after parsing.
			build()
		}
	},
}

// build initializes the Builder and triggers the build process.
// It reads the Kuma configuration file and applies templates to generate the project structure.
func build() {
	// Initialize a new Builder with the provided configurations.
	builder, err := domain.NewBuilder(KumaConfigFile, shared.KumaConfig, domain.NewConfig(ProjectPath, KumaTemplates))
	if err != nil {
		helpers.ErrorPrint(err.Error())
		os.Exit(1)
	}

	// Execute the build process using the BuilderHandler.
	if err = handlers.NewBuilderHandler(builder).Build(map[string]interface{}{}); err != nil {
		helpers.ErrorPrint(err.Error())
		os.Exit(1)
	}
}

// init sets up flags for the 'generate' subcommand and binds them to variables.
func init() {
	// Parser-related flags.
	GenerateCmd.Flags().StringVarP(&ParserToUse, "parser", "", "", "Helper parser you want to use")
	GenerateCmd.Flags().StringVarP(&parser.ParserFilePath, "p-file", "", "", fmt.Sprintf("File path you want to parse\nAvailable parsers:\n - %s", parser.GetAvailableParsersString()))

	// Generate-related flags.
	GenerateCmd.Flags().StringVarP(&KumaConfigFile, "config", "c", "kuma-config.yaml", "Path to the Kuma config file")
	GenerateCmd.Flags().StringVarP(&ProjectPath, "project-path", "p", "kuma-generated", "Path to the project you want to generate")
	GenerateCmd.Flags().StringVarP(&KumaTemplates, "templates-path", "t", "kuma-templates", "Path to the Kuma templates")
}
