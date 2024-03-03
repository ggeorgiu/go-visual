package main

import (
	"math/rand"

	"go-visual/pkg/ui/display"
)

type Alg struct {
	rows int
	cols int

	state   [][]*Cell
	changed *Cell
	current *Cell
	stack   []*Cell

	ended bool
}

func NewAlg(rows, cols int) *Alg {
	a := Alg{
		rows:  rows,
		cols:  cols,
		state: initGrid(rows, cols),
	}

	a.current = a.state[0][0]
	a.current.visited = true
	a.stack = append(a.stack, a.current)

	return &a
}

func initGrid(rows, cols int) [][]*Cell {
	st := make([][]*Cell, rows)
	for i := 0; i < rows; i++ {
		st[i] = make([]*Cell, cols)
		for j := 0; j < cols; j++ {
			st[i][j] = NewCell(i, j)
		}
	}

	return st
}

func (a *Alg) Step() {
	n := a.getNeighs(a.current.i, a.current.j)
	if len(n) == 0 {
		if len(a.stack) == 0 {
			a.ended = true
			return
		}

		a.changed = a.current
		a.current = a.stack[len(a.stack)-1]
		a.stack = a.stack[:len(a.stack)-1]

		return
	}

	r := rand.Intn(len(n))

	a.stack = append(a.stack, a.current)
	next := n[r]
	changeBorders(a.current, next)
	a.changed = a.current
	a.current = next
	a.current.visited = true
}

func changeBorders(current *Cell, next *Cell) {
	x := next.i - current.i
	if x == 1 {
		current.borders[2] = false
		next.borders[0] = false

	} else if x == -1 {
		current.borders[0] = false
		next.borders[2] = false
	}

	y := next.j - current.j
	if y == 1 {
		current.borders[1] = false
		next.borders[3] = false

	} else if y == -1 {
		current.borders[3] = false
		next.borders[1] = false
	}
}

type offset struct {
	i int
	j int
}

var neighs = []offset{
	{-1, 0},
	{0, -1},
	{0, 1},
	{1, 0},
}

func (a *Alg) getNeighs(i, j int) []*Cell {
	var neigh []*Cell
	for _, v := range neighs {
		ni := i + v.i
		nj := j + v.j

		if ni < 0 || ni > a.rows-1 || nj < 0 || nj > a.cols-1 || a.state[ni][nj].visited == true {
			continue
		}

		neigh = append(neigh, a.state[ni][nj])
	}

	return neigh
}

func (a *Alg) GetInitialState() []*display.State {
	var st []*display.State
	for i := 0; i < len(a.state)-1; i++ {
		for j := 0; j < len(a.state[0])-1; j++ {
			st = append(st, display.NewState(display.WithCoords(i, j), display.WithColor(a.state[i][j].GetColor())))
		}
	}

	return st
}

func (a *Alg) GetUpdatedState() []*display.State {
	var st []*display.State

	c := a.changed
	c.SetType(TypeMaze)
	st = append(st, display.NewState(display.WithCoords(c.i, c.j), display.WithColor(c.GetColor()), display.WithBorders(c.borders)))

	c = a.current
	c.SetType(TypeCurrent)
	st = append(st, display.NewState(display.WithCoords(c.i, c.j), display.WithColor(c.GetColor()), display.WithBorders(c.borders)))

	return st
}

func (a *Alg) Ended() bool {
	return a.ended
}
