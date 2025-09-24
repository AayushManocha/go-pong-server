package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var LIVE_GAMES []*Game

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: FIX LATER
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	origin := r.Header["Origin"]
	fmt.Println("Origin header: ", origin)
	fmt.Println("r.host: ", r.Host)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// Create a new game with this connection
	newGame := CreateNewGame()
	newGame.AddConnection(c)

	LIVE_GAMES = append(LIVE_GAMES, &newGame)

	//Broadcast new game to client
	// c.WriteMessage(1, []byte(newGame.toJSON()))
	c.WriteJSON(newGame)

}

type HandlePlayerMoveDTO struct {
	GameId    string `json:"gameId`
	PlayerId  string `json:"playerId"`
	Direction string `json:"direction"`
}

func HandlePlayerMove(w http.ResponseWriter, r *http.Request) {
	var dto HandlePlayerMoveDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	defer r.Body.Close()

	if err != nil {
		panic("Could not read request body")
	}

	fmt.Println(fmt.Sprintf(`dto: %+v`, dto))

	for _, game := range LIVE_GAMES {
		if game.id == dto.GameId {
			game.MovePlayer(dto.PlayerId, dto.Direction)
		}
	}
}

func main() {

	LIVE_GAMES = make([]*Game, 0)

	r := chi.NewRouter()

	r.Use(CorsMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})

	r.Post("/create-game", func(w http.ResponseWriter, r *http.Request) {
		newGame := CreateNewGame()
		LIVE_GAMES = append(LIVE_GAMES, &newGame)
	})

	r.Get("/reset-game", func(w http.ResponseWriter, r *http.Request) {
	})

	r.Get("/echo", echo)
	r.Post("/move-player", HandlePlayerMove)

	http.ListenAndServe(":3000", r)

}
