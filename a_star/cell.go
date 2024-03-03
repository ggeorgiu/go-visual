package main

import (
	"image/color"
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

var typeToColorMap = map[CellType]color.Color{
	TypeCell:     color.White,
	TypeWall:     color.Black,
	TypePossible: color.RGBA{R: 0, G: 255, B: 0, A: 255},   // green
	TypeStart:    color.RGBA{R: 255, G: 255, B: 0, A: 255}, // yellow
	TypeEnd:      color.RGBA{R: 255, G: 255, B: 0, A: 255}, // yellow
	TypeChecked:  color.RGBA{R: 255, G: 0, B: 0, A: 255},   // red
	TypePath:     color.RGBA{R: 0, G: 0, B: 255, A: 255},   // blue
}

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

func (c *Cell) GetColor() color.Color {
	return typeToColorMap[c.t]
}
