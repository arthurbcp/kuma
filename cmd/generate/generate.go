package generate

import (
	"fmt"

	"github.com/arthurbcp/kuma-cli/cmd/parser"
	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/internal/handlers"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/spf13/cobra"
)

var (
	ParserToUse    string
	ProjectPath    string
	KumaConfigFile string
	KumaTemplates  string
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a scaffold for a project based on Go Templates",
	Run: func(cmd *cobra.Command, args []string) {
		if ParserToUse != "" {
			if parser.ParserFilePath == "" {
				fmt.Println("File to parse is required")
				return
			}
			if !helpers.Contains(parser.AvailableParsers, ParserToUse) {
				fmt.Printf("Parser %s not found!\nAvailable parsers:\n - %s",
					ParserToUse, parser.GetAvailableParsersString())
				return
			}
			parser.ParseCmd.Run(cmd, []string{
				ParserToUse,
				"--file",
				parser.ParserFilePath,
			})
			build()
		}
	},
}

func build() {
	builder, err := domain.NewBuilder(KumaConfigFile, shared.KumaConfig, domain.NewConfig(ProjectPath, KumaTemplates))
	if err != nil {
		helpers.ErrorPrint(err.Error())
		return
	}
	if err = handlers.NewBuilderHandler(builder).Build(map[string]interface{}{}); err != nil {
		helpers.ErrorPrint(err.Error())
		return
	}
}

func init() {
	// Parser flags
	GenerateCmd.Flags().StringVarP(&ParserToUse, "parser", "", "", "Helper parse you want use")
	GenerateCmd.Flags().StringVarP(&parser.ParserFilePath, "p-file", "", "", fmt.Sprintf("File path you want to parse\nAvailable parsers:\n - %s", parser.GetAvailableParsersString()))

	// Generate flags
	GenerateCmd.Flags().StringVarP(&KumaConfigFile, "config", "c", "kuma-config.yaml", "Path to the Kuma config file")
	GenerateCmd.Flags().StringVarP(&ProjectPath, "project-path", "p", "kuma-generated", "Path to the project you want to generate")
	GenerateCmd.Flags().StringVarP(&KumaTemplates, "templates-path", "t", "kuma-templates", "Path to the Kuma templates")
}
