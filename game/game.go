package game

import (
	"math"

	"github.com/AayushManocha/go-game-server/utils"
)

type Direction int

const (
	UP = iota
	DOWN
)

const DEFAULT_PADDLE_WIDTH = 20
const DEFAULT_BALL_DIAMETER = 25
const DEFAULT_GUTTER_WIDTH = 50

type Game struct {
	Id           string
	Players      []*Player  `json:"players"`
	CanvasHeight int        `json:"canvasHeight"`
	CanvasWidth  int        `json:"canvasWidth"`
	Ball         *Ball      `json:"ball"`
	GameStatus   GameStatus `json:"gameStatus"`
	Winner       int        `json:"winner"`

	Quit_ch chan bool `json:"-"`
}

func CreateNewGame() *Game {
	ball := &Ball{
		Shape:  &Rectangle{X: 500, Y: 250, Width: DEFAULT_BALL_DIAMETER},
		SpeedX: GenerateRandomSpeed(),
		SpeedY: GenerateRandomSpeed(),
	}

	game := Game{
		Id:           utils.String(10),
		Players:      []*Player{},
		CanvasHeight: 500,
		CanvasWidth:  1000,
		Ball:         ball,
		GameStatus:   ParseGameStatus("CREATED"),
		Quit_ch:      make(chan bool, 2),
		Winner:       0,
	}

	return &game
}

func (g *Game) AddPlayer(p *Player) {
	newPlayerList := append(g.Players, p)
	g.Players = newPlayerList
}

func (g *Game) MovePlayer(playerId int, newY float64) *Player {
	players := g.Players
	player := GetPlayerById(playerId, players)
	player.Shape.Y = newY
	return player
}

func (g *Game) MoveBall(milliseconds int) {
	g.Ball.Shape.X += float64(g.Ball.SpeedX) * float64(milliseconds)
	g.Ball.Shape.Y += float64(g.Ball.SpeedY) * float64(milliseconds)

	paddleCollided := detectPaddleCollision(g)
	detectWallCollision(g, paddleCollided)
}

func detectPaddleCollision(g *Game) bool {
	paddleOne := (g.Players)[0].Shape
	paddleTwo := (g.Players)[1].Shape

	ball := g.Ball

	var leftPaddle *Rectangle
	var rightPaddle *Rectangle

	if paddleOne.X < paddleTwo.X {
		leftPaddle = paddleOne
		rightPaddle = paddleTwo
	} else {
		leftPaddle = paddleTwo
		rightPaddle = paddleOne
	}

	paddleOneXHit := ball.Shape.X <= leftPaddle.X+float64(leftPaddle.Width)
	paddleOneYHit := ball.Shape.Y >= leftPaddle.Y && ball.Shape.Y <= leftPaddle.Y+float64(leftPaddle.Height)

	if paddleOneXHit && paddleOneYHit {
		ball.SpeedX = math.Abs(ball.SpeedX)
		return true
	}

	paddleTwoXHit := ball.Shape.X >= rightPaddle.X-float64(rightPaddle.Width)
	paddleTwoYHit := ball.Shape.Y >= rightPaddle.Y && ball.Shape.Y <= rightPaddle.Y+float64(rightPaddle.Height)

	if paddleTwoXHit && paddleTwoYHit {
		ball.SpeedX = math.Abs(ball.SpeedX) * -1
		return true
	}

	return false
}

func (g *Game) SetWinner(playerIndex int) {
	g.Winner = playerIndex
	g.Quit_ch <- true
}

func detectWallCollision(g *Game, paddleCollided bool) {
	ball := g.Ball

	// if !paddleCollided {
	if ball.Shape.X <= 0 {
		g.SetWinner(2)
	} else if ball.Shape.X >= float64(g.CanvasWidth-ball.Shape.Width) {
		g.SetWinner(1)
	}
	// }

	if ball.Shape.Y <= 0 {
		ball.SpeedY *= -1
	} else if ball.Shape.Y >= float64(g.CanvasHeight-ball.Shape.Width) {
		ball.SpeedY *= -1
	}
}
