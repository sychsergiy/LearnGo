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
	assert.True(t, set.Remove(0))
	assert.Equal(t, []uint64{2}, set.words)
}

func TestIntSet_Remove_NotExistentElement(t *testing.T) {
	set := IntSet{[]uint64{3}} // {0,1}
	assert.False(t, set.Remove(10))
	assert.Equal(t, []uint64{3}, set.words)
}

func TestIntSet_Remove_NotExistentElementInSecondWord(t *testing.T) {
	twoIn64 := uint64(math.Pow(float64(2), float64(64)))
	words := []uint64{0, twoIn64 + twoIn64/2}
	set := IntSet{words} // {0, 126, 127}
	assert.False(t, set.Remove(125))
	assert.Equal(t, words, set.words)
}

func TestIntSet_Remove_ExistentElementInSecondWord(t *testing.T) {
	twoIn64 := uint64(math.Pow(float64(2), float64(64)))
	words := []uint64{0, twoIn64 + twoIn64/2}
	set := IntSet{words} // {126, 127}
	assert.True(t, set.Remove(126))
	assert.Equal(t, []uint64{0, twoIn64}, set.words)
}

func TestIntSet_Clear(t *testing.T) {
	set := IntSet{[]uint64{3}} // {0,1}
	set.Clear()
	assert.Equal(t, set.words, []uint64{})
}

func TestIntSet_Copy(t *testing.T) {
	set := &IntSet{[]uint64{3}} // {0,1}
	setCopy := set.Copy()

	assert.False(t, &set.words == &setCopy.words)
	assert.False(t, &set == &setCopy)
	assert.Equal(t, set.words, setCopy.words)
}

func TestIntSet_AddAll(t *testing.T) {
	set := &IntSet{}
	assert.Equal(t, 3, set.AddAll(1, 2, 3))
	assert.Equal(t, set.words, []uint64{14})
}

func TestIntSet_AddAll_Intersection(t *testing.T) {
	twoIn64 := uint64(math.Pow(float64(2), float64(64)))
	set := &IntSet{[]uint64{2}} // {1}
	assert.Equal(t, set.AddAll(1, 127), 1)
	assert.Equal(t, set.words, []uint64{2, twoIn64})
}

func TestIntSet_Add_New(t *testing.T) {
	set := IntSet{}
	assert.True(t, set.Add(1))
	assert.Equal(t, []uint64{2}, set.words)
}

func TestIntSet_Add_Old(t *testing.T) {
	set := IntSet{[]uint64{2}}
	assert.False(t, set.Add(1))
	assert.Equal(t, []uint64{2}, set.words)
}

func TestIntSet_Elems(t *testing.T) {
	set := IntSet{[]uint64{3, 0, 1}}
	assert.Equal(t, []int{0, 1, 128}, set.Elems())
}

func TestIntSet_String(t *testing.T) {
	set := IntSet{[]uint64{3, 0, 1}}
	assert.Equal(t, "{0, 1, 128}", set.String())
}
