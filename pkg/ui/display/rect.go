package display

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type rect struct {
	rectangle *canvas.Rectangle
	strokes   []fyne.CanvasObject
}

func newRect(row, col, width int) *rect {
	e := rect{
		rectangle: &canvas.Rectangle{
			FillColor: color.White,
		},
	}
	e.rectangle.Move(fyne.Position{
		X: float32(col * width),
		Y: float32(row * width),
	})

	e.rectangle.Resize(fyne.Size{
		Width:  float32(width),
		Height: float32(width),
	})

	e.setBounds([]bool{true, true, true, true})

	return &e
}

func (e *rect) setColor(c color.Color) {
	e.rectangle.FillColor = c
	e.rectangle.Refresh()
}

func (e *rect) get() []fyne.CanvasObject {
	var all []fyne.CanvasObject

	all = append(all, e.rectangle)
	all = append(all, e.strokes...)

	return all
}

func (e *rect) setBounds(b []bool) {
	e.strokes = []fyne.CanvasObject{}

	if b[0] == true {
		from := fyne.NewPos(e.rectangle.Position().X, e.rectangle.Position().Y)
		to := fyne.NewPos(e.rectangle.Position().X+e.rectangle.Size().Width, e.rectangle.Position().Y)

		l := canvas.NewLine(color.Black)
		l.Position1 = to
		l.Position2 = from
		l.StrokeWidth = 2
		e.strokes = append(e.strokes, l)
	}

	if b[1] == true {
		from := fyne.NewPos(e.rectangle.Position().X+e.rectangle.Size().Width, e.rectangle.Position().Y)
		to := fyne.NewPos(e.rectangle.Position().X+e.rectangle.Size().Width, e.rectangle.Position().Y+e.rectangle.Size().Width)

		l := canvas.NewLine(color.Black)
		l.Position1 = to
		l.Position2 = from
		l.StrokeWidth = 2
		e.strokes = append(e.strokes, l)
	}

	if b[2] == true {
		from := fyne.NewPos(e.rectangle.Position().X+e.rectangle.Size().Width, e.rectangle.Position().Y+e.rectangle.Size().Width)
		to := fyne.NewPos(e.rectangle.Position().X, e.rectangle.Position().Y+e.rectangle.Size().Width)

		l := canvas.NewLine(color.Black)
		l.Position1 = from
		l.Position2 = to
		l.StrokeWidth = 2
		e.strokes = append(e.strokes, l)

	}

	if b[3] == true {
		from := fyne.NewPos(e.rectangle.Position().X, e.rectangle.Position().Y+e.rectangle.Size().Width)
		to := fyne.NewPos(e.rectangle.Position().X, e.rectangle.Position().Y)

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
