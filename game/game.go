package game

import (
	"fmt"

	"github.com/AayushManocha/go-game-server/utils"
)

type GameStatus int

const (
	PAUSED GameStatus = iota
	PLAYED
	CREATED
	FINISHED
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
	Players      []*Player `json:"players"`
	CanvasHeight int       `json:"canvasHeight"`
	CanvasWidth  int       `json:"canvasWidth"`
	Ball         *Ball     `json:"ball"`
	GameStatus   string    `json:"gameStatus"`
	Winner       int       `json:"winner"`

	Quit_ch             chan bool      `json:"-"`
	CollisionMessage_ch chan Collision `json:"-"`
}

type Collision struct {
	XSpeed float64
	YSpeed float64
}

func CreateNewGame() *Game {
	ball := &Ball{
		Shape: &Rectangle{X: 500, Y: 250, Width: DEFAULT_BALL_DIAMETER},
		// SpeedX: GenerateRandomSpeed(),
		SpeedX: 0,
		SpeedY: GenerateRandomSpeed(),
	}

	game := Game{
		Id:                  utils.String(10),
		Players:             []*Player{},
		CanvasHeight:        500,
		CanvasWidth:         1000,
		Ball:                ball,
		GameStatus:          "PAUSED",
		Quit_ch:             make(chan bool, 2),
		Winner:              0,
		CollisionMessage_ch: make(chan Collision, 1024),
	}

	return &game
}

func (g *Game) AddPlayer(p *Player) {
	newPlayerList := append(g.Players, p)
	g.Players = newPlayerList
}

func (g *Game) MovePlayer(playerId int, direction string) *Player {
	players := g.Players
	player := GetPlayerById(playerId, players)

	playerIsAtBottomOfCanvas := player.Shape.Y >= float64(g.CanvasHeight-player.Shape.Height)
	playerIsAtTopOfCanvas := player.Shape.Y <= 0

	if direction == "DOWN" && !playerIsAtBottomOfCanvas {
		fmt.Println("Moved Down")
		player.Shape.Y += 10
	} else if direction == "UP" && !playerIsAtTopOfCanvas {
		fmt.Println("Moved UP")
		player.Shape.Y -= 10
	}
	return player
}

func (g *Game) MoveBall(milliseconds int) {
	g.Ball.Shape.X += float64(g.Ball.SpeedX) * float64(milliseconds)
	g.Ball.Shape.Y += float64(g.Ball.SpeedY) * float64(milliseconds)

	// fmt.Printf("New X: %.2f New Y: %.2f (changed x by %.2f) \n", g.Ball.Shape.X, g.Ball.Shape.Y, float64(g.Ball.SpeedX)/1000.0)

	detectWallCollision(g)
	// detectPaddleCollision(g)
}

func detectPaddleCollision(g *Game) {
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

	paddleOneXHit := ball.Shape.X <= leftPaddle.X+20
	paddleOneYHit := ball.Shape.Y >= leftPaddle.Y && ball.Shape.Y <= leftPaddle.Y+100

	if paddleOneXHit && paddleOneYHit {
		g.CollisionMessage_ch <- Collision{
			XSpeed: ball.SpeedX * -1,
			YSpeed: ball.SpeedY,
		}
		ball.SpeedX *= -1
	}

	paddleTwoXHit := ball.Shape.X >= rightPaddle.X-20
	paddleTwoYHit := ball.Shape.Y >= rightPaddle.Y && ball.Shape.Y <= rightPaddle.Y+100

	if paddleTwoXHit && paddleTwoYHit {
		g.CollisionMessage_ch <- Collision{
			XSpeed: ball.SpeedX * -1,
			YSpeed: ball.SpeedY,
		}
		ball.SpeedX *= -1
	}

}

func (g *Game) SetWinner(playerIndex int) {
	g.Winner = playerIndex
	g.Quit_ch <- true
}

func detectWallCollision(g *Game) {
	ball := g.Ball
	if ball.Shape.X <= 0 {
		g.CollisionMessage_ch <- Collision{
			XSpeed: 0,
			YSpeed: 0,
		}
		g.SetWinner(2)
	} else if ball.Shape.X >= float64(g.CanvasWidth-ball.Shape.Width) {
		g.CollisionMessage_ch <- Collision{
			XSpeed: 0,
			YSpeed: 0,
		}
		g.SetWinner(1)
	}

	if ball.Shape.Y <= 0 {
		g.CollisionMessage_ch <- Collision{
			XSpeed: ball.SpeedX,
			YSpeed: ball.SpeedY * -1,
		}
		ball.SpeedY *= -1
	} else if ball.Shape.Y >= float64(g.CanvasHeight-ball.Shape.Width) {
		g.CollisionMessage_ch <- Collision{
			XSpeed: ball.SpeedX,
			YSpeed: ball.SpeedY * -1,
		}
		ball.SpeedY *= -1
	}
}
