package main

import (
	"time"

	"go-visual/pkg/ui/app"
)

func main() {
	rows := 100
	cols := 100
	size := 600

	a := app.NewApp(
		app.WithTitle("A*"),
		app.WithSize(size, size),
		app.WithDisplayGrid(rows, cols),
		app.WithTick(50*time.Millisecond),
	)

	alg := NewAlg(rows, cols)
	a.Run(alg)
}
