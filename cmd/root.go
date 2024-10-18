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

const (
	UnicodeLogo = `
	
	`
)

var rootCmd = &cobra.Command{
	Use:  "kuma",
	Long: fmt.Sprintf("%s \n\n Welcome to Kuma! \n A powerful CLI for generating project scaffolds based on Go templates.", UnicodeLogo),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().BoolVarP(&debug.Debug, "debug", "", false, "Enable debug mode")
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(module.ModuleCmd)
	rootCmd.AddCommand(execRun.ExecCmd)
	rootCmd.AddCommand(modify.ModifyCmd)
}
