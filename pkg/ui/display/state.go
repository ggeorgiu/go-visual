package display

import "image/color"

type StateProvider interface {
	GetInitialState() []*CoordState
	GetUpdatedState() []*CoordState
	Step()
	Ended() bool
}

type CoordState struct {
	i       int
	j       int
	c       color.Color
	borders []bool
}

func NewCoordState(opts ...func(s *CoordState)) *CoordState {
	st := CoordState{}

	for _, opt := range opts {
		opt(&st)
	}

	return &st
}

func WithCoords(i, j int) func(*CoordState) {
	return func(s *CoordState) {
		s.i = i
		s.j = j
	}
}

func WithColor(c color.Color) func(*CoordState) {
	return func(s *CoordState) {
		s.c = c
	}
}

func WithBorders(b []bool) func(*CoordState) {
	return func(s *CoordState) {
		s.borders = b
	}
}
