package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var game Game

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
	defer c.Close()

	ticker := time.NewTicker(500 * time.Millisecond) // every 2 seconds$
	defer ticker.Stop()

	for {

		select {
		case _ = <-ticker.C:
			c.WriteMessage(1, []byte(game.toJSON()))
		}

	}

}

type HandlePlayerMoveDTO struct {
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

	game.MovePlayer(dto.PlayerId, dto.Direction)
}

func main() {

	game = CreateNewGame()

	r := chi.NewRouter()

	r.Use(CorsMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})

	r.Get("/reset-game", func(w http.ResponseWriter, r *http.Request) {
		game = CreateNewGame()
	})

	r.Get("/echo", echo)
	r.Post("/move-player", HandlePlayerMove)

	http.ListenAndServe(":3000", r)

}
