package display

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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

func (g *Grid) InitState(sp StateProvider) {
	for _, s := range sp.GetInitialState() {
		g.pixels[s.i][s.j].update(s.c)
	}

}

func (g *Grid) UpdateState(sp StateProvider) {
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

// TODO refactor
type rect struct {
	*canvas.Rectangle
	strokes []fyne.CanvasObject
}

func newRect(row, col, width int) *rect {
	e := rect{
		Rectangle: &canvas.Rectangle{
			FillColor: color.White,
		},
	}
	e.Move(fyne.Position{
		X: float32(col * width),
		Y: float32(row * width),
	})

	e.Resize(fyne.Size{
		Width:  float32(width),
		Height: float32(width),
	})

	e.setBounds([]bool{true, true, true, true})

	return &e
}

func (e *rect) update(c color.Color) {
	e.FillColor = c
	e.Refresh()
}

func (e *rect) get() []fyne.CanvasObject {
	var all []fyne.CanvasObject

	all = append(all, e.Rectangle)
	all = append(all, e.strokes...)

	return all
}

func (e *rect) setBounds(b []bool) {
	e.strokes = []fyne.CanvasObject{}

	if b[0] == true {
		from := fyne.NewPos(e.Rectangle.Position().X, e.Rectangle.Position().Y)
		to := fyne.NewPos(e.Rectangle.Position().X+e.Rectangle.Size().Width, e.Rectangle.Position().Y)

		l := canvas.NewLine(color.Black)
		l.Position1 = to
		l.Position2 = from
		l.StrokeWidth = 2
		e.strokes = append(e.strokes, l)
	}

	if b[1] == true {
		from := fyne.NewPos(e.Rectangle.Position().X+e.Rectangle.Size().Width, e.Rectangle.Position().Y)
		to := fyne.NewPos(e.Rectangle.Position().X+e.Rectangle.Size().Width, e.Rectangle.Position().Y+e.Rectangle.Size().Width)

		l := canvas.NewLine(color.Black)
		l.Position1 = to
		l.Position2 = from
		l.StrokeWidth = 2
		e.strokes = append(e.strokes, l)
	}

	if b[2] == true {
		from := fyne.NewPos(e.Rectangle.Position().X+e.Rectangle.Size().Width, e.Rectangle.Position().Y+e.Rectangle.Size().Width)
		to := fyne.NewPos(e.Rectangle.Position().X, e.Rectangle.Position().Y+e.Rectangle.Size().Width)

		l := canvas.NewLine(color.Black)
		l.Position1 = from
		l.Position2 = to
		l.StrokeWidth = 2
		e.strokes = append(e.strokes, l)

	}

	if b[3] == true {
		from := fyne.NewPos(e.Rectangle.Position().X, e.Rectangle.Position().Y+e.Rectangle.Size().Width)
		to := fyne.NewPos(e.Rectangle.Position().X, e.Rectangle.Position().Y)

		l := canvas.NewLine(color.Black)
		l.Position1 = from
		l.Position2 = to
		l.StrokeWidth = 2
		e.strokes = append(e.strokes, l)
	}
}

func (e *rect) getBounds() []fyne.CanvasObject {
	return e.strokes
}
