package data_structures

import (
	"github.com/stretchr/testify/assert"
	"reflect"
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

func Test_rotate(t *testing.T) {
	input := []int{1, 2, 3, 4}
	rotate(input, 2)
	assert.Equal(t, input, []int{3, 4, 1, 2})
}
func Test_rotateByCopy(t *testing.T) {
	input := []int{1, 2, 3, 4}
	rotateByCopy(input, 2)
	assert.True(t, reflect.DeepEqual([]int{3, 4, 1, 2}, input))

	input3 := []int{1, 2, 3, 4, 5, 6, 7}
	rotateByCopy(input3, 2)
	assert.True(t, reflect.DeepEqual([]int{3, 4, 5, 6, 7, 1, 2}, input3))

	input2 := []int{1, 2, 3, 4, 5, 6, 7}
	rotateByCopy(input2, 5)
	assert.True(t, reflect.DeepEqual([]int{6, 7, 1, 2, 3, 4, 5}, input2))
}
