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
	GameId int `json:"gameId`
}

func GameStart(w http.ResponseWriter, r *http.Request) {
	var dto GameStartDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	defer r.Body.Close()

	if err != nil {
		panic("Could not read request body")
	}

	game := game.GetGameById(dto.GameId, bootstrap.GetApp().LIVE_GAMES)

	playerCount := len(game.Players)

	if playerCount < 2 {
		w.Write([]byte("Not enough players in game"))
		return
	}

	go func() {
		game.GameStatus = "PLAYED"
	gameloop:
		for {
			select {
			case <-game.Quit_ch:
				fmt.Println("Quit on game w/ id: ", game.Id)
				game.GameStatus = "PAUSED"
				messaging.BroadcastUpdates(game)
				break gameloop
			default:
				time.Sleep(time.Millisecond * 50)
				game.MoveBall()
				messaging.BroadcastUpdates(game)
			}
		}
	}()
}
