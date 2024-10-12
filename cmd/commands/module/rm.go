// module.go
//
// Package get defines the 'module' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package module

import (
	"os"
	"os/exec"

	"github.com/arthurbcp/kuma/cmd/shared"
	"github.com/spf13/cobra"
)

var (
	Module string
)

// Add a Kuma module from a GitHub repository
var ModuleRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a Kuma module",
	Run: func(cmd *cobra.Command, args []string) {
		removeGitSubmodule(Module)
	},
}

// remove git submodule removes a submodule from Kuma
func removeGitSubmodule(module string) error {
	// Full path to the submodule inside .kuma
	fullSubmodulePath := shared.KumaFilesPath + "/" + module

	// 1. Remove submodule config from .git/config
	removeConfigCmd := exec.Command("git", "config", "--remove-section", "submodule."+fullSubmodulePath)
	removeConfigCmd.Stdout = os.Stdout
	removeConfigCmd.Stderr = os.Stderr
	if err := removeConfigCmd.Run(); err != nil {
		return err
	}

	// 2. Remove submodule entry from .gitmodules if it exists
	removeFromGitmodulesCmd := exec.Command("git", "config", "-f", ".gitmodules", "--remove-section", "submodule."+fullSubmodulePath)
	removeFromGitmodulesCmd.Stdout = os.Stdout
	removeFromGitmodulesCmd.Stderr = os.Stderr
	if err := removeFromGitmodulesCmd.Run(); err != nil {
		return err
	}

	// 3. Remove the submodule from git cache
	cmdRmCached := exec.Command("git", "rm", "--cached", fullSubmodulePath)
	cmdRmCached.Stdout = os.Stdout
	cmdRmCached.Stderr = os.Stderr
	if err := cmdRmCached.Run(); err != nil {
		return err
	}

	// 4. Physically remove the submodule directory
	if err := os.RemoveAll(fullSubmodulePath); err != nil {
		return err
	}

	return nil
}

// init sets up flags for the 'rm' subcommand and binds them to variables.
func init() {
	// Module name
	ModuleRmCmd.Flags().StringVarP(&Module, "rm", "m", "", "module to remove")
}
