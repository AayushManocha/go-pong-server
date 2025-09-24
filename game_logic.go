package main

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
)

type Rectangle struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (r Rectangle) toJSON() string {
	return fmt.Sprintf(`{"X": "%d", "Y": "%d"}`, r.X, r.Y)
}

func (r Rectangle) String() string {
	fmt.Sprintf("%d", 5)
	return fmt.Sprintf("Rectangle{X: %d, Y: %d}", r.X, r.Y)
}

type Player struct {
	Id    string     `json:"id"`
	Shape *Rectangle `json:"shape"`
}

type Game struct {
	id            string
	PLAYERS       *[]*Player `json:"players"`
	CANVAS_HEIGHT int        `json:"canvasHeight"`
	CANVAS_WIDTH  int        `json:"canvasWidth"`
	// BALL       Rectangle
	CONNECTIONS *[]*websocket.Conn
}

func (p Player) toJSON() string {
	return fmt.Sprintf(`{"id": "%s", "shape": "%s"}`, p.Id, p.Shape.toJSON())
}

func (p Player) String() string {
	return fmt.Sprintf("Player{id: %s, shape: %s}", p.Id, p.Shape)
}

func (g *Game) toJSON() string {
	json, _ := json.Marshal(g)
	return string(json)
}

func (g Game) String() string {
	return fmt.Sprintf(
		"Game{CANVAS_HEIGHT: %d, CANVAS_WIDTH: %d, PLAYERS: %v}",
		g.CANVAS_HEIGHT,
		g.CANVAS_WIDTH,
		g.PLAYERS,
	)
}

func CreateNewGame() Game {
	p1 := &Player{Id: "1", Shape: &Rectangle{X: 10, Y: 10}}
	p2 := &Player{Id: "2", Shape: &Rectangle{X: 100, Y: 10}}
	connections := make([]*websocket.Conn, 0)

	game := Game{
		PLAYERS:       &[]*Player{p1, p2},
		CANVAS_HEIGHT: 500,
		CANVAS_WIDTH:  500,
		CONNECTIONS:   &connections,
	}

	return game
}

func (g *Game) AddConnection(conn *websocket.Conn) {
	if g.CONNECTIONS != nil {
		new_connections := append(*g.CONNECTIONS, conn)
		g.CONNECTIONS = &new_connections
		return
	}

}

func (g *Game) MovePlayer(playerId string, direction string) {
	fmt.Println("MovePlayer()")

	players := *g.PLAYERS

	for _, player := range players {
		if player.Id == playerId {
			if direction == "DOWN" {
				fmt.Println("Down")
				player.Shape.Y += 1
			} else {
				fmt.Println("Up")
				player.Shape.Y -= 1
			}
		}
	}

	g.BroadcastUpdates()

}

func (g *Game) BroadcastUpdates() {
	connections := *g.CONNECTIONS
	// fmt.Println("# of connections: ", len(connections))
	for _, conn := range connections {
		fmt.Println("Remote Address: ", conn.NetConn().RemoteAddr().String())
		conn.WriteJSON(g)
	}
}
