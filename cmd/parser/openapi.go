// openapi.go
//
// Package parser defines the OpenAPI parser subcommand for the Kuma CLI.
// It processes OpenAPI specification files and integrates the parsed configuration
// into the Kuma CLI's shared configuration.
package parser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/arthurbcp/kuma-cli/cmd/shared"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

// OpenAPIParserCmd represents the 'openapi' parser subcommand.
var OpenAPIParserCmd = &cobra.Command{
	Use:   "openapi",
	Short: "OpenAPI file parser",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure that the file to parse is provided.
		if ParserFilePath == "" {
			fmt.Println("File to parse is required")
			return
		}

		// Parse the OpenAPI file and handle errors.
		config, err := parseOpenAPI(ParserFilePath)
		if err != nil {
			helpers.ErrorPrint("Parsing file error: " + err.Error())
			return
		}

		// Integrate the parsed configuration into the shared Kuma configuration.
		shared.KumaConfig = config
	},
}

// GetAvailableParsersString returns a formatted string of available parsers.
func GetAvailableParsersString() string {
	return strings.Join(AvailableParsers, "\n - ") + "\n"
}

// parseOpenAPI processes the OpenAPI specification file and returns the configuration map.
// It performs the following steps:
// 1. Reads the OpenAPI file content.
// 2. Unmarshals the JSON content into a generic map.
// 3. Converts the generic map into a structured OpenAPI template.
// 4. Marshals the structured template back into JSON.
// 5. Unmarshals the JSON into a configuration map.
func parseOpenAPI(file string) (map[string]interface{}, error) {
	helpers.HeaderPrint("Parsing OpenAPI file")

	// Read the content of the OpenAPI file.
	openAPIFile, err := helpers.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON content into a generic map.
	fileData := make(map[string]interface{})
	err = json.Unmarshal([]byte(openAPIFile), &fileData)
	if err != nil {
		return nil, err
	}

	// Parse the generic map into a structured OpenAPI template.
	fileStruct := openapi.ParseToOpenAPITemplate(fileData)

	// Marshal the structured template back into JSON.
	j, err := json.Marshal(fileStruct)
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
