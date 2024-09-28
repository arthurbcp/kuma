// parse.go
//
// Package parser defines the 'parse' subcommand for the Kuma CLI.
// It manages different parser helpers to process configuration files.
package parser

import (
	"fmt"

	"github.com/spf13/cobra"
)

// OpenAPIParser is a constant representing the OpenAPI parser.
const OpenAPIParser = "openapi"

// AvailableParsers lists all supported parser helpers.
var AvailableParsers = []string{
	OpenAPIParser,
}

var (
	// ParserFilePath specifies the path to the file that needs to be parsed.
	ParserFilePath string
)

// ParseCmd represents the 'parse' subcommand.
var ParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parser helpers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case OpenAPIParser:
				// Execute the OpenAPI parser subcommand.
				OpenAPIParserCmd.Run(cmd, args)
			default:
				// Notify the user if the specified parser is not available.
				fmt.Printf("Parser %s not found!\nAvailable parsers:\n - %s", args[0], GetAvailableParsersString())
			}
		} else {
			// Display help information if no parser is specified.
			cmd.Help()
		}
	},
}

// init initializes the 'parse' subcommand by adding parser-specific subcommands
// and setting up persistent flags.
func init() {
	// Add the OpenAPI parser as a subcommand.
	ParseCmd.AddCommand(OpenAPIParserCmd)

	// Define a persistent flag for specifying the file to parse.
	ParseCmd.PersistentFlags().StringVarP(&ParserFilePath, "file", "f", "", "Path to the file you want to parse")

	// Mark the 'file' flag as required.
	if err := ParseCmd.MarkPersistentFlagRequired("file"); err != nil {
		fmt.Println(err)
		return
	}
}
