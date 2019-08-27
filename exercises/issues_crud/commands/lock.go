package commands

import (
	"issues_crud/cli"
	"issues_crud/github"
	"log"
	"os"
	"strconv"
)

type LockIssue struct {
	GithubClient github.Client
}

func (handler *LockIssue) Execute() {
	args := cli.RetrieveArgs(3)
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
