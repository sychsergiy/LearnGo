package commands

import (
	"encoding/json"
	"issues_crud/cli"
	"issues_crud/github"
	"log"
	"strconv"
)

type EditIssues struct {
	GithubClient github.Client
}

func (handler *EditIssues) Execute() {
	args := cli.RetrieveArgs(4)
	repo, title, state := args[0], args[2], args[3]
	body := cli.RetrieveArgFromEditor()
	issueNumber, convErr := strconv.Atoi(args[1])
	if convErr != nil {
		log.Fatal(convErr)
	}
	data := github.EditIssueRequestData{Title: title, Body: body, State: state}
	issue, err := handler.GithubClient.EditIssue(repo, issueNumber, data)
	if err != nil {
		log.Fatal(err)
	}
	if issue != nil {
		marshaled, _ := json.MarshalIndent(issue, "", "\t")
		log.Println(string(marshaled))
	}
}
