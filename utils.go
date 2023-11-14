package main

import (
	"errors"
	"os"
	"strings"
)

// GetRepoName gets repo name from given url with basic url check.
// When executing shell commands, git will do more checks.
func GetRepoName(url string) (string, error) {
	if len(url) == 0 {
		return "", errors.New("repo url is empty")
	}
	slashParts := strings.Split(url, "/")
	targetParts := strings.Split(slashParts[len(slashParts)-1], ".")
	if len(targetParts) <= 2 {
		return targetParts[0], nil
	} else {
		return "", errors.New("repo url is invalid")
	}
}

// FolderExists checks whether the file at the given path is a folder and exists.
func FolderExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return true, err
	}
	if !info.IsDir() {
		return true, errors.New("given path is not a folder")
	}
	return true, nil
}
