package strings

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_createLetterCountMap(t *testing.T) {
	result := createLettersCountMap("test")
	expected := map[int32]int{
		116: 2, 101: 1, 115: 1,
	}
	assert.True(t, reflect.DeepEqual(result, expected))
}

func Test_isAnagrams(t *testing.T) {
	assert.True(t, isAnagrams("test", "test"))
	assert.True(t, isAnagrams("test1", "estt1"))
	assert.False(t, isAnagrams("test", "tests"))
}
