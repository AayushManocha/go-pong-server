package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AayushManocha/go-game-server/bootstrap"
	"github.com/AayushManocha/go-game-server/game"
	"github.com/AayushManocha/go-game-server/messaging"
)

type HandlePlayerMoveDTO struct {
	GameId    string `json:"gameId`
	PlayerId  int    `json:"playerId"`
	Direction string `json:"direction"`
}

func HandlePlayerMove(w http.ResponseWriter, r *http.Request) {
	var dto HandlePlayerMoveDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	defer r.Body.Close()

	if err != nil {
		panic("Could not read request body")
	}

	game := game.GetGameById(dto.GameId, bootstrap.GetApp().LIVE_GAMES)

	if game.GameStatus == "PLAYED" {
		game.MovePlayer(dto.PlayerId, dto.Direction)
		messaging.BroadcastUpdates(game)
	}

}
