package commands

import (
	"log"
	"overview/exercises/issues_crud/github"
)

type LockIssue struct {
	GithubClient github.Client
}

func (handler *LockIssue) Execute() {
	err := handler.GithubClient.LockIssue("LearnGo", 1, "any reason")
	if err != nil {
		log.Fatal(err)
	}
}
