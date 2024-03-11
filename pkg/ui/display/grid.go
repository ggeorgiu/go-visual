package display

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type pixel interface {
	get() []fyne.CanvasObject

	setColor(color.Color)
	setBounds([]bool)
	getBounds() []fyne.CanvasObject
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

func (g *Grid) InitState(sp StateProvider) {
	for _, s := range sp.GetInitialState() {
		g.pixels[s.i][s.j].setColor(s.c)
	}

}

func (g *Grid) UpdateState(sp StateProvider) {
	for _, v := range sp.GetUpdatedState() {
		for _, b := range g.pixels[v.i][v.j].getBounds() {
			g.content.Remove(b)
		}

		g.pixels[v.i][v.j].setColor(v.c)
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
