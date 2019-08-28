package comic

import (
	"encoding/json"
	"log"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func (c *Comic) Marshal() string {
	marshalled, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(marshalled)
}
