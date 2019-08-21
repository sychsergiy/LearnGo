package data_structures

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWordsFrequency(input string) map[string]int {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()] += 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return counts
}
