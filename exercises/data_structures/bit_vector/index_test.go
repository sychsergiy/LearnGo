package bit_vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSet_Len(t *testing.T) {
	set := IntSet{[]uint64{3}}
	assert.Equal(t, 2, set.Len())

	set2 := IntSet{[]uint64{4}}
	assert.Equal(t, 1, set2.Len())

	set3 := IntSet{[]uint64{2,2}}
	assert.Equal(t, 2, set3.Len())
}
