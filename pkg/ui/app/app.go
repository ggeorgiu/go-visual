package app

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"go-visual/pkg/ui/display"
)

const (
	tick     = 100 * time.Millisecond
	defaultH = float32(200)
	defaultW = float32(200)
)

type StateManager interface {
	InitState(sp display.StateProvider)
	UpdateState(sp display.StateProvider)
}

type App struct {
	app    fyne.App
	window fyne.Window
	height float32
	width  float32
	tick   time.Duration

	sm StateManager
}

func NewApp(opts ...func(a *App)) *App {
	newApp := app.New()

	a := App{
		app:    newApp,
		tick:   tick,
		height: defaultH,
		width:  defaultW,
	}

	for _, opt := range opts {
		opt(&a)
	}

	return &a
}

func WithTitle(title string) func(*App) {
	return func(a *App) {
		w := a.app.NewWindow(title)
		a.window = w
	}
}

func WithSize(w, h float32) func(*App) {
	return func(a *App) {
		if a.window == nil {
			w := a.app.NewWindow("App")
			a.window = w
		}
		a.window.Resize(fyne.NewSize(w, h))

		a.height = h
		a.width = w
	}
}

func WithTick(t time.Duration) func(*App) {
	return func(a *App) {
		a.tick = t
	}
}

func WithDisplayGrid(rows, cols int) func(*App) {
	return func(a *App) {
		disp := display.NewGrid(rows, cols, int(a.width))
		a.window.SetContent(disp.Content())
		a.sm = disp
	}
}

func (a *App) GetStateManager() StateManager {
	return a.sm
}

func (a *App) Run(sp display.StateProvider) {
	go func() {
		a.sm.InitState(sp)
		for range time.Tick(tick) {
			if sp.Ended() {
				break
			}

			sp.Step()
			a.sm.UpdateState(sp)
		}
	}()
	a.window.ShowAndRun()
}
