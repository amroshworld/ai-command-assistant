package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	gl "cloud.google.com/go/ai/generativelanguage/apiv1beta2"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta2/generativelanguagepb"
	"google.golang.org/api/option"
)

const apiKeyFile = "apiKey.txt"

// get api key from usr
func getAPIKey() string {
	// Check if API key  exists
	if file, err := os.ReadFile(apiKeyFile); strings.TrimSpace(string(file)) == "" {
		// Ask user for API key
		var key string
		fmt.Print("Enter your API key: ")
		fmt.Scanln(&key)

		// Trim newline or space
		apiKey := strings.TrimSpace(key)

		// Save API key to file
		err = os.WriteFile(apiKeyFile, []byte(apiKey), 0644)
		Error(err)

	}

	// Read API key from file
	apiKey, err := os.ReadFile(apiKeyFile)
	Error(err)

	return string(apiKey)
}
func generateText() string {
	var output string

	ctx := context.Background()

	//Get API key
	apiKey := getAPIKey()

	client, err := gl.NewTextRESTClient(ctx, option.WithAPIKey(apiKey))
	Error(err)

	defer client.Close()
	sliceOfStrings := os.Args[1:]
	joinedString := strings.Join(sliceOfStrings, " ")
	overrideAction := "write command for " + joinedString
	req := &pb.GenerateTextRequest{
		Model: "models/text-bison-001",
		Prompt: &pb.TextPrompt{
			Text: overrideAction,
		},
	}

	resp, err := client.GenerateText(ctx, req)
	Error(err)

	command := resp.Candidates[0].Output

	if command != "" {
		// extract command from response
		if strings.HasPrefix(command, "```") && strings.HasSuffix(command, "```") {
			withoutMarks := strings.TrimPrefix(strings.TrimSuffix(command, "```"), "```")
			withoutSpaces := strings.TrimPrefix(strings.TrimSuffix(withoutMarks, "\n"), "\n")
			output = withoutSpaces
		} else {

			output = command
		}
	} else {
		fmt.Print("no response received")
	}
	return output
}

// remove apikey from apikey.txt
func clearFile(filename string) error {
	// Open file in write mode
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	Error(err)

	// Close file without writing anything
	err = file.Close()
	Error(err)

	return nil
}

// store text to bash history
func storeText(generativeOutput string) {
	homeDir, err := os.UserHomeDir()
	Error(err)

	historyFile := filepath.Join(homeDir, ".bash_history")

	f, err := os.OpenFile(historyFile, os.O_WRONLY|os.O_APPEND, 0644)
	Error(err)

	defer f.Close()

	_, err = fmt.Fprintf(f, "%s\n", generativeOutput)
	Error(err)

	reloadHistory()

}

// reload history in terminal after storing text
func reloadHistory() {

	err := syscall.Exec("/bin/bash", []string{"history", "-r"}, syscall.Environ())
	Error(err)

	fmt.Println("History successfully reloaded.")
}

func main() {

	frgtCommand := flag.NewFlagSet("frgt", flag.ExitOnError)

	flag.Parse()

	if len(os.Args) < 1 {
		fmt.Println("what command should i search for?")
		os.Exit(2)
	}

	switch os.Args[0] {
	case "frgt":
		if os.Args[1] != "rmvkey" {
			response := generateText()
			if response != "" {
				frgtCommand.Parse(os.Args[2:])
				fmt.Println(response)
				storeText(response)
			} else {
				fmt.Println("no response received")
			}

		} else {
			err := clearFile(apiKeyFile)
			if err != nil {
				fmt.Println("Error clearing API key file:", err)
			} else {
				fmt.Println("API key cleared")
			}
		}
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

}

func Error(err error) {
	if err != nil {
		panic(err)
	}
}
