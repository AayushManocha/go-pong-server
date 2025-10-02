package game

import "math/rand"

type Ball struct {
	Shape  *Rectangle
	SpeedX int
	SpeedY int
}

func GenerateRandomSpeed() int {
	baseSpeed := 50

	directionBase := rand.Intn(100)
	if directionBase > 50 {
		return baseSpeed
	} else {
		return baseSpeed * -1
	}
}
