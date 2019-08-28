package json

import "xkcd/comic"

type IndexJson struct {
}

func (index *IndexJson) Create() {

}

func (index *IndexJson) Drop() {

}

func (index *IndexJson) AddComic(comic *comic.Comic) {

}

func (index *IndexJson) BulkAddComic(comic []comic.Comic) {

}
