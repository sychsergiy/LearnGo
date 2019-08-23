package main

import "log"

func main() {
	//SearchIssues()
	issues, err := ListIssues(AuthCredentials{"test", "test"})

	if issues != nil {
		for _, issue := range *issues {
			log.Print(issue)
		}
	}
	if err != nil {
		log.Fatal(err)
	}

}
