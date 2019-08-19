package dashboard

type dashboard struct {
	cellsMatrix [][]int
}

func (d dashboard) Width() int {
	return len(d.cellsMatrix[0])
}

func (d dashboard) Height() int {
	return len(d.cellsMatrix)
}

func New(width, height int) *dashboard {
	d := new(dashboard)
	for i := 0; i < height; i++ {
		d.cellsMatrix = append(d.cellsMatrix, make([]int, width))
	}
	return d
}

func (d dashboard) getCellState(x, y int) cellState {
	if d.cellsMatrix[x][y] == 0 {
		return dead
	} else if d.cellsMatrix[x][y] == 1 {
		return alive
	} else {
		panic("unresolved cell state")
	}
}

func (d dashboard) GetCell(x, y int) cell {
	x, y = d.normalizeCoordinates(x, y)
	cellState := d.getCellState(x, y)
	return cell{x, y, cellState}
}

func (d dashboard) normalizeCoordinates(x, y int) (int, int) {
	return normalizeCoordinate(x, d.Width()-1), normalizeCoordinate(y, d.Height()-1)
}

func (d dashboard) GetCellsNearby(x, y int) [8]cell {
	type cellCoordinates struct{ x, y int }
	var cells [8]cell

	cellsAroundCoordinates := [8]cellCoordinates{
		{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1},
		{x + 1, y + 1}, {x - 1, y - 1}, {x + 1, y - 1}, {x - 1, y + 1},
	}
	for index, value := range cellsAroundCoordinates {
		cells[index] = d.GetCell(value.x, value.y)
	}
	return cells
}

func normalizeCoordinate(coordinate, max int) int {
	if coordinate > max {
		return coordinate - max - 1
	} else if coordinate < 0 {
		return max + coordinate + 1
	} else {
		return coordinate
	}
}
