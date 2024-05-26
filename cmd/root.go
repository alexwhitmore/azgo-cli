/*
Copyright © 2024 Alex Whitmore heyimalexw@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "azgo",
	Short: "A CLI tool to help you deploy your static site to Azure.",
	Long:  `A CLI tool to create a static site resource in Azure, and also deploy your application code to it.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
