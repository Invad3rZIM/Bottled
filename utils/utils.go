package utils

import (
	"math/rand"
	"time"
)

//filler id generator function. needs rework later
func GenFloat(min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano())

	return (max-min)*rand.Float64() + min
}

//filler id generator function. needs rework later
func GenInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}
