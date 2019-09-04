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

func (index *ComicIndex) ReadAll() []item.Comic {
	file, err := os.OpenFile(index.getSourceFilePath(), os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	indexItems := make([]item.Comic, 0, 0)
	err = json.Unmarshal(content, &indexItems)
	if err != nil {
		log.Fatal(err)
	}
	return indexItems
}

func New(comics []comic.Comic) ComicIndex {
	comicIndex := ComicIndex{}
	comicIndex.Write(comics)
	return comicIndex
}

func NewEmpty() ComicIndex {
	ins := ComicIndex{}
	ins.writeEmpty()
	return ins
}

func (index *ComicIndex) Drop() {
	err := os.Remove(index.getSourceFilePath())
	if err != nil {
		log.Fatal(err)
	}
}

func (index *ComicIndex) Append(comics []comic.Comic) {
	indexItems := index.ReadAll()
	items := toComicIndexItems(comics)
	index.write(append(indexItems, items...))
}

func toComicIndexItems(comics []comic.Comic) []item.Comic {
	comicIndexItems := make([]item.Comic, 0, len(comics))
	for _, comic_ := range comics {
		comicIndexItems = append(comicIndexItems, *item.New(comic_))
	}
	return comicIndexItems
}

func (index *ComicIndex) Write(comics []comic.Comic) {
	index.write(toComicIndexItems(comics))
}

func (index *ComicIndex) writeEmpty() {
	file, err := os.OpenFile(index.getSourceFilePath(), os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal([]string{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

func (index *ComicIndex) write(comicIndexItems []item.Comic) {
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
