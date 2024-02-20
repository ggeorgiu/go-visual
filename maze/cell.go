package main

type CellType int

const (
	TypeCell CellType = iota
	TypeMaze
	TypeCurrent
)

type Cell struct {
	i int
	j int

	visited bool
	borders []bool
	t       CellType
}

func NewCell(i, j int) *Cell {
	c := Cell{
		i:       i,
		j:       j,
		borders: []bool{true, true, true, true},
		t:       TypeCell,
	}

	return &c
}

func (c *Cell) SetType(t CellType) {
	c.t = t
}

func (c *Cell) GetType() CellType {
	return c.t
}
