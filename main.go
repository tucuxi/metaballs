package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Metaballs")
	model := newRandomEnsemble(10)
	view := newMetaballsView(model, color.NRGBA{R: 200, G: 200, B: 0, A: 255})
	content := container.NewWithoutLayout()
	content.Resize(fyne.NewSize(500, 500))
	w.SetContent(content)
	go func() {
		for range time.Tick(time.Millisecond * 50) {
			model.move()
			content.RemoveAll()
			view.draw(content)
			content.Refresh()
		}
	}()
	w.ShowAndRun()
}
