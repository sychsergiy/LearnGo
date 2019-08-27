package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"xkcd/comic"
)

func FetchOne(number int) *comic.Comic {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", number)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	c := &comic.Comic{}
	if err := json.NewDecoder(resp.Body).Decode(c); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	return c
}
