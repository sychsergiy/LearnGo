package index

import (
	"xkcd/comic"
)

type Index interface {
	Create() error
	Drop() error
	AddComic(comic comic.Comic) error
	BulkAddComic(comics []comic.Comic)
	RemoveComic()
}
