package parser

import (
	"fmt"

	"github.com/spf13/cobra"
)

const OpenAPIParser = "openapi"

var AvailableParsers = []string{
	OpenAPIParser,
}

var (
	ParserFilePath string
)

var ParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parser helpers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if args[0] == OpenAPIParser {
				OpenAPIParserCmd.Run(cmd, args)
				return
			}
			fmt.Printf("Parser %s not found!\nAvailable parsers:\n - %s", args[0], GetAvailableParsersString())
		}
	},
}

func init() {
	ParseCmd.AddCommand(OpenAPIParserCmd)
	ParseCmd.PersistentFlags().StringVarP(&ParserFilePath, "file", "f", "", "Path to the file you want to parse")
	if err := ParseCmd.MarkPersistentFlagRequired("file"); err != nil {
		fmt.Println(err)
		return
	}
}
