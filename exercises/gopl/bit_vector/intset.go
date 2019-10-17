package bit_vector

import (
	"bytes"
	"fmt"
	"strings"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	var wordsCopy = make([]uint64, len(s.words))
	copy(wordsCopy, s.words)
	newS := &IntSet{wordsCopy}
	return newS
}

func (s *IntSet) Elems() []int {
	var items []int

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				items = append(items, 64*i+j)
			}
		}
	}
	return items
}

func (s *IntSet) AddAll(items ...int) int {
	var count = 0
	for _, y := range items {
		if s.Add(y) {
			count += 1
		}
	}
	return count
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Remove(x int) bool {
	word, bit := x/64, uint(x%64)

	if word < len(s.words) {
		if s.words[word]&(1<<bit) == 0 {
			return false
		} else {
			s.words[word] ^= 1 << bit
			return true
		}
	} else {
		return false
	}
}

func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count += 1
			}
		}
	}
	return count
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) bool {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	previous := s.words[word]
	s.words[word] |= 1 << bit
	if previous != s.words[word] {
		return true
	}
	return false
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	buf.WriteString(strings.Trim(strings.Replace(fmt.Sprint(s.Elems()), " ", ", ", -1), "[]"))
	buf.WriteByte('}')
	return buf.String()
}
