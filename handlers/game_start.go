package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AayushManocha/go-game-server/bootstrap"
	"github.com/AayushManocha/go-game-server/game"
	"github.com/AayushManocha/go-game-server/messaging"
)

type GameStartDTO struct {
	GameId string `json:"gameId`
}

func GameStart(w http.ResponseWriter, r *http.Request) {
	var dto GameStartDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	defer r.Body.Close()

	if err != nil {
		panic("Could not read request body")
	}

	g := game.GetGameById(dto.GameId, bootstrap.GetApp().LIVE_GAMES)

	messaging.BroadcastToAllPlayers(g, messaging.NewGameStartMessage())

	playerCount := len(g.Players)
	if playerCount < 2 {
		w.Write([]byte("Not enough players in game"))
		return
	}

	go func() {
		g.GameStatus = game.ParseGameStatus("IN_PLAY")
		ticksSinceCorrection := 0
	gameloop:
		for {
			select {
			case <-g.Quit_ch:
				if g.Winner != 0 {
					g.GameStatus = game.ParseGameStatus("FINISHED")
					messaging.BroadcastToAllPlayers(g, messaging.NewGameWinMessage(g))
					bootstrap.GetApp().RemoveGame(g.Id)
					fmt.Printf("Games remaining: %+v", bootstrap.GetApp().LIVE_GAMES)
					break gameloop
				} else {
					g.GameStatus = game.ParseGameStatus("PAUSED")
					messaging.BroadcastToAllPlayers(g, messaging.NewGameStopMessage(g))
					break gameloop
				}

			default:
				time.Sleep(time.Millisecond * 50)
				g.MoveBall(50)
				ticksSinceCorrection += 1
				if ticksSinceCorrection >= 20 {
					messaging.BroadcastToAllPlayers(g, messaging.NewBallCorectionMessage(g))
					ticksSinceCorrection = 0
				}
			}
		}
	}()

}
