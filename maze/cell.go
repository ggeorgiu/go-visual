package main

import "image/color"

type CellType int

const (
	TypeCell CellType = iota
	TypeMaze
	TypeCurrent
)

var typeToColorMap = map[CellType]color.Color{
	TypeCell:    color.White,
	TypeMaze:    color.RGBA{R: 0, G: 255, B: 0, A: 100},   // green
	TypeCurrent: color.RGBA{R: 255, G: 255, B: 0, A: 255}, // yellow
}

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
	}

	return &c
}

func (c *Cell) SetType(t CellType) {
	c.t = t
}

func (c *Cell) GetType() CellType {
	return c.t
}

func (c *Cell) GetColor() color.Color {
	return typeToColorMap[c.t]
}
