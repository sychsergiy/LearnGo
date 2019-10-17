package gopl

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_countWordsFrequency(t *testing.T) {
	result := countWordsFrequency("test asdf test 15 15 abc")
	assert.True(t, reflect.DeepEqual(result, map[string]int{"test": 2, "asdf": 1, "15": 2, "abc": 1}))
}
