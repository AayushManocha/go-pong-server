package messaging

import "github.com/AayushManocha/go-game-server/game"

type PlayerMoveMessage struct {
	Type        string  `json:"type"`
	PlayerIndex int     `json:"playerIndex"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
}

func NewPlayerMoveMessage(p *game.Player) PlayerMoveMessage {
	return PlayerMoveMessage{
		Type:        "PLAYER_MOVE_MESSAGE",
		PlayerIndex: p.Index,
		X:           p.Shape.X,
		Y:           p.Shape.Y,
	}
}
