package commands

import (
	"log"
	"os"
	"overview/exercises/issues_crud/cli"
	"overview/exercises/issues_crud/github"
	"strconv"
)

type LockIssue struct {
	BaseCommand  cli.BaseCommand
	GithubClient github.Client
}

func (handler *LockIssue) Execute() {
	args := handler.BaseCommand.RetrieveArgs(3)
	issueNumber, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	err2 := handler.GithubClient.LockIssue(args[0], issueNumber, args[2])
	if err2 != nil {
		log.Fatal(err2)
	}
}

func (handler *LockIssue) parseArgs() (string, int, string) {
	if len(os.Args) < 5 {
		log.Fatal("Not enough arguments")
	}
	repo := os.Args[2]
	issueNumber, err := strconv.Atoi(os.Args[3])
	reason := os.Args[4]
	if err != nil {
		log.Fatal(err)
	}
	return repo, issueNumber, reason
}
