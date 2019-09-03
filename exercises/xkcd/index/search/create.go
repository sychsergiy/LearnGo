package search

import (
	"xkcd/comic"
)

func CreateSearchIndex(comics []comic.Comic) {
	comicIndexItems := make([]ComicIndexItem, len(comics))

	for _, comic_ := range comics {
		comicIndexItems = append(comicIndexItems, *CreateComicIndexItem(comic_))
	}

	WriteComicIndex(comicIndexItems)
}
