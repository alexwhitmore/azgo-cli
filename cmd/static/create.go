/*
Copyright © 2024 Alex Whitmore heyimalexw@gmail.com
*/
package static

import (
	"azgo/cmd"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create Static Site in Azure",
	Long:    `Create a Static Site resource in Azure and deploy your application to it.`,
	Example: `azgo create staticsite --name myStaticSite
  azgo c staticsite --name myStaticSite`,
}

func init() {
	cmd.RootCmd.AddCommand(createCmd)
}
