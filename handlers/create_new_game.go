package handlers

import (
	"fmt"
	"net/http"

	"github.com/AayushManocha/go-game-server/bootstrap"
	"github.com/AayushManocha/go-game-server/game"
)

func CreateGame(w http.ResponseWriter, r *http.Request) {

	app := bootstrap.GetApp()

	newGame := game.CreateNewGame()
	app.LIVE_GAMES = append(app.LIVE_GAMES, newGame)

	p1 := game.NewPlayer(1, game.DEFAULT_GUTTER_WIDTH)
	newGame.AddPlayer(p1)

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"gameId": "%s", "playerIndex": %d, "xSpeed": "%d", "ySpeed":"%d"}`, newGame.Id, p1.Index, newGame.Ball.SpeedX, newGame.Ball.SpeedY)))
}
