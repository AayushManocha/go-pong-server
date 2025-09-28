package messaging

import (
	"github.com/AayushManocha/go-game-server/game"
)

type PlayerMessage struct {
	Type   string
	Player *game.Player
}

type GameMessage struct {
	Type string
	Game *game.Game
}

type GenericErrorMessage struct {
	Type    string
	Message string
}

func NewPlayerMessage(p *game.Player) PlayerMessage {
	return PlayerMessage{
		Type:   "PLAYER_MESSAGE",
		Player: p,
	}
}

func NewGameMessage(g *game.Game) GameMessage {
	return GameMessage{
		Type: "GAME_MESSAGE",
		Game: g,
	}
}

func NewGenericErrorMessage(msg string) GenericErrorMessage {
	return GenericErrorMessage{
		Type:    "ERROR_MESSAGE",
		Message: msg,
	}
}

func BroadcastUpdates(g *game.Game) {
	players := g.Players

	for _, p := range players {
		conn := p.Connection

		conn.Mu.Lock()
		conn.Connection.WriteJSON(NewGameMessage(g))
		conn.Mu.Unlock()
	}
}
