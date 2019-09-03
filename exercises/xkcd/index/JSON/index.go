package JSON

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (index *Index) RetrieveComic(num int) *comic.Comic {
	filename := fmt.Sprintf("%d.json", num)
	file, err := os.OpenFile(path.Join(index.getDirName(), filename), os.O_RDONLY, 0660)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	c := &comic.Comic{}
	err = json.Unmarshal(content, c)
	if err != nil {
		log.Fatal(err)
	}
	return c

}

func (index *Index) BulkRetrieveComic(nums []int) []comic.Comic {
	comics := make([]comic.Comic, 0, len(nums))
	for _, num := range nums {
		comics = append(comics, *index.RetrieveComic(num))
	}
	return comics
}
