package game

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Player struct {
	Index      int               `json:"index"`
	Shape      *Rectangle        `json:"shape"`
	Connection *PlayerConnection `json:"_"`
}

type PlayerConnection struct {
	Mu         sync.Mutex
	Connection *websocket.Conn
}

func NewPlayer(playerIndex int, xPos int) *Player {
	return &Player{
		Index: playerIndex,
		Shape: &Rectangle{
			X:     xPos,
			Y:     100,
			Width: 20,
		},
		Connection: &PlayerConnection{},
	}
}

func (p *Player) SetConnection(conn *websocket.Conn) {
	p.Connection = &PlayerConnection{
		Connection: conn,
	}
}
