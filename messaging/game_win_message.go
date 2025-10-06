package messaging

import "github.com/AayushManocha/go-game-server/game"

type GameWinMessage struct {
	Type        string `json:"type"`
	PlayerIndex int    `json:"playerIndex"`
}

func NewGameWinMessage(g *game.Game) GameWinMessage {
	return GameWinMessage{
		Type:        "GAME_WIN_MESSAGE",
		PlayerIndex: g.Winner,
	}
}
