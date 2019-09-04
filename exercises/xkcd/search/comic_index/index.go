package comic_index

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"xkcd/comic"
	"xkcd/search/comic_index/item"
)

type ComicIndex struct {
}

const name = "xkcd"

const fileNamePrefix = "search_index_"

func (index *ComicIndex) getSourceFilePath() string {
	return fileNamePrefix + name + ".json"
}

func (index *ComicIndex) ReadAll() []comic.Comic {
	file, err := os.OpenFile(index.getSourceFilePath(), os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	indexItems := make([]comic.Comic, 0, 0)
	err = json.Unmarshal(content, &indexItems)
	if err != nil {
		log.Fatal(err)
	}
	return indexItems
}

func (index *ComicIndex) Write(comicIndexItems []item.Comic) {
	file, err := os.OpenFile(index.getSourceFilePath(), os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(comicIndexItems)

	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

func New(comics []comic.Comic) ComicIndex {
	comicIndexItems := make([]item.Comic, 0, len(comics))

	for _, comic_ := range comics {
		comicIndexItems = append(comicIndexItems, *item.New(comic_))
	}

	comicIndex := ComicIndex{}
	comicIndex.Write(comicIndexItems)
	return comicIndex
}
