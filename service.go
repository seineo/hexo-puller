package main

import (
	"os"
	"os/exec"
)

func ExecuteCommand(workDir string, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = workDir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run() // blocks until sub process completes
	if err != nil {
		return err
	}
	return nil
}
