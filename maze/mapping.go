package main

import (
	"go-visual/pkg/display"
	"image/color"
)

var typeToColorMap = map[CellType]color.Color{
	TypeCell:    color.White,
	TypeMaze:    color.RGBA{R: 0, G: 255, B: 0, A: 100},   // gree
	TypeCurrent: color.RGBA{R: 255, G: 255, B: 0, A: 255}, // yellow
}

func getColor(c *Cell) color.Color {
	return typeToColorMap[c.t]
}

func ToDisplayState(a *Alg) [][]color.Color {
	st := make([][]color.Color, a.rows)
	for i := 0; i < len(a.state)-1; i++ {

		sr := make([]color.Color, a.cols)
		for j := 0; j < len(a.state[0])-1; j++ {
			sr[j] = getColor(a.state[i][j])
		}
		st[i] = sr
	}

	return st
}

func ToUpdatedState(a *Alg) []*display.State {
	var st []*display.State

	c := a.changed
	c.SetType(TypeMaze)
	st = append(st, display.NewStateWithBorders(c.i, c.j, getColor(c), c.borders))

	c = a.current
	c.SetType(TypeCurrent)
	st = append(st, display.NewStateWithBorders(c.i, c.j, getColor(c), c.borders))

	return st
}
