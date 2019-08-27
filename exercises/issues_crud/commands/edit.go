package commands

import (
	"encoding/json"
	"log"
	"overview/exercises/issues_crud/github"
)

type EditIssues struct {
	GithubClient github.Client
}

func (handler *EditIssues) Execute() {
	data := github.EditIssueRequestData{Title: "New title", Body: "edited Body", State: "closed"}
	issue, err := handler.GithubClient.EditIssue("LearnGo", 3, data)
	if err != nil {
		log.Fatal(err)
	}
	if issue != nil {
		marshaled, _ := json.MarshalIndent(issue, "", "\t")
		log.Println(string(marshaled))
	}
}
