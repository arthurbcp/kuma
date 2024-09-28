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

var OpenAPIParserCmd = &cobra.Command{
	Use:   "openapi",
	Short: "OpenAPI file parser",
	Run: func(cmd *cobra.Command, args []string) {
		if ParserFilePath == "" {
			fmt.Println("File to parse is required")
			return
		}
		config, err := parseOpenAPI(ParserFilePath)
		if err != nil {
			helpers.ErrorPrint("Parsing file error: " + err.Error())
			return
		}
		shared.KumaConfig = config
	},
}

func GetAvailableParsersString() string {
	return strings.Join(AvailableParsers, "\n - ") + "\n"
}

func parseOpenAPI(file string) (map[string]interface{}, error) {
	helpers.HeaderPrint("Parsing OpenAPI file")
	openAPIFile, err := helpers.ReadFile(file)
	if err != nil {
		return nil, err
	}
	fileData := make(map[string]interface{})
	err = json.Unmarshal([]byte(openAPIFile), &fileData)
	if err != nil {
		return nil, err
	}
	fileStruct := openapi.ParseToOpenAPITemplate(fileData)
	j, err := json.Marshal(fileStruct)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}
	err = json.Unmarshal(j, &config)
	if err != nil {
		return nil, err
	}
	helpers.CheckMarkPrint("OpenAPI file parsed successfully!")
	return config, nil
}

func init() {

}
