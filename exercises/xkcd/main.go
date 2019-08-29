package main

import (
	"fmt"
	"log"
	"xkcd/index"
	"xkcd/index/JSON"
)

func main() {
	jsonIndex := &JSON.Index{Name: "test"}
	_ = jsonIndex.Drop()
	_ = jsonIndex.Create()
	failed := index.BulkFill(jsonIndex, 100)
	log.Println(fmt.Sprintf("Finish filling index, failed items: %d", failed))
	//_ = jsonIndex.Drop()
}
