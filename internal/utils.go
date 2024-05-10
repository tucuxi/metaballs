package internal

import "math/rand/v2"

func randRangeInt(min, max int) int {
	return rand.IntN(max-min) + min
}

func randRangeFloat(min, max float32) float32 {
	return rand.Float32()*(max-min) + min
}

func square(x float32) float32 {
	return x * x
}
