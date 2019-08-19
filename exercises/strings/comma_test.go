package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_commaRecursive(t *testing.T) {
	result := commaRecursive("10000001")
	assert.Equal(t, "10,000,001", result)
}

func Test_insertCommasToInteger(t *testing.T) {
	result := insertCommasToIntegerPart("10000001", 3)
	assert.Equal(t, "10,000,001", result)

	result2 := insertCommasToIntegerPart("100", 3)
	assert.Equal(t, "100", result2)

	result3 := insertCommasToIntegerPart("10", 3)
	assert.Equal(t, "10", result3)
}

func Test_insertCommasToFloat(t *testing.T) {
	result := insertCommasToFloatPart("10000001", 3)
	assert.Equal(t, "100,000,01", result)

	result2 := insertCommasToFloatPart("100", 3)
	assert.Equal(t, "100", result2)

	result3 := insertCommasToFloatPart("10", 3)
	assert.Equal(t, "10", result3)
}

func Test_insertCommas(t *testing.T) {
	result := insertCommas("1000.0001", 3)
	assert.Equal(t, "1,000.000,1", result)

	result2 := insertCommas("100.01", 2)
	assert.Equal(t, "1,00.01", result2)

	result3 := insertCommas("1000", 2)
	assert.Equal(t, "10,00", result3)
}
