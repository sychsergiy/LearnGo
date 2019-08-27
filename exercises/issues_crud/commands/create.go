package commands

import (
	"encoding/json"
	"issues_crud/cli"
	"issues_crud/github"
	"log"
)

type CreateIssue struct {
	GithubClient github.Client
}

func (handler *CreateIssue) Execute() {
	args := cli.RetrieveArgs(2)
	repo, title := args[0], args[1]
	body := cli.RetrieveArgFromEditor()
	data := github.CreateIssueRequestData{Title: title, Body: body}
	issue, err := handler.GithubClient.CreateIssue(repo, data)
	if err != nil {
		log.Fatal(err)
	}
	if issue != nil {
		marshaled, _ := json.MarshalIndent(issue, "", "\t")
		log.Println(string(marshaled))
	}
}
