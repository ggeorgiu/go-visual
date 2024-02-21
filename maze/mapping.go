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

type mapper struct {
	a *Alg
}

func (m *mapper) GetInitialState() []*display.State {
	var st []*display.State
	for i := 0; i < len(m.a.state)-1; i++ {
		for j := 0; j < len(m.a.state[0])-1; j++ {
			st = append(st, display.NewState(display.WithCoords(i, j), display.WithColor(getColor(m.a.state[i][j]))))
		}
	}

	return st
}

func (m *mapper) GetUpdatedState() []*display.State {
	var st []*display.State

	c := m.a.changed
	c.SetType(TypeMaze)
	st = append(st, display.NewState(display.WithCoords(c.i, c.j), display.WithColor(getColor(c)), display.WithBorders(c.borders)))

	c = m.a.current
	c.SetType(TypeCurrent)
	st = append(st, display.NewState(display.WithCoords(c.i, c.j), display.WithColor(getColor(c)), display.WithBorders(c.borders)))

	return st
}
