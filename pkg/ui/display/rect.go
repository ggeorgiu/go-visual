package display

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

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

func (e *rect) setColor(c color.Color) {
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
