package main

import (
	"image"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/StephaneBunel/bresenham"
)

const (
	iso    = 1
	raster = 128
)

type metaballsWidget struct {
	widget.BaseWidget

	model  *ensemble
	raster *canvas.Raster
}

func newMetaballsWidget(m *ensemble) *metaballsWidget {
	mw := &metaballsWidget{model: m}
	mw.raster = canvas.NewRaster(mw.draw)
	mw.ExtendBaseWidget(mw)
	return mw
}

func (mw *metaballsWidget) animate() {
	go func() {
		for range time.Tick(time.Millisecond * 10) {
			mw.model.move()
			fyne.Do(func() { mw.Refresh() })
		}
	}()
}

func (mw *metaballsWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(mw.raster)
}

func (mw *metaballsWidget) draw(w, h int) image.Image {
	fgcolor := theme.Color(theme.ColorNameForeground)
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	size := min(h, w)

	// dx and dy are the offsets to center the rendering area.
	dx := float32(w-size) / 2
	dy := float32(h-size) / 2

	gridStep := float32(1) / raster

	row1 := make([]float32, raster+1)
	row2 := make([]float32, raster+1)

	g := int(float32(size) / raster) // Cell size in pixels

	for j := range row1 {
		row1[j] = mw.model.value(float32(j)*gridStep, 0)
	}

	for i := range raster {
		y := int(dy + float32(i)*float32(size)/raster)

		// Calculate the next row of values
		for j := range raster + 1 {
			row2[j] = mw.model.value(float32(j)*gridStep, float32(i+1)*gridStep)
		}

		for j := range raster {
			// Get values at the corners of the grid cell
			a, b := row1[j], row1[j+1]
			c, d := row2[j+1], row2[j]

			// Calculate pixel coordinates for the left of the current grid cell and cell size
			x := int(dx + float32(j)*float32(size)/raster)

			// Calculate intersection points using linear interpolation
			a1, a2 := lerp(x, x+g, (iso-a)/(b-a)), y   // Top edge
			b1, b2 := x+g, lerp(y, y+g, (iso-b)/(c-b)) // Right edge
			c1, c2 := lerp(x, x+g, (iso-d)/(c-d)), y+g // Bottom edge
			d1, d2 := x, lerp(y, y+g, (iso-a)/(d-a))   // Left edge

			switch state(a, b, c, d) {
			case 1, 14:
				bresenham.DrawLine(img, c1, c2, d1, d2, fgcolor)
			case 2, 13:
				bresenham.DrawLine(img, b1, b2, c1, c2, fgcolor)
			case 3, 12:
				bresenham.DrawLine(img, b1, b2, d1, d2, fgcolor)
			case 4, 11:
				bresenham.DrawLine(img, a1, a2, b1, b2, fgcolor)
			case 5, 10:
				bresenham.DrawLine(img, a1, a2, d1, d2, fgcolor)
				bresenham.DrawLine(img, b1, b2, c1, c2, fgcolor)
			case 6, 9:
				bresenham.DrawLine(img, a1, a2, c1, c2, fgcolor)
			case 7, 8:
				bresenham.DrawLine(img, a1, a2, d1, d2, fgcolor)
			}
		}
		row1, row2 = row2, row1
	}
	return img
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
