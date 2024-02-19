package main

import (
	"go-visual/pkg/display"
	"image/color"
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

	for i := 0; i < len(a.closedSet); i++ {
		c := a.closedSet[i]

		c.SetType(TypeChecked)
		st = append(st, display.NewState(c.i, c.j, getColor(c)))
	}

	for i := 0; i < len(a.openSet); i++ {
		c := a.openSet[i]

		c.SetType(TypePossible)
		st = append(st, display.NewState(c.i, c.j, getColor(c)))
	}

	for i := 0; i < len(a.path); i++ {
		c := a.path[i]

		c.SetType(TypePath)
		st = append(st, display.NewState(c.i, c.j, getColor(c)))
	}

	a.start.SetType(TypeStart)
	a.end.SetType(TypeEnd)

	st = append(st,
		display.NewState(a.start.i, a.start.j, getColor(a.start)),
		display.NewState(a.end.i, a.end.j, getColor(a.end)),
	)

	return st
}
