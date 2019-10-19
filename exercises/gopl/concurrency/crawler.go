package main

import (
	"fmt"
	"log"
	"os"
	"overview/exercises/gopl/concurrency/links"
	"strconv"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type searchLink struct {
	link  string
	depth int
}

func main() {
	worklist := make(chan []searchLink)  // lists of URLs, may have duplicates
	unseenLinks := make(chan searchLink) // de-duplicated URLs

	// Add command-line arguments to worklist.
	firstLink := searchLink{os.Args[1], 1}
	go func() { worklist <- []searchLink{firstLink} }()

	maxDepth, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for item := range unseenLinks {
				foundLinks := crawl(item.link)

				if !(item.depth > maxDepth) {
					searchItems := make([]searchLink, len(foundLinks))
					for _, link := range foundLinks {
						searchItems = append(searchItems, searchLink{link, item.depth + 1})
					}
					go func() { worklist <- searchItems }()
				}

			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, searchItem := range list {
			if !seen[searchItem.link] {
				seen[searchItem.link] = true
				unseenLinks <- searchItem
			}
		}
	}
}
