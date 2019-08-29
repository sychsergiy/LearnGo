package main

import (
	"xkcd/index"
	"xkcd/index/JSON"
)

func main() {
	jsonIndex := &JSON.Index{}
	_ = jsonIndex.Create()
	index.Fill(jsonIndex)
	jsonIndex.Drop()
}
