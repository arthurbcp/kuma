// module.go
//
// Package get defines the 'module' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package module

import "github.com/spf13/cobra"

var ModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Manage Kuma modules",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ModuleCmd.AddCommand(ModuleAddCmd)
	ModuleCmd.AddCommand(ModuleRmCmd)
}
