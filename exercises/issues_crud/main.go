package main

import (
	"os"
	"overview/exercises/issues_crud/cli"
	"overview/exercises/issues_crud/commands"
	"overview/exercises/issues_crud/github"
)

const (
	LIST   = "list"
	EDIT   = "edit"
	CREATE = "create"
	LOCK   = "lock"
	UNLOCK = "unlock"
)

func getBasicAuthCredsFromEnv() github.BasicAuthCreds {
	username := os.Getenv("GIT_HUB_USERNAME")
	password := os.Getenv("GIT_HUB_PASSWORD")
	return github.BasicAuthCreds{Username: username, Password: password}
}
func main() {
	client := github.Client{AuthCreds: getBasicAuthCredsFromEnv()}
	CLI := cli.New()
	CLI.RegisterCommand(LIST, &commands.ListIssues{GithubClient: client})
	CLI.RegisterCommand(EDIT, &commands.EditIssues{GithubClient: client})
	CLI.RegisterCommand(CREATE, &commands.CreateIssue{GithubClient: client})
	CLI.RegisterCommand(LOCK, &commands.LockIssue{GithubClient: client})
	CLI.RegisterCommand(UNLOCK, &commands.UnlockIssue{GithubClient: client})
	CLI.Run()
}
