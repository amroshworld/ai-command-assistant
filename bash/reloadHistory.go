package bash

import (
	"frgt/errormessage"
	"syscall"
)

// reload history in terminal after storing text
func ReloadHistory() {

	// store current session history to history file prevent from unexpected history conflict from other sessions

	err := syscall.Exec("/bin/bash", []string{"history", "-r"}, syscall.Environ())
	errormessage.Error(err)

	err = syscall.Exec("/bin/bash", []string{"history", "-a"}, syscall.Environ())
	errormessage.Error(err)

}
