// generate.go
//
// Package generate defines the 'generate' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	// KumaConfigFilePath specifies the path to the Kuma configuration file.
	KumaConfigFilePath string

	// KumaTemplatesPath defines the path to the directory containing Kuma templates.
	KumaTemplatesPath string

	//KumaConfigParsedFileTargetPath specifies the target directory for the parsed file.
	KumaConfigParsedFileTargetPath string
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
			if !helpers.StringContains(parser.AvailableParsers, ParserToUse) {
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
	builder, err := domain.NewBuilder(KumaConfigFilePath, shared.KumaConfig, domain.NewConfig(ProjectPath, KumaTemplatesPath))
	if err != nil {
		helpers.ErrorPrint(err.Error())
		os.Exit(1)
	}

	// Execute the build process using the BuilderHandler.
	if err = handlers.NewBuilderHandler(builder).Build(map[string]interface{}{}); err != nil {
		helpers.ErrorPrint(err.Error())
		os.Exit(1)
	}

	if KumaConfigParsedFileTargetPath != "" {
		if err := helpers.CreateDirectoryIfNotExists(KumaConfigParsedFileTargetPath); err != nil {
			helpers.ErrorPrint("creating target directory error: " + err.Error())
			os.Exit(1)
		}
		fileName := "parsed-" +
			KumaConfigFilePath[strings.LastIndex(KumaConfigFilePath, "/")+1:]
		if err := helpers.WriteFile(filepath.Join(KumaConfigParsedFileTargetPath, fileName), builder.ParsedFile); err != nil {
			helpers.ErrorPrint("writing file error: " + err.Error())
			os.Exit(1)
		}
		helpers.CheckMarkPrint(fmt.Sprintf("File %s written successfully!", fileName))
	}
}

// init sets up flags for the 'generate' subcommand and binds them to variables.
func init() {
	// Parser-related flags.
	GenerateCmd.Flags().StringVarP(&ParserToUse, "parser", "", "", "Helper parser you want to use")
	GenerateCmd.Flags().StringVarP(&parser.ParserFilePath, "p-file", "", "", fmt.Sprintf("File path you want to parse\nAvailable parsers:\n - %s", parser.GetAvailableParsersString()))

	// Generate-related flags.
	// Target file directory
	GenerateCmd.Flags().StringVarP(&KumaConfigParsedFileTargetPath, "target-dir", "", "", "target directory for the kuma parsed config file")
	GenerateCmd.Flags().StringVarP(&KumaConfigFilePath, "config", "c", "kuma-config.yaml", "Path to the Kuma config file")
	GenerateCmd.Flags().StringVarP(&ProjectPath, "project-path", "p", "kuma-generated", "Path to the project you want to generate")
	GenerateCmd.Flags().StringVarP(&KumaTemplatesPath, "templates-path", "t", "kuma-templates", "Path to the Kuma templates")
}
