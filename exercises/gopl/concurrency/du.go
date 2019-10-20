// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type dirInfo struct {
	name string
	size int64
}

type rootInfo struct {
	size   int64
	nFiles int64
}

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	rootsMap := make(map[string]*rootInfo)
	fileSizes := make(chan dirInfo)
	var n sync.WaitGroup
	for _, root := range roots {
		rootsMap[root] = &rootInfo{}
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case dirInfo, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			rootsMap[dirInfo.name].nFiles += 1
			rootsMap[dirInfo.name].size += dirInfo.size
		case <-tick:
			for key, value := range rootsMap {
				printDiskUsage(key, value.nFiles, value.size)
			}
			printDiskUsage("total", nfiles, nbytes)
		}
	}
	for key, value := range rootsMap {
		printDiskUsage(key, value.nFiles, value.size)
	}
	printDiskUsage("total", nfiles, nbytes) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(dirname string, nfiles, nbytes int64) {
	fmt.Printf("%s dir %d files  %.1f GB\n", dirname, nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(currentDir, rootDir string, n *sync.WaitGroup, fileSizes chan<- dirInfo) {
	defer n.Done()
	for _, entry := range dirents(currentDir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(currentDir, entry.Name())
			go walkDir(subdir, rootDir, n, fileSizes)
		} else {
			fileSizes <- dirInfo{rootDir, entry.Size()}
		}
	}
}

//!-walkDirint

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...int
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
