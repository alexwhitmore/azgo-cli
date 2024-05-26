/*
Copyright © 2024 Alex Whitmore heyimalexw@gmail.com
*/
package static

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var siteName string
var resourceGroup string
var sourceRepo string
var branchName string

var staticSiteCmd = &cobra.Command{
	Use:   "static",
	Short: "Create an Azure Static Web App and deploy an application from GitHub to it.",
	Run: func(cmd *cobra.Command, args []string) {
		createStaticSite(siteName, resourceGroup, sourceRepo, branchName)
	},
}

func init() {
	createCmd.AddCommand(staticSiteCmd)
	staticSiteCmd.Flags().StringVarP(&siteName, "name", "n", "", "Name of the static site (required)")
	staticSiteCmd.Flags().StringVarP(&resourceGroup, "resource-group", "r", "", "Azure resource group (required)")
	staticSiteCmd.Flags().StringVarP(&sourceRepo, "source", "s", "", "Source GitHub repository URL (required)")
	staticSiteCmd.Flags().StringVarP(&branchName, "branch", "b", "main", "Branch name (default is 'main')")
	staticSiteCmd.MarkFlagRequired("name")
	staticSiteCmd.MarkFlagRequired("resource-group")
	staticSiteCmd.MarkFlagRequired("source")
}

func createStaticSite(siteName, resourceGroup, sourceRepo, branchName string) {
	fmt.Printf("Creating Azure Static Web App with name: %s\n", siteName)

	createCmd := exec.Command("az", "staticwebapp", "create", "--name", siteName, "--source", sourceRepo, "--resource-group", resourceGroup, "--branch", branchName, "--login-with-github")
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	err := createCmd.Run()
	if err != nil {
		log.Fatalf("Failed to create static site: %v", err)
	}

	fmt.Println("Azure Static Web App created successfully")
}
