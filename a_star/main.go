package main

import (
	"go-visual/pkg/display"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var tick = 500 * time.Millisecond

func main() {
	rows := 20
	cols := 20
	size := 600

	myApp := app.New()
	w := myApp.NewWindow("A*")
	w.Resize(fyne.NewSize(float32(size), float32(size)))

	disp := display.NewGrid(rows, cols, size)
	w.SetContent(disp.Content())

	alg := NewAlg(rows, cols)
	m := mapper{a: alg}
	disp.SetState(&m)
	go runApp(&m, alg, disp)

	w.ShowAndRun()
}

func runApp(m *mapper, a *Alg, g *display.Grid) {
	for range time.Tick(tick) {
		if a.ended {
			break
		}

		a.Step()
		g.UpdateState(m)
	}
}
