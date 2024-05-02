package internal

import (
	"image/color"
	"math"
	"math/rand"
)

type Position struct {
	x float32
	y float32
}

func NewrandomPosition() Position {
	return Position{
		x: rand.Float32()*0.8 + 0.1,
		y: rand.Float32()*0.8 + 0.1,
	}
}

type Velocity struct {
	x float32
	y float32
}

func NewVelocity(s BallSpeed) Velocity {
	return Velocity{
		x: randRangeFloat(s.minVelocity, s.maxVelocity),
		y: randRangeFloat(s.minVelocity, s.maxVelocity),

	}
}

type Ball struct {
	pos    Position
	vel    Velocity
	radius float32
	color  color.Color
}

func NewRandomBall(s BallSpeed,z BallSize) Ball {
	return Ball{
		pos:    NewrandomPosition(),
		radius: randRangeFloat(z.minRadius, z.maxRadius),
		vel:    NewVelocity(s),
	}
}

func (b *Ball) move() {
	b.pos.x += b.vel.x
	if b.pos.x <= 0.1 || b.pos.x >= 0.9 {
		b.vel.x = -b.vel.x
	}
	b.pos.y += b.vel.y
	if b.pos.y <= 0.1 || b.pos.y >= 0.9 {
		b.vel.y = -b.vel.y
	}
}

// return if the point is on the contour of the ball
// and return the distance from the point to the center of the ball
func (b *Ball) IsOnPointOnContour(x, y float32) (bool, float32) {
	dx := x - b.pos.x
	dy := y - b.pos.y
	d := float32(math.Sqrt(float64(square(dx) + square(dy))))
	return d >= b.radius-0.01 && d <= b.radius+0.01, d
}