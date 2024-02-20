package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"go-visual/pkg/display"
	"time"
)

var tick = 500 * time.Millisecond

func main() {
	rows := 10
	cols := 10

	myApp := app.New()
	w := myApp.NewWindow("Maze")
	w.Resize(fyne.NewSize(600, 600))

	disp := display.NewGrid(rows, cols, 600)
	w.SetContent(disp.Content())

	alg := NewAlg(rows, cols)
	go runApp(alg, w, disp)

	w.ShowAndRun()
}

func runApp(a *Alg, w fyne.Window, g *display.Grid) {
	for range time.Tick(tick) {
		if a.ended {
			break
		}

		a.Step()
		g.UpdateState(ToUpdatedState(a))
		w.SetContent(g.Content())
	}
}
