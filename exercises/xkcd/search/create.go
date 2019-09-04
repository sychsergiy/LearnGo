package search

import (
	"xkcd/comic"
	"xkcd/index"
)

func CreateSearchIndex(comics []comic.Comic) {
	comicIndexItems := make([]ComicIndexItem, 0, len(comics))

	for _, comic_ := range comics {
		comicIndexItems = append(comicIndexItems, *CreateComicIndexItem(comic_))
	}

	comicIndex := ComicIndex{}
	comicIndex.Write(comicIndexItems)
}

func CreateSearchIndexFromOfflineIndex(index index.Index) {
	comics := index.BulkRetrieveComic([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) // todo: change on retrieve on from index
	CreateSearchIndex(comics)
}
