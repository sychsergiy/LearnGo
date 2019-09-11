package JSON

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"xkcd/comic"
)

const indexDirPrefix = "json_index"

type Index struct {
	Name string
}

func (index *Index) getDirName() string {
	return indexDirPrefix + "_" + index.Name
}

func (index *Index) Create() {
	err := os.Mkdir(index.getDirName(), 0700)
	if err != nil {
		log.Fatal(err)
	}
}

func (index *Index) Drop() {
	err := os.RemoveAll(index.getDirName())
	if err != nil {
		log.Fatal(err)
	}
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

func (index *Index) RetrieveAllComics() []comic.Comic {
	nums := index.getAllComicsNums()
	return index.BulkRetrieveComic(nums)
}

func (index *Index) AllComicsIterator(chunkSize int) (func() ([]comic.Comic, bool), bool) {
	nums := index.getAllComicsNums()
	var chunkEnd, chunkStart int
	totalLen := len(nums)

	return func() ([]comic.Comic, bool) {
		chunkEnd = chunkStart + chunkSize
		if chunkEnd > totalLen {
			chunkEnd = totalLen
		}
		comics := index.BulkRetrieveComic(nums[chunkStart:chunkEnd])
		chunkStart = chunkEnd
		return comics, chunkEnd < totalLen
	}, chunkEnd < totalLen
}

func (index *Index) getAllComicsNums() []int {
	filesNames := getDirFiles(index.getDirName())

	nums := make([]int, 0, len(filesNames))
	for _, filename := range filesNames {
		chunks := strings.Split(filename, ".")
		num, err := strconv.Atoi(chunks[0])
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func getDirFiles(dirPath string) []string {
	files, err := ioutil.ReadDir(dirPath)

	filesNames := make([]string, 0, len(files))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		filesNames = append(filesNames, f.Name())
	}
	return filesNames
}
