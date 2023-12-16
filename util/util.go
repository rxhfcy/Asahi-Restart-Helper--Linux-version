package util

import (
	"fmt"
	"os"
	"os/exec"
)

func RequireCommand(name string) string {
	absPath, err := exec.LookPath(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Required command not found:", name)
		os.Exit(1)
	}
	return absPath
}
