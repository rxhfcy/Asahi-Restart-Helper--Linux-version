package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RequireCommand(name string, allowedPaths ...string) string {
	var absPath string
	var err error
	if len(allowedPaths) > 0 {
		// look only in allowedPaths
		for _, prefix := range allowedPaths {
			candidatePath := filepath.Join(prefix, name)
			if _, err = os.Stat(candidatePath); err == nil {
				absPath = candidatePath
				break
			}
		}
	} else {
		// look anywhere in PATH
		absPath, err = exec.LookPath(name)
	}
	if err != nil || absPath == "" {
		fmt.Fprintln(os.Stderr, "Required command not found:", name)
		os.Exit(1)
	}
	return absPath
}
