package display

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type rect struct {
	*canvas.Rectangle
}

func newRect() *rect {
	e := rect{
		canvas.NewRectangle(color.White),
	}

	return &e
}

func (e *rect) update(c color.Color) {
	e.FillColor = c
	e.Refresh()
}

func (e *rect) get() fyne.CanvasObject {
	return e.Rectangle
}
