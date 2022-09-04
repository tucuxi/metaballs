package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("com.github.tucuxi.metaballs")
	w := a.NewWindow("Metaballs")
	model := newRandomEnsemble(8)
	widget := newMetaballsWidget(model)
	w.SetContent(widget)
	widget.animate()
	w.ShowAndRun()
}
