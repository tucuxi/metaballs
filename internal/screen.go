package internal

import (
	"image"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/StephaneBunel/bresenham"

	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ISO Contour Value
const iso = 1

type Screen struct {
	widget.BaseWidget
	raster *canvas.Raster
	group  *Group

	resolution Resolution

	painter chan *Line
	img     *image.RGBA
	color   color.Color
}

func NewScreen(g *Group, r Resolution,color BallColor) *Screen {
	s := &Screen{
		group:      g,
		resolution: r,
		painter:    make(chan *Line),
		color:     color,
	}
	s.raster = canvas.NewRaster(s.draw)
	// go s.StartPainter()
	return s
}

func (bw *Screen) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(bw.raster)
}

func (s *Screen) StartPainter() {

	for {
		select {
		case line := <-s.painter:
			go bresenham.DrawLine(s.img, line.x1, line.y1, line.x2, line.y2, s.color)
		}
	}
}

// Called when s.Refresh() is called
// Re-render the screen
func (s *Screen) draw(w, h int) image.Image {
	// start := time.Now()
	// fgcolor := theme.WarningColor()
	img := image.NewRGBA(image.Rect(0, 0, int(s.Size().Height), int(s.Size().Width)))

	size := float32(max(w, h))
	g := int(math.Ceil(float64(size) / float64(s.resolution)))

	for row := 0; row < h; row += g {
		y := float32(row) / size
		y2 := float32(row+g) / size

		for col := 0; col < w; col += g {

			x := float32(col) / size
			x2 := float32(col+g) / size

			a, b, c, d := s.group.valueV2(x, x2, y, y2)
			sq := Square{
				a: a,
				b: b,
				c: c,
				d: d,
			}
			// lines := sq.March(col, row, g)
			// for _, line := range lines {

			// 	bresenham.DrawLine(img, line.x1, line.y1, line.x2, line.y2, color)
			// }
			// go sq.MarchV3(col, row, g)
			sq.MarchV2(col, row, g, img, s.color)
		}

	}

	// elapsed := time.Since(start)
	// fmt.Println("draw time: ", elapsed)

	return img
}
