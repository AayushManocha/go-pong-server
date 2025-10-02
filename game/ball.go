package game

import "math/rand"

type Ball struct {
	Shape  *Rectangle
	SpeedX float64
	SpeedY float64
}

// Note speed is in px/ms
func GenerateRandomSpeed() float64 {
	baseSpeed := 0.1

	directionBase := rand.Intn(100)
	if directionBase > 50 {
		return baseSpeed
	} else {
		return baseSpeed * -1
	}
}
