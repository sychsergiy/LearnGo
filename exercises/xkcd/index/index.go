package index

import (
	"fmt"
	"log"
	"strconv"
	"xkcd/comic"
	"xkcd/index/client"
)

type Index interface {
	Create() error
	Drop() error
	AddComic(comic *comic.Comic) error
	BulkAddComic(comics []comic.Comic) int
	RetrieveComic(num int) *comic.Comic
	BulkRetrieveComic(nums []int) []comic.Comic
}

func Fill(index Index) {
	log.Println("Fetching last comic")
	firstComic := client.FetchLast()
	err := index.AddComic(firstComic)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < firstComic.Num; i++ {
		log.Println("Fetching comic num: " + strconv.Itoa(i))
		err = index.AddComic(client.FetchOne(i))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BulkFill(index Index, bulkSize int) int {
	log.Println("Fetching last comic")
	firstComic := client.FetchLast()
	err := index.AddComic(firstComic)
	if err != nil {
		log.Fatal(err)
	}

	failed := 0
	for i := 1; i < firstComic.Num; i += bulkSize {
		log.Println(fmt.Sprintf("Fetching next comics bulk, first index: %d", i))
		nums := make([]int, 0, bulkSize)
		for j := i; j < i+bulkSize && j < firstComic.Num; j++ {
			nums = append(nums, j)
		}
		bulkFailed := index.BulkAddComic(client.Fetch(nums))
		failed += bulkFailed
		log.Println(fmt.Sprintf("Comics bulk saved, failed items: %d", bulkFailed))
	}
	return failed
}
