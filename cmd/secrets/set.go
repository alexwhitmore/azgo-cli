/*
Copyright © 2024 Alex Whitmore heyimalexw@gmail.com
*/
package secrets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"azgo/cmd/auth"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set key=value [key=value...]",
	Short: "Set one or more secrets for Azure Static Web Apps",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		secrets := make(map[string]string)
		for _, kv := range args {
			parts := strings.SplitN(kv, "=", 2)
			if len(parts) != 2 {
				log.Fatalf("Invalid key-value pair: %s", kv)
			}
			key, value := parts[0], parts[1]
			secrets[key] = value
		}
		setSecrets(secrets)
	},
}

func init() {
	secretCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&siteName, "name", "n", "", "Name of the static site (required)")
	setCmd.Flags().StringVarP(&resourceGroup, "resource-group", "r", "", "Azure resource group (required)")
	setCmd.MarkFlagRequired("name")
	setCmd.MarkFlagRequired("resource-group")
}

// Merge existing and new secrets, then upload all secrets to the Azure Static Site.
func setSecrets(newSecrets map[string]string) {
	token, err := auth.GetAccessToken()
	if err != nil {
		log.Fatalf("failed to get access token: %v", err)
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		log.Fatalf("Azure subscription ID must be set in environment variables")
	}

	currentSecrets, err := getCurrentSecrets(token, subscriptionID, resourceGroup, siteName)
	if err != nil {
		log.Fatalf("failed to retrieve current secrets: %v", err)
	}

	for key, value := range newSecrets {
		currentSecrets[key] = value
	}

	err = updateSecrets(token, subscriptionID, resourceGroup, siteName, currentSecrets)
	if err != nil {
		log.Fatalf("failed to set secrets: %v", err)
	}

	fmt.Println("Secrets set successfully")
}

// Retrieve all current secrets inside an Azure Static Site.
func getCurrentSecrets(token, subscriptionID, resourceGroup, siteName string) (map[string]string, error) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/staticSites/%s/listAppSettings?api-version=2023-12-01", subscriptionID, resourceGroup, siteName)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to retrieve secrets: %s, response: %s", resp.Status, responseBody)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	properties, ok := result["properties"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	secrets := make(map[string]string)
	for key, value := range properties {
		secrets[key] = fmt.Sprintf("%v", value)
	}

	return secrets, nil
}

// Upload secrets to an Azure Static Site.
func updateSecrets(token, subscriptionID, resourceGroup, siteName string, secrets map[string]string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/staticSites/%s/config/appsettings?api-version=2023-12-01", subscriptionID, resourceGroup, siteName)

	payload := map[string]interface{}{
		"properties": secrets,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		responseBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to set secrets: %s, response: %s", resp.Status, responseBody)
	}

	return nil
}
