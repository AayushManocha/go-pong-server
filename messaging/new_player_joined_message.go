package messaging

import "github.com/AayushManocha/go-game-server/game"

type PlayerMoveMessage struct {
	Type        string  `json:"type"`
	PlayerIndex int     `json:"playerIndex"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
}

func NewPlayerJoinedMessage(p *game.Player) PlayerJoinedMessage {
	return PlayerJoinedMessage{
		Type:   "PLAYER_MESSAGE",
		Player: p,
	}
}
