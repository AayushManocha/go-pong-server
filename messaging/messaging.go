package messaging

import (
	"fmt"
	"io"

	"github.com/AayushManocha/go-game-server/game"
	"github.com/gorilla/websocket"
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

		err := conn.Connection.WriteJSON(NewGameMessage(g))
		if err != nil {
			if err == io.EOF {
				fmt.Printf("EOF Error received \n")
			} else if err == websocket.ErrCloseSent {
				fmt.Printf("ErrClose Error received \n")
			} else {
				fmt.Printf("Received err: %s \n", err.Error())
			}

			// Stop game and remove player
			g.Quit_ch <- true

		}

		conn.Mu.Unlock()
	}
}
