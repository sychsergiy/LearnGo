package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_commaRecursive(t *testing.T) {
	result := commaRecursive("10000001")
	assert.Equal(t, "10,000,001", result)
}

func Test_commaLoop(t *testing.T) {
	result := commaLoop("10000001")
	assert.Equal(t, "10,000,001", result)

	result2 := commaLoop("100")
	assert.Equal(t, "100", result2)

	result3 := commaLoop("10")
	assert.Equal(t, "10", result3)
}
