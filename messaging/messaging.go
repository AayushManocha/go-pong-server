package messaging

import (
	"fmt"
	"io"

	"github.com/AayushManocha/go-game-server/game"
	"github.com/gorilla/websocket"
)

type Message interface{}

type GameMessage struct {
	Type string `json:"type"`
	Game *game.Game
}

func NewGameMessage(g *game.Game) GameMessage {
	return GameMessage{
		Type: "GAME_MESSAGE",
		Game: g,
	}
}

func BroadcastGame(g *game.Game) {
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

func BroadcastToAllPlayers(g *game.Game, msg Message) {
	players := g.Players
	for _, p := range players {
		conn := p.Connection
		conn.Mu.Lock()
		if err := conn.Connection.WriteJSON(msg); err != nil {
			HandleMessageError()
		}
		conn.Mu.Unlock()
	}
}

func BroadcastToOtherPlayers(from *game.Player, g *game.Game, msg Message) {
	players := g.Players
	for _, p := range players {
		conn := p.Connection
		if p.Index != from.Index {
			conn.Mu.Lock()
			if err := conn.Connection.WriteJSON(msg); err != nil {
				HandleMessageError()
			}
			conn.Mu.Unlock()
		}
	}
}

func HandleMessageError() {}
