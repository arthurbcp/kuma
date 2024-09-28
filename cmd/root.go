/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/arthurbcp/kuma-cli/cmd/generate"
	"github.com/arthurbcp/kuma-cli/cmd/parser"
	"github.com/spf13/cobra"
)

const (
	UnicodeLogo = `
▗▘▖▘▚▘▘▘▖▖▘▖▝▖▞▗▘▚▘▘▝▝▗▗ ▚▗    ▘▘▘▖▘▘▖▖▞ ▚▗▘▘▘▝▖▘▘▘▚▗▝ ▘▞▗ ▘
▗▝▗▘▘▝▝▝▗▗▘▝▝▗▗▗▝▖▝▝▝▝     ▖▄▐▗▗▗   ▝ ▖▖▚▗▗▝▝▝▖▖▘▘▚▘▖▘▘▚ ▘▞ 
▗▘▖▞▝▐▝▝▖▖▞▝▝▖▖▖▘▝▝▝▝ ▗▄▙▛▀▀▀▝▝▀▀▀▛▟▄   ▗▗▘▝▞▝▗▝▝▝▖▖▖▘▚▗▘▚ ▘
▗▝▗▗▘▘▝▝▗▗▗▘▚▗▗▝▝▝▝ ▗▟▀▝            ▀▀▙▖  ▝▝ ▘▘▞▝▝▗▗▝▝▖▖▞ ▘▘
▗▘▖▖▞▝▝▞▗▗▗▝▖▖▖▘▘ ▗▟▀   ▗▄▙▙    ▄▟▄▖   ▜▙▖ ▘▘▘▚ ▘▚▗▗▘▘▖▖▖▘▘▘
▗▝▗▗▗▘▚ ▘▖▖▘▖▖▞▝ ▄▜    ▗▛▄▄▐█▖▗█▀▖▌█▖   ▝▚▖ ▘▘▘▞▝▖▖▘▝▖▖▞ ▚▘▘
▗▘▖▘▖▝▖▘▚▗▝▖▖▖▖ ▗▞▘    ▐▛▄▐▐▞▌▝▙▜▐▐▚▌    ▝▜▖ ▘▚ ▚▗▝▝▖▖▞ ▘▘▝ 
▗▝▗▘▝▝ ▚▗▗▘▖▝▗  ▟▘  ▄▄▄ ▜▙▙▙▛  ▜▙▙▛▛ ▗▄▄  ▝▟  ▖▚▗ ▘▚▗▗▗▘▘▚▘▘
▗▘▖▘▚▘▘▘▖▖▝▗▘▘ ▐▌ ▗█▀▜▐▙            ▟▛▚▜█▖ ▜▘ ▗▗▗▘▚▗▗▗▗▝▝▖▝▖
▗▝▗▘▖▞▝▝▗▝▝▖▝▖ ▟▘ ▜▟▐▚▚▛▌  ▗▄▄▄▄▖  ▐▙▚▚▚▟▚ ▐▙ ▗ ▘▝▖▖▖▘▖▘▚▝▖▖
▗▘▖▚▗▝▝▝▖▘▚▝▖▖ ▛▌  ▛▙▛▙▛  ▞▙▚▚▚▜▟▄  ▐█▟▜▞  ▗▛  ▐▝▝▗▗▝▖▞▝▗▗▗ 
▗▝▖▖▖▘▘▚ ▚▗▗▗  ▜▖   ▘▀  ▗▛▛▞▞▞▞▄▚▜▜▖  ▘▘   ▐▜  ▘▝▝▖▖▘▖▖▘▘▖▘▖
▗▘▖▝▗▘▘▘▞▗▗▘▖▘ ▐▙      ▟▛▚▚▚▜▐▞▞▞▞▞▛▙    ▗ ▟▘ ▝▝▞▝▗▝▖▖▞▝▝▗▘▖
▗▝▗▘▘▝▝▝▗▗▗▗▝▝  ▙▘    ▐▙▞▌▌▙▚▚▚▜▐▞▞▞▛▌    ▗▜  ▞▝ ▚▗▘▖▖▖▘▘▘▖▖
▗▘▖▘▘▚▘▘▘▖▘▖▘▘▚ ▝▛▖   ▐▟▟▟▟▞▟▚▙▙▌▙▜▞█▘   ▗▜▘ ▖▞▝▝▖▖▖▖▞ ▚▘▚▗ 
▗▝▗▘▘▘▞▝▝▗▘▝▝▝▖▖ ▝▙▚   ▜▟▞▟▟▜▜▟▟▟▜▚▛▌▘  ▐▜  ▗▝ ▘▚▗▗▝▗▗▘▘▝▖▖▖
▗▘▖▘▘▚ ▚▘▘▝▞▝▝▖▘▘  ▀▄▖  ▝▀▘▀▀▘▝▝▘▀▀▀  ▗▞▛  ▐▗▘▚▘▘▖▖▘▘▖▝▝▞▗▗ 
▗▝▗▘▘▚▝▖▝▝▝ ▘▚▝▝▚▚  ▝▀▙▄            ▗▟▀   ▞▖▖▚ ▘▚▗▝▝▝▖▚▘▖▖▖▘
▗▘▖▘▚▗▘▝▝▞▝▞▝▖▝▞▗▖▌▚▗  ▘▀▙▄▄▄▄▄▄▄▟▞▀▘▘ ▗▝▞▗▝▖▘▞▝▖▖▘▚▘▖▘▖▞▗▝ 
▗▝▗▘▘▖▞▝▝▗▘▖▚▝▝ ▘▖▞▝▖▚▘▖     ▘▘     ▖▖▞▝▞▗▘▞ ▚▗▘▖▝▝▖▝▗▘▖▖▘▞ 
▗▘▘▝▝▗▗▘▘▘▖▞▗▝▝▞▝▖▞▝▞▖▚▝▞▐▗▗▖▖▄▝▖▖▌▚▝▖▚▚▝▖▘▖▚▗▗▝▗▘▚▝▝▖▝▖▝▖▖▘`
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "kuma",
	Long: fmt.Sprintf("%s \n\n Welcome to Kuma! \n A powerful CLI for generating project scaffolds based on Go templates.", UnicodeLogo),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(parser.ParseCmd)
}
