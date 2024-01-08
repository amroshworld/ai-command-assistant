package textcleaning

import (
	"regexp"
)

func ExtractCommand(command string) string {

	var output string
	// Regular expression to match strings between "***"
	re := regexp.MustCompile(`\*\*\*(.*?)\*\*\*`)

	matches := re.FindAllStringSubmatch(command, -1)

	// Extract and print the matched strings, ensuring clarity and conciseness
	for _, match := range matches {
		extractedString := match[1]
		output = extractedString

	}
	return output
}
