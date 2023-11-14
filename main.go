package main

import (
	"hexo-puller/config"
	"log"
	"path"
)

func main() {
	conf := config.GetConfig()
	log.Printf("url: %v\npath: %v\n", conf.Url, conf.Path)
	// check whether repo directory exists
	repoName, err := GetRepoName(conf.Url)
	if err != nil {
		log.Fatalf("failed to get repo name, error: %v\n", err.Error())
	}
	log.Printf("repo name is %v\n", repoName)
	// check parent folder first
	exists, err := FolderExists(conf.Path)
	if err != nil {
		log.Fatalf("given path is not valid, error: %v\n", err.Error())
	}
	if !exists {
		log.Fatalf("given path not exists")
	}
	// check repo directory
	repoLocal := path.Join(conf.Path, repoName)
	exists, err = FolderExists(repoLocal)
	if err != nil {
		log.Fatalf("local repo path is not valid, error: %v\n", err.Error())
	}
	if !exists {
		log.Printf("Local repo does not exist, trying to clone repo from remote....\n")
		err = ExecuteCommand(conf.Path, "git", "clone", conf.Url)
		if err != nil {
			log.Printf("git clone error: %v\n", err.Error())
		}
	} else {
		log.Printf("Local repo exists, trying to pull repo from remote....\n")
		err = ExecuteCommand(repoLocal, "git", "pull", conf.Url)
		if err != nil {
			log.Printf("git pull error: %v\n", err.Error())
		}
	}
}
