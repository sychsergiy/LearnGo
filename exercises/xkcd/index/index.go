package index

import (
	"log"
	"xkcd/comic"
	"xkcd/index/client"
)

type Index interface {
	Create() error
	Drop() error
	AddComic(comic *comic.Comic) error
	BulkAddComic(comics []comic.Comic)
	RemoveComic(num int) error
}

func Fill(index Index) {
	firstComic := client.FetchLast()
	err := index.AddComic(firstComic)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < firstComic.Num; i++ {
		err = index.AddComic(client.FetchOne(i))
		if err != nil {
			log.Fatal(err)
		}
	}
}
