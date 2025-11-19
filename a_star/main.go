package main

import (
	"time"

	"go-visual/pkg/ui/app"
)

func main() {
	rows := 10
	cols := 10
	size := 600

	a := app.NewApp(
		app.WithTitle("A*"),
		app.WithSize(size, size),
		app.WithDisplayGrid(rows, cols),
		app.WithTick(1*time.Second),
	)

	alg := NewAlg(rows, cols)
	a.Run(alg)
}
