package main

import (
	"go-visual/pkg/ui/app"
	"time"
)

func main() {
	rows := 60
	cols := 60
	size := 600

	a := app.NewApp(
		app.WithTitle("Maze"),
		app.WithSize(size, size),
		app.WithDisplayGrid(rows, cols),
		app.WithTick(10*time.Millisecond),
	)

	alg := NewAlg(rows, cols)

	a.Run(alg)
}
