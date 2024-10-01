// parse.go
//
// Package parser defines the 'parse' subcommand for the Kuma CLI.
// It manages different parser helpers to process configuration files.
package parser

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// OpenAPIParser is a constant representing the OpenAPI V2.0 parser.
const OpenAPIParser2 = "openapi2"

// AvailableParsers lists all supported parser helpers.
var AvailableParsers = []string{
	OpenAPIParser2,
}

var (
	// ParserFilePath specifies the path to the file that needs to be parsed.
	ParserFilePath string
	// TargetDir specifies the target directory for the parsed file.
	ParsedFileTargetPath string
	// FileContent specifies the content after being parsed.
	ParsedFileContent string
)

// ParseCmd represents the 'parse' subcommand.
var ParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parser helpers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case OpenAPIParser2:
				// Execute the OpenAPI parser subcommand.
				OpenAPI2ParserCmd.Run(cmd, args)
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

// GetAvailableParsersString returns a formatted string of available parsers.
func GetAvailableParsersString() string {
	return strings.Join(AvailableParsers, "\n - ") + "\n"
}

// init initializes the 'parse' subcommand by adding parser-specific subcommands
// and setting up persistent flags.
func init() {
	// Add the OpenAPI parser as a subcommand.
	ParseCmd.AddCommand(OpenAPI2ParserCmd)

	// Define a persistent flag for specifying the file to parse.
	ParseCmd.PersistentFlags().StringVarP(&ParserFilePath, "file", "f", "", "Path to the file you want to parse")
	// Target file directory
	ParseCmd.PersistentFlags().StringVarP(&ParsedFileTargetPath, "out-dir", "o", "", "output directory for the parsed file")

	// Mark the 'file' flag as required.
	if err := ParseCmd.MarkPersistentFlagRequired("file"); err != nil {
		fmt.Println(err)
		return
	}
}
