package messaging

import (
	"fmt"
	"io"

	"github.com/AayushManocha/go-game-server/game"
	"github.com/gorilla/websocket"
)

type PlayerMessage struct {
	Type   string `json:"type"`
	Player *game.Player
}

type GameMessage struct {
	Type string `json:"type"`
	Game *game.Game
}

type GenericErrorMessage struct {
	Type    string `json:"type"`
	Message string
}

type CollisionMessage struct {
	Type   string `json:"type"`
	XSpeed int    `json:"xSpeed'`
	YSpeed int    `json:"ySpeed"`
}

type PlayerMoveMessage struct {
	Type        string `json:"type"`
	PlayerIndex int    `json:"playerIndex"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
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

func BroadcastPlayerMove(g *game.Game, p *game.Player) {
	players := g.Players

	msg := PlayerMoveMessage{
		Type:        "PLAYER_MOVE_MESSAGE",
		PlayerIndex: p.Index,
		X:           p.Shape.X,
		Y:           p.Shape.Y,
	}

	fmt.Printf("Broadcast message: %+v \n", msg)

	for _, p := range players {
		conn := p.Connection

		conn.Mu.Lock()
		err := conn.Connection.WriteJSON(msg)
		if err != nil {
			fmt.Printf("Received err: %s \n", err.Error())
			// Stop game and remove player
			// g.Quit_ch <- true
		}

		conn.Mu.Unlock()
	}
}

func BroadcastCollison(g *game.Game, collision game.Collision) {
	players := g.Players

	for _, p := range players {
		conn := p.Connection

		conn.Mu.Lock()

		fmt.Printf("Collision on game %s, writing message: %+v \n", g.Id, collision)

		err := conn.Connection.WriteJSON(CollisionMessage{
			Type:   "COLLISION_MESSAGE",
			XSpeed: collision.XSpeed,
			YSpeed: collision.YSpeed,
		})

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
