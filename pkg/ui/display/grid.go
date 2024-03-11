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
	chunks := make([][]chunk, r/10)

	for i := 1; i < r; i += 10 {
		ch := make([]chunk, c/10)
		for j := 1; j < c; j += 10 {
			ch[j/10] = chunk{pixels: make([][]pixel, 10)}
		}

		chunks[i/10] = ch
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			pix := newRect(i, j, ws/r)

			chunks[i/10][j/10].pixels[i%10] = append(chunks[i/10][j/10].pixels[i%10], pix)
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
		g.chunks[s.i/10][s.j/10].pixels[s.i%10][s.j%10].setColor(s.c)
	}
}

func (g *Grid) UpdateState(sp StateProvider) {
	for _, v := range sp.GetUpdatedState() {

		for _, b := range g.chunks[v.i/10][v.j/10].pixels[v.i%10][v.j%10].getBounds() {
			g.chunks[v.i/10][v.j/10].content.Remove(b)
		}

		g.chunks[v.i/10][v.j/10].pixels[v.i%10][v.j%10].setColor(v.c)
		if v.borders != nil {
			g.chunks[v.i/10][v.j/10].pixels[v.i%10][v.j%10].setBounds(v.borders)
		}

		for _, b := range g.chunks[v.i/10][v.j/10].pixels[v.i%10][v.j%10].getBounds() {
			g.chunks[v.i/10][v.j/10].content.Add(b)
		}

		g.chunks[v.i/10][v.j/10].content.Refresh()
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

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			obj = append(obj, c.pixels[i][j].get()...)
		}
	}

	c.content = container.NewWithoutLayout(obj...)
	return c.content
}
