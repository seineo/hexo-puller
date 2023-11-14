package main

import (
	"encoding/json"
	"fmt"
	"hexo-puller/config"
	"log"
	"net/http"
)

type TaskInfo struct {
	RepoUrl   string `json:"repoUrl"`
	TargetDir string `json:"targetDir"`
}

func updateRepo(w http.ResponseWriter, r *http.Request) {
	log.Println("got update request")

	// get repo path and target dir from requests
	decoder := json.NewDecoder(r.Body)
	var t TaskInfo
	err := decoder.Decode(&t)
	if err != nil {
		log.Printf("requst body is invalid: %v\n", err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// update repo
	log.Println("------ start to update repository ------")
	err = UpdateRepo(t.RepoUrl, t.TargetDir)
	if err != nil {
		log.Printf("failed to update repo: %v\n", err.Error())
		return
	}
	log.Println("------ repository successfully updated ------")
}

// just for test
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
}

func main() {
	config := config.GetConfig()

	http.HandleFunc("/", updateRepo)
	http.HandleFunc("/hello", HelloServer)

	port := ":33333"
	log.Printf("server starts, listening at port%v\n", port)
	err := http.ListenAndServeTLS(port, config.Tls.Crt, config.Tls.Key, nil)
	if err != nil {
		log.Fatalf("server errror: %v\n", err.Error())
	}
}
