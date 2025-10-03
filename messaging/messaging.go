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

type PlayerMoveMessage struct {
	Type        string  `json:"type"`
	PlayerIndex int     `json:"playerIndex"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
}

type GameStartMessage struct {
	Type string `json:"type"`
}

type BallCorrectionMessage struct {
	Type   string `json:"type"`
	SpeedX float64
	SpeedY float64
	X      float64
	Y      float64
}

type GameStopMessage struct {
	Type string `json:"type"`
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

func BroadcastPlayerMove(g *game.Game, movedPlayer *game.Player) {
	players := g.Players

	msg := PlayerMoveMessage{
		Type:        "PLAYER_MOVE_MESSAGE",
		PlayerIndex: movedPlayer.Index,
		X:           movedPlayer.Shape.X,
		Y:           movedPlayer.Shape.Y,
	}

	for _, p := range players {
		conn := p.Connection

		// if p.Index != movedPlayer.Index {
		conn.Mu.Lock()
		err := conn.Connection.WriteJSON(msg)
		if err != nil {
			fmt.Printf("Received err: %s \n", err.Error())
			// Stop game and remove player
			// g.Quit_ch <- true
		}

		conn.Mu.Unlock()
	}

	// }
}

func BroadcastGameStart(g *game.Game) {
	players := g.Players
	for _, p := range players {
		conn := p.Connection
		conn.Mu.Lock()
		err := conn.Connection.WriteJSON(GameStartMessage{
			Type: "GAME_START_MESSAGE",
		})
		if err != nil {
			fmt.Printf("Received err: %s \n", err.Error())
			// Stop game and remove player
			g.Quit_ch <- true
		}
		conn.Mu.Unlock()
	}
}

func BroadcastBallCorrection(g *game.Game) {
	players := g.Players
	for _, p := range players {
		conn := p.Connection
		conn.Mu.Lock()
		err := conn.Connection.WriteJSON(BallCorrectionMessage{
			Type:   "BALL_CORRECTION_MESSAGE",
			SpeedX: g.Ball.SpeedX,
			SpeedY: g.Ball.SpeedY,

			X: g.Ball.Shape.X,
			Y: g.Ball.Shape.Y,
		})

		if err != nil {
			fmt.Printf("Received err: %s \n", err.Error())
			// Stop game and remove player
			g.Quit_ch <- true
		}
		conn.Mu.Unlock()
	}
}

func BroadcastGameStop(g *game.Game) {
	players := g.Players
	for _, p := range players {
		conn := p.Connection
		conn.Mu.Lock()
		err := conn.Connection.WriteJSON(GameStopMessage{
			Type: "GAME_STOP_MESSAGE",
		})
		fmt.Println("GAME_STOP_MESSAGE")
		if err != nil {
			fmt.Printf("Received err: %s \n", err.Error())
			// Stop game and remove player
			g.Quit_ch <- true
		}
		conn.Mu.Unlock()
	}
}
