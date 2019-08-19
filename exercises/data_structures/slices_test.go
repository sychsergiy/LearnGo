package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	expected := []int{4, 3, 0, 1}
	input := []int{1, 0, 3, 4}
	reverse(input)
	assert.Equal(t, input, expected)
}

func Test_reverseArray(t *testing.T) {
	expected := []int{4, 3, 0, 1}
	input := []int{1, 0, 3, 4}
	reverseArray(&input)
	assert.Equal(t, input, expected)
}
