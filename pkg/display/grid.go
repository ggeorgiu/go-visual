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
	getBounds() []fyne.CanvasObject
	get() []fyne.CanvasObject
}

type Grid struct {
	rows int
	cols int

	pixels  [][]pixel
	content *fyne.Container
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

type StateUpdater interface {
	GetInitialState() []*State
	GetUpdatedState() []*State
}

type State struct {
	i       int
	j       int
	c       color.Color
	borders []bool
}

func NewState(opts ...func(s *State)) *State {
	st := State{}

	for _, opt := range opts {
		opt(&st)
	}

	return &st
}

func WithCoords(i, j int) func(*State) {
	return func(s *State) {
		s.i = i
		s.j = j
	}
}

func WithColor(c color.Color) func(*State) {
	return func(s *State) {
		s.c = c
	}
}

func WithBorders(b []bool) func(*State) {
	return func(s *State) {
		s.borders = b
	}
}

func (g *Grid) SetState(sp StateUpdater) {
	for _, s := range sp.GetInitialState() {
		g.pixels[s.i][s.j].update(s.c)
	}

}

func (g *Grid) UpdateState(sp StateUpdater) {
	for _, v := range sp.GetUpdatedState() {
		for _, b := range g.pixels[v.i][v.j].getBounds() {
			g.content.Remove(b)
		}

		g.pixels[v.i][v.j].update(v.c)
		if v.borders != nil {
			g.pixels[v.i][v.j].setBounds(v.borders)
		}

		for _, b := range g.pixels[v.i][v.j].getBounds() {
			g.content.Add(b)
		}
	}

	g.content.Refresh()
}

func (g *Grid) Content() fyne.CanvasObject {
	var obj []fyne.CanvasObject

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			obj = append(obj, g.pixels[i][j].get()...)
		}
	}

	g.content = container.NewWithoutLayout(obj...)
	return g.content
}
