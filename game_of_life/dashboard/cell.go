package dashboard

type cellState int

const (
	alive cellState = iota
	dead
)

type cell struct {
	x, y int
	cellState
}

func (c cell) isAlive() bool {
	return c.cellState == alive
}

func (c cell) isDead() bool {
	return c.cellState == dead
}
