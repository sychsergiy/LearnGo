package bit_vector

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestIntSet_Len(t *testing.T) {
	set := IntSet{[]uint64{3}}
	assert.Equal(t, 2, set.Len())

	set2 := IntSet{[]uint64{4}}
	assert.Equal(t, 1, set2.Len())

	set3 := IntSet{[]uint64{2, 2}}
	assert.Equal(t, 2, set3.Len())
}

func TestIntSet_Remove_ExistentElement(t *testing.T) {
	set := IntSet{[]uint64{3}} // {0,1}
	assert.Equal(t, set.Remove(0), true)
	assert.Equal(t, []uint64{2}, set.words)
}

func TestIntSet_Remove_NotExistentElement(t *testing.T) {
	set := IntSet{[]uint64{3}} // {0,1}
	assert.Equal(t, set.Remove(10), false)
	assert.Equal(t, []uint64{3}, set.words)
}

func TestIntSet_Remove_NotExistentElementInSecondWord(t *testing.T) {
	twoIn64 := uint64(math.Pow(float64(2), float64(64)))
	words := []uint64{0, twoIn64 + twoIn64/2}
	set := IntSet{words} // {0, 126, 127}
	assert.Equal(t, set.Remove(125), false)
	assert.Equal(t, words, set.words)
}

func TestIntSet_Remove_ExistentElementInSecondWord(t *testing.T) {
	twoIn64 := uint64(math.Pow(float64(2), float64(64)))
	words := []uint64{0, twoIn64 + twoIn64/2}
	set := IntSet{words} // {126, 127}
	assert.Equal(t, set.Remove(126), true)
	assert.Equal(t, []uint64{0, twoIn64}, set.words)
}

func TestIntSet_Clear(t *testing.T) {
	set := IntSet{[]uint64{3}} // {0,1}
	set.Clear()
	assert.Equal(t, set.words, []uint64{})
}
