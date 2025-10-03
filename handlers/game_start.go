package handlers

import (
	"encoding/json"
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

	messaging.BroadcastGameStart(g)

	// playerCount := len(game.Players)

	// if playerCount < 2 {
	// 	w.Write([]byte("Not enough players in game"))
	// 	return
	// }
	//

	go func() {
		g.GameStatus = game.ParseGameStatus("IN_PLAY")
		ticksSinceCorrection := 0
	gameloop:
		for {
			select {
			case <-g.Quit_ch:
				if g.Winner != 0 {
					g.GameStatus = game.ParseGameStatus("FINISHED")
					messaging.BroadcastGameWinMessage(g)
					break gameloop
				} else {
					g.GameStatus = game.ParseGameStatus("PAUSED")
					messaging.BroadcastGameStop(g)
					break gameloop
				}

			default:
				time.Sleep(time.Millisecond * 50)
				g.MoveBall(50)
				ticksSinceCorrection += 1
				if ticksSinceCorrection >= 20 {
					// messaging.BroadcastBallCorrection(g)
					ticksSinceCorrection = 0
				}
			}
		}
	}()

}
