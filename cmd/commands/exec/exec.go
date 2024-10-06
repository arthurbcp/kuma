// run.go
//
// Package run defines the 'run' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package exec

import (
	execHandlers "github.com/arthurbcp/kuma-cli/cmd/commands/exec/handlers"

	"github.com/spf13/cobra"
)

var (
	Run string
)

// ExecRunCmd represents the 'run' subcommand.
var ExecRunCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a specific pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		vars := map[string]interface{}{
			"data": map[string]interface{}{},
		}
		execHandlers.HandleRun("initial", vars)
	},
}

// init sets up flags for the 'run' subcommand and binds them to variables.
func init() {
	// Repository name
	ExecRunCmd.Flags().StringVarP(&Run, "run", "r", "", "run to use")
	ExecRunCmd.MarkFlagRequired("run")
}
