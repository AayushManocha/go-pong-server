package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/AayushManocha/go-game-server/bootstrap"
	"github.com/AayushManocha/go-game-server/game"
	"github.com/AayushManocha/go-game-server/messaging"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: FIX LATER
		return true
	},
}

func Echo(w http.ResponseWriter, r *http.Request) {
	origin, err := url.Parse(r.RequestURI)

	if err != nil {
		fmt.Printf("Error parsing origin")
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Upgrade error: ", err.Error())
		return
	}

	gameId := origin.Query().Get("gameId")
	current_game := game.GetGameById(gameId, bootstrap.GetApp().LIVE_GAMES)

	playerIndex := origin.Query().Get("playerIndex")
	parsedPlayerIndex, _ := strconv.Atoi(playerIndex)

	var newPlayer *game.Player
	existingPlayer := game.GetPlayerById(parsedPlayerIndex, current_game.Players)

	if existingPlayer == nil {
		newPlayerIdx := len(current_game.Players) + 1

		if newPlayerIdx > 2 {
			// c.WriteJSON(messaging.NewGenericErrorMessage("This game is already full"))
			return
		}

		newPlayer = game.NewPlayer(newPlayerIdx, current_game.CanvasWidth-(game.DEFAULT_GUTTER_WIDTH+game.DEFAULT_PADDLE_WIDTH))
		newPlayer.SetConnection(c)
		current_game.AddPlayer(newPlayer)
	} else {
		existingPlayer.SetConnection(c)
	}

	if len(current_game.Players) == 2 {
		current_game.GameStatus = game.ParseGameStatus("READY")
	}

	//Optionally write playerMessage
	if newPlayer != nil {
		c.WriteJSON(messaging.NewPlayerJoinedMessage(newPlayer))
		messaging.BroadcastGame(current_game)
	}

	// Write initial game state to new client
	err = c.WriteJSON(messaging.NewGameMessage(current_game))

	if err != nil {
		fmt.Printf("Error writing JSON: %s \n", err.Error())
	}

}
