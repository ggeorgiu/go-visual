package display

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type pixel interface {
	fyne.CanvasObject

	update(color.Color)
	setBounds([]bool)
	get() []fyne.CanvasObject
}

type Grid struct {
	rows int
	cols int

	pixels [][]pixel
}

func NewGrid(r, c, ws int) *Grid {
	pixels := make([][]pixel, r)

	for i := 0; i < r; i++ {
		pr := make([]pixel, c)
		for j := 0; j < c; j++ {
			pr[j] = newRect(i, j, ws/r)
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
	i       int
	j       int
	c       color.Color
	borders []bool
}

func NewState(i, j int, c color.Color) *State {
	st := State{
		i: i,
		j: j,
		c: c,
	}

	return &st
}

func NewStateWithBorders(i, j int, c color.Color, b []bool) *State {
	st := State{
		i:       i,
		j:       j,
		c:       c,
		borders: b,
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
		if v.borders != nil {
			g.pixels[v.i][v.j].setBounds(v.borders)
		}
	}
}

func (g *Grid) Content() fyne.CanvasObject {
	var obj []fyne.CanvasObject

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			obj = append(obj, g.pixels[i][j].get()...)
		}
	}

	return container.NewWithoutLayout(obj...)
}
