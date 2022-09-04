package main

import (
	"math/rand"
	"sync"
)

type ball struct {
	x  float32
	y  float32
	r  float32
	vx float32
	vy float32
}

type ensemble struct {
	balls []ball
	mutex sync.RWMutex
}

func newRandomBall() ball {
	return ball{
		x:  rand.Float32()*0.8 + 0.1,
		y:  rand.Float32()*0.8 + 0.1,
		r:  rand.Float32()*0.1 + 0.025,
		vx: (rand.Float32() - 0.5) * 0.005,
		vy: (rand.Float32() - 0.5) * 0.005,
	}
}

func (b *ball) move() {
	b.x += b.vx
	if b.x < b.r {
		b.x = b.r
		b.vx = -b.vx
	} else if b.x > 1-b.r {
		b.x = 1 - b.r
		b.vx = -b.vx
	}
	b.y += b.vy
	if b.y < b.r {
		b.y = b.r
		b.vy = -b.vy
	} else if b.y > 1-b.r {
		b.y = 1 - b.r
		b.vy = -b.vy
	}
}

func (e *ensemble) value(x, y float32) float32 {
	rl := e.mutex.RLocker()
	rl.Lock()
	defer rl.Unlock()
	var res float32
	for _, b := range e.balls {
		dx := x - b.x
		dy := y - b.y
		res += b.r * b.r / (dx*dx + dy*dy)
	}
	return res
}

func newRandomEnsemble(n int) *ensemble {
	balls := make([]ball, 0, n)
	for i := 0; i < n; i++ {
		balls = append(balls, newRandomBall())
	}
	return &ensemble{balls: balls}
}

func (e *ensemble) move() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for i := range e.balls {
		e.balls[i].move()
	}
}
