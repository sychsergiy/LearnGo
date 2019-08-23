package main

import (
	"encoding/json"
	"log"
	"os"
)

func getBasicAuthCredsFromEnv() BasicAuthCreds {
	username := os.Getenv("GIT_HUB_USERNAME")
	password := os.Getenv("GIT_HUB_PASSWORD")
	return BasicAuthCreds{username, password}
}

func main() {
	creds := getBasicAuthCredsFromEnv()
	issues, err := ListIssues(creds)

	if issues != nil {
		for _, issue := range *issues {
			json, _ := json.MarshalIndent(issue, "", "\t")
			log.Println(string(json))
		}
	}
	if err != nil {
		log.Fatal(err)
	}

}
