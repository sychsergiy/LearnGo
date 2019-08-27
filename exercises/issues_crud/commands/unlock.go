package commands

import (
	"issues_crud/cli"
	"issues_crud/github"
	"log"
	"strconv"
)

type UnlockIssue struct {
	GithubClient github.Client
}

func (handler *UnlockIssue) Execute() {
	args := cli.RetrieveArgs(2)
	repo := args[0]
	issueNumber, convErr := strconv.Atoi(args[1])
	if convErr != nil {
		log.Fatal(convErr)
	}
	err := handler.GithubClient.UnlockIssue(repo, issueNumber)
	if err != nil {
		log.Fatal(err)
	}
}
