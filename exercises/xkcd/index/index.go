package index

import (
	"xkcd/comic"
)

type Index interface {
	Create()
	Drop()
	AddComic(comic comic.Comic)
	BulkAddComic(comics []comic.Comic)
}
