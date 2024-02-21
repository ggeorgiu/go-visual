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

	for i := 0; i < len(m.a.closedSet); i++ {
		c := m.a.closedSet[i]

		c.SetType(TypeChecked)
		st = append(st, display.NewState(display.WithCoords(c.i, c.j), display.WithColor(getColor(c))))
	}

	for i := 0; i < len(m.a.openSet); i++ {
		c := m.a.openSet[i]

		c.SetType(TypePossible)
		st = append(st, display.NewState(display.WithCoords(c.i, c.j), display.WithColor(getColor(c))))
	}

	for i := 0; i < len(m.a.path); i++ {
		c := m.a.path[i]

		c.SetType(TypePath)
		st = append(st, display.NewState(display.WithCoords(c.i, c.j), display.WithColor(getColor(c))))
	}

	m.a.start.SetType(TypeStart)
	m.a.end.SetType(TypeEnd)

	st = append(st,
		display.NewState(display.WithCoords(m.a.start.i, m.a.start.j), display.WithColor(getColor(m.a.start))),
		display.NewState(display.WithCoords(m.a.end.i, m.a.end.j), display.WithColor(getColor(m.a.end))),
	)

	return st
}
