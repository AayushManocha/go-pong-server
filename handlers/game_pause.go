package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AayushManocha/go-game-server/bootstrap"
	"github.com/AayushManocha/go-game-server/game"
)

type GamePauseDto struct {
	GameId int
}

func GamePause(w http.ResponseWriter, r *http.Request) {
	var dto GamePauseDto
	json.NewDecoder(r.Body).Decode(&dto)
	defer r.Body.Close()

	games := bootstrap.GetApp().LIVE_GAMES
	game := game.GetGameById(dto.GameId, games)
	fmt.Println("GamePause() -> ", game.Id)

	game.Quit_ch <- true

	w.Write([]byte("Done writing to quit channel"))
}
