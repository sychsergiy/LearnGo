package dashboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDashboard_New(t *testing.T) {
	d := New(1, 2)
	expected := [][]int{{0}, {0}}
	assert.Equal(t, expected, d.cellsMatrix)
}

func TestDashboard_getCellState(t *testing.T) {
	d := new(dashboard)
	d.cellsMatrix = [][]int{{0, 1}}
	assert.Equal(t, dead, d.getCellState(0, 0))
	assert.Equal(t, alive, d.getCellState(0, 1))
}

func TestDashboard_GetCellsNearby(t *testing.T) {
	d := new(dashboard)
	d.cellsMatrix = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	result := d.GetCellsNearby(0, 0)
	expected := [8]cell{
		{1, 0, dead}, {2, 0, dead}, {0, 1, dead}, {0, 2, dead},
		{1, 1, dead}, {2, 2, dead}, {1, 2, dead}, {2, 1, dead},
	}
	assert.Equal(t, expected, result)
}

func TestDashboard_normalizeCoordinates(t *testing.T) {
	d := new(dashboard)
	d.cellsMatrix = [][]int{{0, 0, 0}, {0, 0, 0}}
	x, y := d.normalizeCoordinates(4, 3)
	assert.Equal(t, 1, x)
	assert.Equal(t, 1, y)
}

func TestDashboard_Height(t *testing.T) {
	d := new(dashboard)
	d.cellsMatrix = [][]int{{1, 1, 1}, {1, 1, 1}}
	assert.Equal(t, d.Width(), 3)
}

func TestDashboard_Width(t *testing.T) {
	d := new(dashboard)
	d.cellsMatrix = [][]int{{1, 2, 3}, {1, 2, 3}}
	assert.Equal(t, d.Height(), 2)
}

func TestDashboard_normalizeCoordinate(t *testing.T) {
	var result int
	result = normalizeCoordinate(4, 3)
	assert.Equal(t, 0, result)
	result = normalizeCoordinate(-1, 2)
	assert.Equal(t, 2, result)
}

func TestDashboard_GetCell(t *testing.T) {
	d := new(dashboard)
	d.cellsMatrix = [][]int{{1, 1}, {1, 1}}
	result := d.GetCell(1, 1)
	assert.Equal(t, result, cell{1, 1, alive})
}
