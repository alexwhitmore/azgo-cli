/*
Copyright © 2024 Alex Whitmore heyimalexw@gmail.com
*/
package secrets

import (
	"azgo/cmd"

	"github.com/spf13/cobra"
)

var siteName string
var resourceGroup string

var secretCmd = &cobra.Command{
	Use:     "secret",
	Short:   "Manage secrets for Azure Static Web Apps",
	Long:    "Create, list, update, or delete Azure app settings/environment variables.",
	Example: `azgo secret set -n myStaticSiteNew -r remix-rg MY_SECRET_KEY1=mysecretvalue1 MY_SECRET_KEY2=mysecretvalue2`,
}

func init() {
	cmd.RootCmd.AddCommand(secretCmd)
}
