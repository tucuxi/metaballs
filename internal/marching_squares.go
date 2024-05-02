package internal

import (
	"image"
	"image/color"
	"math"

	"github.com/StephaneBunel/bresenham"
)

type Square struct {
	a    float32
	b    float32
	c    float32
	d    float32
	send chan *Line
}
type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

// Apply the marching squares algorithm to the square
// and return the lines that make up the contour.
func (sq *Square) March(col, row, g int) []Line {

	lines := []Line{}

	a1, a2 := lerp(col, col+g, (iso-sq.a)/(sq.b-sq.a)), row
	b1, b2 := col+g, lerp(row, row+g, (iso-sq.b)/(sq.c-sq.b))
	c1, c2 := lerp(col, col+g, (iso-sq.d)/(sq.c-sq.d)), row+g
	d1, d2 := col, lerp(row, row+g, (iso-sq.a)/(sq.d-sq.a))

	switch mask(sq.a, sq.b, sq.c, sq.d) {
	case 1, 14:
		lines = append(lines, Line{c1, c2, d1, d2})

	case 2, 13:
		lines = append(lines, Line{b1, b2, c1, c2})

	case 3, 12:
		lines = append(lines, Line{b1, b2, d1, d2})

	case 4:
		lines = append(lines, Line{a1, a2, b1, b2})

	case 5:
		lines = append(lines, Line{a1, a2, d1, d2})
		lines = append(lines, Line{b1, b2, c1, c2})

	case 6:
		lines = append(lines, Line{a1, a2, c1, c2})

	case 7, 8:
		lines = append(lines, Line{a1, a2, d1, d2})

	case 9:
		lines = append(lines, Line{a1, a2, c1, c2})

	case 10:
		lines = append(lines, Line{a1, a2, b1, b2})
		lines = append(lines, Line{c1, c2, d1, d2})

	case 11:
		lines = append(lines, Line{a1, a2, b1, b2})

	// if the square is inside the contour line, fill the square with the color
	case 15:

		lines = append(lines, Line{a1, a2, b1, b2})
		lines = append(lines, Line{b1, b2, c1, c2})
		lines = append(lines, Line{c1, c2, d1, d2})
		lines = append(lines, Line{d1, d2, a1, a2})

	}

	return lines
}

func (sq *Square) MarchV2(col, row, g int, img *image.RGBA, color color.Color) {

	a1, a2 := lerp(col, col+g, (iso-sq.a)/(sq.b-sq.a)), row
	b1, b2 := col+g, lerp(row, row+g, (iso-sq.b)/(sq.c-sq.b))
	c1, c2 := lerp(col, col+g, (iso-sq.d)/(sq.c-sq.d)), row+g
	d1, d2 := col, lerp(row, row+g, (iso-sq.a)/(sq.d-sq.a))

	switch mask(sq.a, sq.b, sq.c, sq.d) {
	case 1, 14:

		bresenham.DrawLine(img, c1, c2, d1, d2, color)

	case 2, 13:

		bresenham.DrawLine(img, b1, b2, c1, c2, color)

	case 3, 12:
		bresenham.DrawLine(img, b1, b2, d1, d2, color)

	case 4:
		bresenham.DrawLine(img, a1, a2, b1, b2, color)

	case 5:

		bresenham.DrawLine(img, a1, a2, d1, d2, color)
		bresenham.DrawLine(img, b1, b2, c1, c2, color)

	case 6:

		bresenham.DrawLine(img, a1, a2, c1, c2, color)

	case 7, 8:
		// lines = append(lines, Line{a1, a2, d1, d2})
		bresenham.DrawLine(img, a1, a2, d1, d2, color)

	case 9:
		bresenham.DrawLine(img, a1, a2, c1, c2, color)

	case 10:
		bresenham.DrawLine(img, a1, a2, b1, b2, color)
		bresenham.DrawLine(img, c1, c2, d1, d2, color)

	case 11:
		// lines = append(lines, Line{a1, a2, b1, b2})
		bresenham.DrawLine(img, a1, a2, b1, b2, color)

	// if the square is inside the contour line, fill the square with the color
	case 15:
		// fill the square with the color
		bresenham.DrawLine(img, a1, a2, b1, b2, color)
		bresenham.DrawLine(img, b1, b2, c1, c2, color)
		bresenham.DrawLine(img, c1, c2, d1, d2, color)
		bresenham.DrawLine(img, d1, d2, a1, a2, color)
		
		// draw the center
		bresenham.DrawLine(img, (a1+b1+c1+d1)/4, (a2+b2+c2+d2)/4, (a1+b1+c1+d1)/4, (a2+b2+c2+d2)/4, color)


	}

	
}

// Apply the marching squares algorithm to the square
// and return the lines that make up the contour.
func (sq *Square) MarchV3(col, row, g int) {

	a1, a2 := lerp(col, col+g, (iso-sq.a)/(sq.b-sq.a)), row
	b1, b2 := col+g, lerp(row, row+g, (iso-sq.b)/(sq.c-sq.b))
	c1, c2 := lerp(col, col+g, (iso-sq.d)/(sq.c-sq.d)), row+g
	d1, d2 := col, lerp(row, row+g, (iso-sq.a)/(sq.d-sq.a))

	switch mask(sq.a, sq.b, sq.c, sq.d) {
	case 1, 14:
		sq.send <- &Line{c1, c2, d1, d2}

	case 2, 13:
		sq.send <- &Line{b1, b2, c1, c2}

	case 3, 12:
		sq.send <- &Line{b1, b2, d1, d2}

	case 4:
		sq.send <- &Line{a1, a2, b1, b2}

	case 5:

		sq.send <- &Line{a1, a2, d1, d2}
		sq.send <- &Line{b1, b2, c1, c2}

	case 6:
		sq.send <- &Line{a1, a2, c1, c2}

	case 7, 8:
		sq.send <- &Line{a1, a2, d1, d2}

	case 9:
		sq.send <- &Line{a1, a2, c1, c2}

	case 10:

		sq.send <- &Line{a1, a2, b1, b2}
		sq.send <- &Line{c1, c2, d1, d2}

	case 11:

		sq.send <- &Line{a1, a2, b1, b2}

	// if the square is inside the contour line, fill the square with the color
	case 15:

		sq.send <- &Line{a1, a2, b1, b2}
		sq.send <- &Line{b1, b2, c1, c2}
		sq.send <- &Line{c1, c2, d1, d2}
		sq.send <- &Line{d1, d2, a1, a2}

	}

}

// mask is a function that returns a 4-bit number that represents the contour line's position in the cell.
// as you know marching squares algorithm divides the cell into 16 different cases.
// the mask function returns which case the contour line is in.
func mask(tl, tr, br, bl float32) int {
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

// Linear interpolation
// The goal of this function is to find the value where the contour line really intersects the cell.
// intersect = By+(Dy−By)t
// t = 1−f(Bx,By) / f(Dx,Dy)−f(Bx,By)
func lerp(a, b int, t float32) int {
	if t < 0 {
		return a
	}
	if t > 1 {
		return b
	}
	return int(math.Round(float64(a) + float64(b-a)*float64(t)))
}
