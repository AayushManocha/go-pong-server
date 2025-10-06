package messaging

import "github.com/AayushManocha/go-game-server/game"

type GameStopMessage struct {
	Type string `json:"type"`
}

func NewGameStopMessage(g *game.Game) GameStopMessage {
	return GameStopMessage{
		Type: "GAME_STOP_MESSAGE",
	}
}
