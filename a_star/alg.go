package main

import (
	"fmt"
	"math"
	"math/rand/v2"

	"go-visual/pkg/ui/display"
)

var (
	heuristicAbs = func(c, e *Cell) float64 {
		return math.Abs(float64(c.i-e.i)) + math.Abs(float64(c.j-e.j))
	}
	heuristicSqrt = func(c, e *Cell) float64 {
		return math.Sqrt(float64((c.i-e.i)*(c.i-e.i) + (c.j-e.j)*(c.j-e.j)))
	}
)

type Alg struct {
	rows int
	cols int

	start   *Cell
	current *Cell
	end     *Cell

	state     [][]*Cell
	openSet   []*Cell
	closedSet []*Cell
	path      []*Cell

	ended bool
}

func NewAlg(rows, cols int) *Alg {
	a := Alg{
		rows:  rows,
		cols:  cols,
		state: initGrid(rows, cols),
	}

	a.start = a.state[0][0]
	a.end = a.state[rows-1][cols-1]

	a.start.g = 0
	a.start.h = heuristicAbs(a.start, a.end)
	a.openSet = append(a.openSet, a.start)

	return &a
}

func (a *Alg) Step() {
	if len(a.openSet) > 0 {
		winner := 0
		for i, v := range a.openSet {
			if v.f < a.openSet[winner].f {
				winner = i
			}
		}

		a.current = a.openSet[winner]
		current := a.current

		if current == a.end {
			a.ended = true
			a.path = []*Cell{}
			a.path = append(a.path, current)
			for current.prev != nil {
				a.path = append(a.path, current.prev)
				current = current.prev
			}
			return
		}

		a.openSet = append(a.openSet[:winner], a.openSet[winner+1:]...)
		a.closedSet = append(a.closedSet, current)

		neigs := a.getNeighs(current.i, current.j)
		for _, n := range neigs {
			isClosed := false
			for _, cs := range a.closedSet {
				if cs == n {
					isClosed = true
					break
				}
			}

			if isClosed {
				continue
			}

			currentG := current.g + 1

			isOpen := false
			for _, os := range a.openSet {
				if os == n {
					isOpen = true
					break
				}
			}

			if !isOpen {
				a.openSet = append(a.openSet, n)
			}

			if currentG > n.g {
				continue
			}

			n.g = currentG
			n.h = heuristicAbs(n, a.end)
			n.f = float64(n.g) + n.h

			n.prev = current
		}
	} else {
		fmt.Println("NO SOLUTION")
		a.ended = true
	}
}

func initGrid(rows, cols int) [][]*Cell {
	st := make([][]*Cell, rows)
	for i := 0; i < rows; i++ {
		st[i] = make([]*Cell, cols)
		for j := 0; j < cols; j++ {
			st[i][j] = NewCell(i, j)

			r := rand.Float64()
			if r < 0.4 {
				st[i][j].SetType(TypeWall)
			}
		}
	}

	return st
}

type offset struct {
	i int
	j int
}

var neighs = []offset{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func (a *Alg) getNeighs(i, j int) []*Cell {
	var neigh []*Cell
	for _, v := range neighs {
		ni := i + v.i
		nj := j + v.j

		if ni < 0 || ni > a.rows-1 || nj < 0 || nj > a.cols-1 || a.state[ni][nj].IsWall() {
			continue
		}

		neigh = append(neigh, a.state[ni][nj])
	}

	return neigh
}

func (a *Alg) GetInitialState() []*display.CoordState {
	var st []*display.CoordState
	for i := 0; i < len(a.state); i++ {
		for j := 0; j < len(a.state[0]); j++ {
			st = append(st, display.NewCoordState(display.WithCoords(i, j), display.WithColor(a.state[i][j].GetColor())))
		}
	}

	return st
}

func (a *Alg) GetUpdatedState() []*display.CoordState {
	var st []*display.CoordState

	for i := 0; i < len(a.closedSet); i++ {
		c := a.closedSet[i]

		c.SetType(TypeChecked)
		st = append(st, display.NewCoordState(display.WithCoords(c.i, c.j), display.WithColor(c.GetColor())))
	}

	for i := 0; i < len(a.openSet); i++ {
		c := a.openSet[i]

		c.SetType(TypePossible)
		st = append(st, display.NewCoordState(display.WithCoords(c.i, c.j), display.WithColor(c.GetColor())))
	}

	for i := 0; i < len(a.path); i++ {
		c := a.path[i]

		c.SetType(TypePath)
		st = append(st, display.NewCoordState(display.WithCoords(c.i, c.j), display.WithColor(c.GetColor())))
	}

	a.start.SetType(TypeStart)
	a.end.SetType(TypeEnd)

	st = append(st,
		display.NewCoordState(display.WithCoords(a.start.i, a.start.j), display.WithColor(a.start.GetColor())),
		display.NewCoordState(display.WithCoords(a.end.i, a.end.j), display.WithColor(a.end.GetColor())),
	)

	return st
}

func (a *Alg) Ended() bool {
	return a.ended
}
