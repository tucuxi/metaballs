package main

import (
	"image"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/StephaneBunel/bresenham"
)

const iso = 1

type metaballsRenderer struct {
	raster  *canvas.Raster
	objects []fyne.CanvasObject
	fgcolor color.Color
	bgcolor color.Color
	widget  *metaballsWidget
}

func (r *metaballsRenderer) ApplyTheme() {
	r.fgcolor = theme.ForegroundColor()
	r.bgcolor = theme.BackgroundColor()
}

func (r *metaballsRenderer) Destroy() {
}

func (r *metaballsRenderer) draw(w, h int) image.Image {
	g := grid(w, h)
	gx := float32(g) / float32(w)
	gy := float32(g) / float32(h)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for row := 0; row < h; row += g {
		y := float32(row) / float32(h)
		for col := 0; col < w; col += g {
			x := float32(col) / float32(w)
			m := r.widget.model
			a := m.value(x, y)
			b := m.value(x+gx, y)
			c := m.value(x+gx, y+gy)
			d := m.value(x, y+gy)

			a1, a2 := lerp(col, col+g, (iso-a)/(b-a)), row
			b1, b2 := col+g, lerp(row, row+g, (iso-b)/(c-b))
			c1, c2 := lerp(col, col+g, (iso-d)/(c-d)), row+g
			d1, d2 := col, lerp(row, row+g, (iso-a)/(d-a))

			switch state(a, b, c, d) {
			case 1, 14:
				bresenham.DrawLine(img, c1, c2, d1, d2, r.fgcolor)
			case 2, 13:
				bresenham.DrawLine(img, b1, b2, c1, c2, r.fgcolor)
			case 3, 12:
				bresenham.DrawLine(img, b1, b2, d1, d2, r.fgcolor)
			case 4:
				bresenham.DrawLine(img, a1, a2, b1, b2, r.fgcolor)
			case 5:
				bresenham.DrawLine(img, a1, a2, d1, d2, r.fgcolor)
				bresenham.DrawLine(img, b1, b2, c1, c2, r.fgcolor)
			case 6:
				bresenham.DrawLine(img, a1, a2, c1, c2, r.fgcolor)
			case 7, 8:
				bresenham.DrawLine(img, a1, a2, d1, d2, r.fgcolor)
			case 9:
				bresenham.DrawLine(img, a1, a2, c1, c2, r.fgcolor)
			case 10:
				bresenham.DrawLine(img, a1, a2, b1, b2, r.fgcolor)
				bresenham.DrawLine(img, c1, c2, d1, d2, r.fgcolor)
			case 11:
				bresenham.DrawLine(img, a1, a2, b1, b2, r.fgcolor)
			}
		}
	}
	return img
}

func (r *metaballsRenderer) Layout(size fyne.Size) {
	r.raster.Resize(size)
}

func (r *metaballsRenderer) MinSize() fyne.Size {
	return fyne.NewSize(float32(64), float32(64))
}

func (r *metaballsRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *metaballsRenderer) Refresh() {
	canvas.Refresh(r.raster)
}

type metaballsWidget struct {
	widget.BaseWidget

	model *ensemble
}

func newMetaballsWidget(m *ensemble) *metaballsWidget {
	w := &metaballsWidget{model: m}
	w.ExtendBaseWidget(w)
	return w
}

func (w *metaballsWidget) animate() {
	go func() {
		for range time.Tick(time.Millisecond * 50) {
			w.model.move()
			w.Refresh()
		}
	}()
}

func (w *metaballsWidget) CreateRenderer() fyne.WidgetRenderer {
	renderer := &metaballsRenderer{widget: w}
	raster := canvas.NewRaster(renderer.draw)
	renderer.raster = raster
	renderer.objects = []fyne.CanvasObject{raster}
	renderer.ApplyTheme()
	return renderer
}

func grid(w, h int) int {
	return int(math.Ceil(math.Sqrt(float64(w*h) / 16384)))
}

func state(tl, tr, br, bl float32) int {
	res := 0
	if tl >= iso {
		res |= 8
	}
	if tr >= iso {
		res |= 4
	}
	if br >= iso {
		res |= 2
	}
	if bl >= iso {
		res |= 1
	}
	return res
}

func lerp(a, b int, t float32) int {
	if t < 0 {
		return a
	}
	if t > 1 {
		return b
	}
	return int(math.Round(float64(a) + float64(b-a)*float64(t)))
}
