package main

import (
	"fmt"

	"github.com/goccy/go-json"
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

type NewObject struct {
Id string
number int
}

type Player struct {
	Id    string    `json:"id"`
	Shape Rectangle `json:"shape"`
}

type Game struct {
	PLAYERS       []Player `json:"players"`
	CANVAS_HEIGHT int      `json:"canvasHeight"`
	CANVAS_WIDTH  int      `json:"canvasWidth"`
	// BALL       Rectangle
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
	game := Game{
		PLAYERS:       []Player{},
		CANVAS_HEIGHT: 500,
		CANVAS_WIDTH:  500,
	}
	game.PLAYERS = append(game.PLAYERS, Player{Id: "1", Shape: Rectangle{X: 10, Y: 10}})
	game.PLAYERS = append(game.PLAYERS, Player{Id: "2", Shape: Rectangle{X: 100, Y: 10}})

	return game
}

func (g *Game) MovePlayer(p *Player, direction string) {
	fmt.Println("MovePlayer()")
	if direction == "UP" && p.Shape.Y < g.CANVAS_HEIGHT {
		fmt.Println("Move Player Up")
		p.Shape.Y += 1
	} else if p.Shape.Y > 0 {
		fmt.Println("Move Player down")
		p.Shape.Y -= 1
	}
}
