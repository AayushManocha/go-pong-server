package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AayushManocha/go-game-server/bootstrap"
	"github.com/AayushManocha/go-game-server/game"
	"github.com/AayushManocha/go-game-server/messaging"
)

type HandlePlayerMoveDTO struct {
	GameId   string  `json:"gameId`
	PlayerId int     `json:"playerId"`
	NewY     float64 `json:"newY"`
}

func HandlePlayerMove(w http.ResponseWriter, r *http.Request) {
	var dto HandlePlayerMoveDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	defer r.Body.Close()

	if err != nil {
		panic("Could not read request body")
	}

	game := game.GetGameById(dto.GameId, bootstrap.GetApp().LIVE_GAMES)

	updatedPlayer := game.MovePlayer(dto.PlayerId, dto.NewY)
	messaging.BroadcastToOtherPlayers(updatedPlayer, game, messaging.NewPlayerMoveMessage(updatedPlayer))

	w.Write([]byte("Moved player"))
}
