package commands

import (
	"log"
	"overview/exercises/issues_crud/github"
)

type UnlockIssue struct {
	GithubClient github.Client
}

func (handler *UnlockIssue) Execute() {
	err := handler.GithubClient.UnlockIssue("LearnGo", 1)
	if err != nil {
		log.Fatal(err)
	}
}
