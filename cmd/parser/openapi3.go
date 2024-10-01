// openapi3.go
//
// Package parser defines the OpenAPI parser subcommand for the Kuma CLI.
// It processes OpenAPI specification files and integrates the parsed configuration
// into the Kuma CLI's shared configuration.
package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	openapi3 "github.com/arthurbcp/kuma-cli/pkg/openapi/v3"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// OpenAPI3ParserCmd represents the 'openapi' parser subcommand.
var OpenAPI3ParserCmd = &cobra.Command{
	Use:   "openapi3",
	Short: "OpenAPI file parser",
	Run: func(cmd *cobra.Command, args []string) {
		helpers := helpers.NewHelpers()
		fs := filesystem.NewFileSystem(afero.NewOsFs())
		// Ensure that the file to parse is provided.
		if ParserFilePath == "" {
			fmt.Println("File to parse is required")
			os.Exit(1)
		}

		// Parse the OpenAPI file and handle errors.
		config, err := parseOpenAPI3(ParserFilePath)
		if err != nil {
			helpers.ErrorPrint("parsing file error: " + err.Error())
			os.Exit(1)
		}

		// Integrate the parsed configuration into the shared Kuma configuration.
		shared.TemplateVariables = config
		if ParsedFileTargetPath != "" {
			if err := fs.CreateDirectoryIfNotExists(ParsedFileTargetPath); err != nil {
				helpers.ErrorPrint("creating target directory error: " + err.Error())
				os.Exit(1)
			}
			fileName := "kuma-" + strings.Replace(
				ParserFilePath[strings.LastIndex(ParserFilePath, "/")+1:],
				filepath.Ext(ParserFilePath), ".json", 1)
			ParsedFileContent = helpers.PrettyJson(ParsedFileContent)
			if err := fs.WriteFile(filepath.Join(ParsedFileTargetPath, fileName), ParsedFileContent); err != nil {
				helpers.ErrorPrint("writing file error: " + err.Error())
				os.Exit(1)
			}
			helpers.CheckMarkPrint(fmt.Sprintf("File %s written successfully!", fileName))
		}
	},
}

func parseOpenAPI3(file string) (map[string]interface{}, error) {
	helpers := helpers.NewHelpers()
	fileData, err := shared.UnmarshalFile(file)
	if err != nil {
		return nil, err
	}

	// Parse the generic map into a structured OpenAPI template.
	fileStruct := openapi3.ParseToOpenAPITemplate(helpers, fileData)

	// Marshal the structured template back into JSON.
	j, err := json.Marshal(fileStruct)
	ParsedFileContent = string(j)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into a configuration map.
	var config map[string]interface{}
	err = json.Unmarshal(j, &config)
	if err != nil {
		return nil, err
	}

	// Indicate successful parsing to the user.
	helpers.CheckMarkPrint("OpenAPI file parsed successfully!")
	return config, nil
}
