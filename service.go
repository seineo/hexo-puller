package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

func UpdateRepo(repoRemote string, targetDir string) error {
	// check whether repo directory exists
	repoName, err := GetRepoName(repoRemote)
	if err != nil {
		return fmt.Errorf("failed to get repo name from repo remote url: %v", err.Error())
	}
	log.Printf("repo name is %v\n", repoName)
	// check parent folder first
	exists, err := FolderExists(targetDir)
	if err != nil {
		return fmt.Errorf("target folder path is not valid: %v", err.Error())
	}
	if !exists {
		return errors.New("target folder path not exists")
	}
	// check repo directory
	repoLocal := path.Join(targetDir, repoName)
	exists, err = FolderExists(repoLocal)
	if err != nil {
		return fmt.Errorf("target repo path is not valid: %v", err.Error())
	}
	if !exists {
		log.Printf("Local repo does not exist, trying to clone repo from remote....\n")
		err = ExecuteCommand(targetDir, "git", "clone", repoRemote)
		if err != nil {
			return fmt.Errorf("git clone error: %v", err.Error())
		}
	} else {
		log.Printf("Local repo exists, trying to pull repo from remote....\n")
		err = ExecuteCommand(repoLocal, "git", "pull", repoRemote)
		if err != nil {
			return fmt.Errorf("git pull error: %v", err.Error())
		}
	}
	return nil
}

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
