// root.go
//
// Package cmd serves as the entry point for the Kuma CLI application.
// It defines the root command and integrates all subcommands into the CLI.
package cmd

import (
	"fmt"
	"os"

	"github.com/arthurbcp/kuma/cmd/commands/create"
	execRun "github.com/arthurbcp/kuma/cmd/commands/exec"
	"github.com/arthurbcp/kuma/cmd/commands/modify"
	"github.com/arthurbcp/kuma/cmd/commands/module"
	"github.com/arthurbcp/kuma/internal/debug"
	"github.com/spf13/cobra"
)

// UnicodeLogo holds the ASCII or Unicode art logo for Kuma CLI.
// Currently, it's an empty string but can be populated with logo art.
const (
	UnicodeLogo = `
	
	`
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:  "kuma",
	Long: fmt.Sprintf("%s \n\n Welcome to Kuma! \n A powerful CLI for generating project scaffolds based on Go templates.", UnicodeLogo),
	Run: func(cmd *cobra.Command, args []string) {
		// Display help information when no subcommand is provided.
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		// Exit with status code 1 if an error occurs during command execution.
		os.Exit(1)
	}
}

// init initializes the root command by setting up completion options
// and adding all available subcommands.
func init() {
	// Hide the default completion command to prevent clutter.
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&debug.Debug, "debug", "", false, "Enable debug mode")
	// Add subcommands to the root command.
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(module.ModuleCmd)
	rootCmd.AddCommand(execRun.ExecCmd)
	rootCmd.AddCommand(modify.ModifyCmd)
}
