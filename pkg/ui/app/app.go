package app

import (
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"

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

	canPause bool
	paused   bool
	stepKey  rune

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

func WithSize(w, h int) func(*App) {
	return func(a *App) {
		if a.window == nil {
			w := a.app.NewWindow("App")
			a.window = w
		}
		a.window.Resize(fyne.NewSize(float32(w+7), float32(h+7)))

		a.height = float32(h)
		a.width = float32(w)
	}
}

func WithTick(t time.Duration) func(*App) {
	return func(a *App) {
		a.tick = t
	}
}

func WithStepOnKey(k rune) func(*App) {
	return func(a *App) {
		a.canPause = true
		a.stepKey = k

		if deskCanvas, ok := a.window.Canvas().(desktop.Canvas); ok {
			deskCanvas.SetOnKeyUp(func(key *fyne.KeyEvent) {
				if strings.ToLower(string(key.Name)) == string(k) {
					a.paused = false
				}
			})
		}
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
		for range time.Tick(a.tick) {
			if sp.Ended() {
				break
			}

			if a.paused {
				continue
			}

			sp.Step()
			a.sm.UpdateState(sp)

			if a.canPause && !a.paused {
				a.paused = true
			}
		}
	}()

	a.window.ShowAndRun()
}
