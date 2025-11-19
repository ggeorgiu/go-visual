package display

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type pixel interface {
	setColor(color.Color)
	setBounds([]bool)
	getBounds() []fyne.CanvasObject

	get() []fyne.CanvasObject
}

const gridSize = 5

// 10 X 10
type chunk struct {
	pixels  [][]pixel
	content *fyne.Container
}

type Grid struct {
	rows int
	cols int

	chunks [][]chunk
}

func NewGrid(r, c, ws int) *Grid {
	chunks := make([][]chunk, r/gridSize)

	for i := 1; i < r; i += gridSize {
		ch := make([]chunk, c/gridSize)
		for j := 1; j < c; j += gridSize {
			ch[j/gridSize] = chunk{pixels: make([][]pixel, gridSize)}
		}

		chunks[i/gridSize] = ch
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			pix := newRect(i, j, ws/r)

			chunks[i/gridSize][j/gridSize].pixels[i%gridSize] = append(chunks[i/gridSize][j/gridSize].pixels[i%gridSize], pix)
		}
	}

	return &Grid{
		rows:   r,
		cols:   c,
		chunks: chunks,
	}
}

func (g *Grid) InitState(sp StateProvider) {
	for _, s := range sp.GetInitialState() {
		g.chunks[s.i/gridSize][s.j/gridSize].pixels[s.i%gridSize][s.j%gridSize].setColor(s.c)
	}
}

func (g *Grid) UpdateState(sp StateProvider) {
	for _, v := range sp.GetUpdatedState() {
		c := g.chunks[v.i/gridSize][v.j/gridSize]
		//cl := color.RGBA{
		//	R: 255,
		//	G: 255,
		//	B: 0,
		//	A: 222,
		//}
		//c.pixels[0][0].setColor(cl)
		//c.pixels[0][gridSize-1].setColor(cl)
		//c.pixels[gridSize-1][0].setColor(cl)
		//c.pixels[gridSize-1][gridSize-1].setColor(cl)

		for _, b := range c.pixels[v.i%gridSize][v.j%gridSize].getBounds() {
			c.content.Remove(b)
		}

		c.pixels[v.i%gridSize][v.j%gridSize].setColor(v.c)
		if v.borders != nil {
			c.pixels[v.i%gridSize][v.j%gridSize].setBounds(v.borders)
		}

		for _, b := range c.pixels[v.i%gridSize][v.j%gridSize].getBounds() {
			c.content.Add(b)
		}

		c.content.Refresh()
	}
}

func (g *Grid) Content() fyne.CanvasObject {
	var obj []fyne.CanvasObject

	for i := 0; i < len(g.chunks); i++ {
		for j := 0; j < len(g.chunks[0]); j++ {
			obj = append(obj, g.chunks[i][j].Content())
		}
	}

	return container.NewWithoutLayout(obj...)
}

func (c *chunk) Content() fyne.CanvasObject {
	var obj []fyne.CanvasObject

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			obj = append(obj, c.pixels[i][j].get()...)
		}
	}

	c.content = container.NewWithoutLayout(obj...)
	return c.content
}
