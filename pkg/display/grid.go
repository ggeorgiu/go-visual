package display

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type pixel interface {
	fyne.CanvasObject

	update(c color.Color)
	get() fyne.CanvasObject
}

type Grid struct {
	rows int
	cols int

	pixels [][]pixel
}

func NewGrid(r, c int) *Grid {
	pixels := make([][]pixel, r)

	for i := 0; i < r; i++ {
		pr := make([]pixel, c)
		for j := 0; j < c; j++ {
			pr[j] = newRect()
		}

		pixels[i] = pr
	}

	return &Grid{
		rows:   r,
		cols:   c,
		pixels: pixels,
	}
}

type State struct {
	i int
	j int
	c color.Color
}

func NewState(i, j int, c color.Color) *State {
	st := State{
		i: i,
		j: j,
		c: c,
	}

	return &st
}

func (g *Grid) SetState(state [][]color.Color) {
	for i := 0; i < len(state)-1; i++ {
		for j := 0; j < len(state[0])-1; j++ {
			g.pixels[i][j].update(state[i][j])
		}
	}
}

func (g *Grid) UpdateState(state []*State) {
	for _, v := range state {
		g.pixels[v.i][v.j].update(v.c)
	}
}

func (g *Grid) Content() fyne.CanvasObject {
	cont := make([]fyne.CanvasObject, g.rows)

	for i := 0; i < g.rows; i++ {
		obj := make([]fyne.CanvasObject, g.cols)

		for j := 0; j < g.cols; j++ {
			obj[j] = g.pixels[i][j].get()
		}

		c := container.New(layout.NewGridLayout(g.cols), obj...)
		cont[i] = c
	}

	return container.New(layout.NewGridLayoutWithRows(g.rows), cont...)
}
