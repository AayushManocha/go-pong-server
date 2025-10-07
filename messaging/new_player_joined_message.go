package messaging

import "github.com/AayushManocha/go-game-server/game"

type PlayerJoinedMessage struct {
	Type   string `json:"type"`
	Player *game.Player
}

func NewPlayerJoinedMessage(p *game.Player) PlayerJoinedMessage {
	return PlayerJoinedMessage{
		Type:   "PLAYER_JOINED_MESSAGE",
		Player: p,
	}
}
