package main

import (
	"math"
)

type CellType int

const (
	TypeCell CellType = iota
	TypeWall
	TypePath
	TypeStart
	TypeEnd
	TypePossible
	TypeChecked
)

type Cell struct {
	i int
	j int

	g int
	f float64
	h float64

	t CellType

	prev *Cell
}

func NewCell(i, j int) *Cell {
	c := Cell{
		i: i,
		j: j,
		g: math.MaxInt,
		f: math.MaxInt,
	}

	return &c
}

func (c *Cell) SetType(t CellType) {
	c.t = t
}

func (c *Cell) GetType() CellType {
	return c.t
}

func (c *Cell) IsWall() bool {
	return c.t == TypeWall
}
