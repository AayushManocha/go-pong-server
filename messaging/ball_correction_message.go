package messaging

import "github.com/AayushManocha/go-game-server/game"

type BallCorrectionMessage struct {
	Type   string `json:"type"`
	SpeedX float64
	SpeedY float64
	X      float64
	Y      float64
}

func NewBallCorectionMessage(g *game.Game) BallCorrectionMessage {
	return BallCorrectionMessage{
		Type:   "BALL_CORRECTION_MESSAGE",
		SpeedX: g.Ball.SpeedX,
		SpeedY: g.Ball.SpeedY,

		X: g.Ball.Shape.X,
		Y: g.Ball.Shape.Y,
	}
}
