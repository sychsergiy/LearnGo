package main

import (
	"encoding/json"
	"fmt"
	"log"
	"xkcd/client"
)

func main() {
	comic := client.FetchLast()
	marshalledComic, err := json.Marshal(comic)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(marshalledComic))
}
