package display

import "image/color"

type StateProvider interface {
	GetInitialState() []*State
	GetUpdatedState() []*State
	Step()
	Ended() bool
}

type State struct {
	i       int
	j       int
	c       color.Color
	borders []bool
}

func NewState(opts ...func(s *State)) *State {
	st := State{}

	for _, opt := range opts {
		opt(&st)
	}

	return &st
}

func WithCoords(i, j int) func(*State) {
	return func(s *State) {
		s.i = i
		s.j = j
	}
}

func WithColor(c color.Color) func(*State) {
	return func(s *State) {
		s.c = c
	}
}

func WithBorders(b []bool) func(*State) {
	return func(s *State) {
		s.borders = b
	}
}
