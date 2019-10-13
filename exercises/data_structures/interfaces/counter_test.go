package interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordsCounter_Write(t *testing.T) {
	var counter WordsCounter = 0
	integer, err := counter.Write([]byte("first  second, third"))

	assert.Equal(t, 3, integer)
	assert.Equal(t, WordsCounter(3), counter)
	assert.Nil(t, err)
}

func TestLinesCounter_Write(t *testing.T) {
	var counter LinesCounter = 0
	integer, err := counter.Write([]byte("first\n second\t third"))
	assert.Equal(t, LinesCounter(2), counter)
	assert.Equal(t, integer, 2)
	assert.Nil(t, err)
}
