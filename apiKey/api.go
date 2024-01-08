package apikey

import (
	"fmt"
	"frgt/errormessage"
	"os"
	"strings"
)

const apiKeyFile = "apiKey.txt"

// get api key from usr
func GetAPIKey() string {
	// Check if API key  exists
	if file, err := os.ReadFile(apiKeyFile); strings.TrimSpace(string(file)) == "" {
		// Ask user for API key
		var key string
		fmt.Println("To use this app, you will need to get your own API key from Gemini Pro. Here are the steps to do that:\n\n1. Go to [makersuite.google.com/app/apikey](https://beebom.com/how-use-google-gemini-api-key/) and sign in with your Google account.\n2. Under API keys, click the 'Create API key in new project' button.\n3. Copy the API key and keep it private. Do not publish or share the API key publicly.\n4. Paste the API key in the config file of this app.\n5. Enjoy using the Gemini Pro API command line app.\n\nEnter your API key: ")
		fmt.Scanln(&key)

		// Trim newline or space
		apiKey := strings.TrimSpace(key)

		// Save API key to file
		err = os.WriteFile(apiKeyFile, []byte(apiKey), 0644)
		errormessage.Error(err)

	}

	// Read API key from file
	apiKey, err := os.ReadFile(apiKeyFile)
	errormessage.Error(err)

	return string(apiKey)
}

// remove apikey from apikey.txt
func ClearFile() {
	// Open file in write mode
	file, err := os.OpenFile(apiKeyFile, os.O_WRONLY|os.O_TRUNC, 0644)
	errormessage.Error(err)

	// Close file without writing anything
	err = file.Close()
	errormessage.Error(err)

	fmt.Println("API key removed.")
}
