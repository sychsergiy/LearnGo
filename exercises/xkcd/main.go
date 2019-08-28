package main

import (
	"xkcd/client"
	"xkcd/index/JSON"
)

func main() {
	comic := client.FetchLast()

	index := JSON.Index{}
	_ = index.Create()
	_ = index.AddComic(comic)
	_ = index.Drop()
}
