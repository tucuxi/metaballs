package internal

import (
	"sync"
)

type Group struct {
	balls []Ball
	mutex *sync.RWMutex
}

func newRandomGroup(n int, s BallSpeed, z BallSize) *Group {
	balls := make([]Ball, n)
	for i := 0; i < n; i++ {
		balls[i] = NewRandomBall(s, z)
	}
	return &Group{
		balls: balls,
		mutex: &sync.RWMutex{},
	}
}

//	 n
//		∑     =  r(i)^2/ (x-x(i))^2 + (y-y(i))^2
//	 i=0
//
// This is the function of drawing a metaball
func (e *Group) value(x, y float32) float32 {
	// rl := e.mutex.RLocker()
	// rl.Lock()
	// defer rl.Unlock()

	var res float32
	for _, b := range e.balls {
		dx := x - b.pos.x
		dy := y - b.pos.y
		res += square(b.radius) / (square(dx) + square(dy))
	}

	return res
}

//	 n
//		∑     =  r(i)^2/ (x-x(i))^2 + (y-y(i))^2
//	 i=0
func (e *Group) valueV2(x1, x2, y1, y2 float32) (float32, float32, float32, float32) {
	// there isn't really a need to lock since this will rarely confilct with move(),
	// since they both work on different intervals
	// even when it happens its not really noticable
	// rl := e.mutex.RLocker()
	// rl.Lock()
	// defer rl.Unlock()

	var a, b, c, d float32
	for _, ball := range e.balls {
		a += circleFormula(ball.radius, ball.pos.x, ball.pos.y, x1, y1)
		b += circleFormula(ball.radius, ball.pos.x, ball.pos.y, x2, y1)
		c += circleFormula(ball.radius, ball.pos.x, ball.pos.y, x2, y2)
		d += circleFormula(ball.radius, ball.pos.x, ball.pos.y, x1, y2)
	}

	return a, b, c, d
}
func circleFormula(r, px, py, x, y float32) float32 {
	return square(r) / (square(x-px) + square(y-py))
}

func (e *Group) move() {
	// e.mutex.Lock()
	// defer e.mutex.Unlock()
	for i := range e.balls {
		e.balls[i].move()
	}
}
