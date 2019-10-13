package interfaces

import (
	"bufio"
	"io"
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

type BytesCounter struct {
	writer  io.Writer
	counter int64
}

func (bc *BytesCounter) Write(p []byte) (int, error) {
	n, err := bc.writer.Write(p)
	bc.counter += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := BytesCounter{w, 0}
	return &bc, &bc.counter
}
