package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const iso = 1

type metaballsView struct {
	fg    color.Color
	model *ensemble
}

func newMetaballsView(m *ensemble, fg color.Color) *metaballsView {
	v := &metaballsView{
		fg:    fg,
		model: m,
	}
	return v
}

func (v *metaballsView) draw(container *fyne.Container) {
	size := container.Size()
	g := grid(size.Width, size.Height)
	dx := 1 / size.Width
	dy := 1 / size.Height

	for row := float32(0); row < size.Height; row += g {
		y := row * dy
		for col := float32(0); col < size.Width; col += g {
			x := col * dx

			m := v.model
			a := m.value(x, y)
			b := m.value(x+dx*g, y)
			c := m.value(x+dx*g, y+dy*g)
			d := m.value(x, y+dy*g)

			a1, a2 := lerp(col, col+g, (iso-a)/(b-a)), row
			b1, b2 := col+g, lerp(row, row+g, (iso-b)/(c-b))
			c1, c2 := lerp(col, col+g, (iso-d)/(c-d)), row+g
			d1, d2 := col, lerp(row, row+g, (iso-a)/(d-a))

			switch state(a, b, c, d) {
			case 1, 14:
				line(container, v.fg, c1, c2, d1, d2)
			case 2, 13:
				line(container, v.fg, b1, b2, c1, c2)
			case 3, 12:
				line(container, v.fg, b1, b2, d1, d2)
			case 4:
				line(container, v.fg, a1, a2, b1, b2)
			case 5:
				line(container, v.fg, a1, a2, d1, d2)
				line(container, v.fg, b1, b2, c1, c2)
			case 6:
				line(container, v.fg, a1, a2, c1, c2)
			case 7, 8:
				line(container, v.fg, a1, a2, d1, d2)
			case 9:
				line(container, v.fg, a1, a2, c1, c2)
			case 10:
				line(container, v.fg, a1, a2, b1, b2)
				line(container, v.fg, c1, c2, d1, d2)
			case 11:
				line(container, v.fg, a1, a2, b1, b2)
			}
		}
	}
}

func line(container *fyne.Container, color color.Color, x1, y1, x2, y2 float32) {
	l := canvas.NewLine(color)
	l.Position1 = fyne.NewPos(x1, y1)
	l.Position2 = fyne.NewPos(x2, y2)
	container.Add(l)
}

func grid(w, h float32) float32 {
	return float32(math.Ceil(math.Sqrt(float64(w*h) / 10000)))
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

func lerp(a, b, t float32) float32 {
	if t < 0 {
		return a
	}
	if t > 1 {
		return b
	}
	return float32(math.Round(float64(a + (b-a)*t)))
}
