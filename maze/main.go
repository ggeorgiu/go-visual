package main

import (
	"time"

	"go-visual/pkg/ui/app"
)

func main() {
	rows := 30
	cols := 30
	size := 600

	a := app.NewApp(
		app.WithTitle("Maze"),
		app.WithSize(size, size),
		app.WithDisplayGrid(rows, cols),
		app.WithTick(100*time.Millisecond),
	)

	alg := NewAlg(rows, cols)

	a.Run(alg)
}
