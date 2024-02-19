package main

import (
	"go-visual/pkg/display"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var tick = 50 * time.Millisecond

func main() {
	rows := 50
	cols := 50

	myApp := app.New()
	w := myApp.NewWindow("A*")
	w.Resize(fyne.NewSize(600, 600))

	disp := display.NewGrid(rows, cols)
	w.SetContent(disp.Content())

	alg := NewAlg(rows, cols)
	disp.SetState(ToDisplayState(alg))

	go runApp(alg, disp)

	w.ShowAndRun()
}

func runApp(a *Alg, g *display.Grid) {
	for range time.Tick(tick) {
		if a.ended {
			break
		}

		a.Step()
		g.UpdateState(ToUpdatedState(a))
	}
}
