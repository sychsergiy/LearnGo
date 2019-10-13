package interfaces

import (
	"bufio"
	"strings"
)

type WordsCounter int

func (wc *WordsCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	c := 0
	for scanner.Scan() {
		c++
	}
	*wc += WordsCounter(c)
	return c, nil
}

type LinesCounter int

func (lc *LinesCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	c := 0
	for scanner.Scan() {
		c++
	}
	*lc += LinesCounter(c)
	return c, nil
}
