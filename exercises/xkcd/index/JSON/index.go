package JSON

import (
	"fmt"
	"log"
	"os"
	"path"
	"xkcd/comic"
)

const indexDirPrefix = "json_index"

type Index struct {
	Name string
}

func (index *Index) getDirName() string {
	return indexDirPrefix + "_" + index.Name
}

func (index *Index) Create() error {
	err := os.Mkdir(index.getDirName(), 0700)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (index *Index) Drop() error {
	err := os.RemoveAll(index.getDirName())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (index *Index) AddComic(comic *comic.Comic) error {
	filename := fmt.Sprintf("%d.json", comic.Num)
	f, err := os.Create(path.Join(index.getDirName(), filename))
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = f.WriteString(comic.Marshal())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (index *Index) BulkAddComic(comics []comic.Comic) int {
	var failed int
	for _, c := range comics {
		err := index.AddComic(&c)
		if err != nil {
			failed += 1
		}
	}
	return failed
}
