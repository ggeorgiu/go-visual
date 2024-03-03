package main

import (
	"go-visual/pkg/ui/app"
)

func main() {
	rows := 20
	cols := 20
	size := float32(600)

	a := app.NewApp(
		app.WithTitle("A*"),
		app.WithSize(size, size),
		app.WithDisplayGrid(rows, cols),
	)

	alg := NewAlg(rows, cols)
	a.Run(alg)
}
