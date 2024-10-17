package exec

import (
	execModule "github.com/arthurbcp/kuma/cmd/commands/exec/module"
	execRun "github.com/arthurbcp/kuma/cmd/commands/exec/run"
	"github.com/spf13/cobra"
)

var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Manage Kuma execs",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ExecCmd.AddCommand(execRun.ExecCmd)
	ExecCmd.AddCommand(execModule.ExecModuleCmd)
}
