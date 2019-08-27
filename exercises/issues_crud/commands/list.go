package commands

import (
	"encoding/json"
	"log"
	"overview/exercises/issues_crud/github"
)

type ListIssues struct {
	GithubClient github.Client
}

func (handler *ListIssues) Execute() {
	issues, err := handler.GithubClient.ListIssues()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ISSUES:")
	for _, issue := range issues {
		marshaled, _ := json.MarshalIndent(issue, "", "\t")
		log.Println(string(marshaled))
	}
}
