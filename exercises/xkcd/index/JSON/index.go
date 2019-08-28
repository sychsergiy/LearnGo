package JSON

import (
	"fmt"
	"log"
	"os"
	"path"
	"xkcd/comic"
)

const indexDirPath = "json_index"

type Index struct {
}

func (index *Index) Create() error {
	err := os.Mkdir(indexDirPath, 0700)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (index *Index) Drop() error {
	err := os.RemoveAll(indexDirPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (index *Index) AddComic(comic *comic.Comic) error {
	filename := fmt.Sprintf("%d.json", comic.Num)
	f, err := os.Create(path.Join(indexDirPath, filename))
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

func (index *Index) BulkAddComic(comic []comic.Comic) {

}
