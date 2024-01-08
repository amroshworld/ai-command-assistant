package bash

import (
	"fmt"
	"frgt/errormessage"
	"os"
	"path/filepath"
)

// store text to bash history
func StoreText(generativeOutput string) {

	homeDir, err := os.UserHomeDir()
	errormessage.Error(err)

	historyFile := filepath.Join(homeDir, ".bash_history")

	f, err := os.OpenFile(historyFile, os.O_WRONLY|os.O_APPEND, 0644)
	errormessage.Error(err)

	defer f.Close()

	_, err = fmt.Fprintf(f, "%s\n", generativeOutput)
	errormessage.Error(err)

}
