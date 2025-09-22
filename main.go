package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow any domain (you might restrict this in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type HandlePlayerMoveDTO struct {
	PlayerId  string `json:"playerId"`
	Direction string `json:"direction"`
}

func HandlePlayerMove(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("https://example.com")
	var v struct{}
	json.NewDecoder(resp.Body).Decode(&v)

	if r.Method == "POST" {
		var dto HandlePlayerMoveDTO
		err := json.NewDecoder(r.Body).Decode(&dto)

		fmt.Println(fmt.Sprintf(`dto: %+v`, dto))

		if err != nil {
			panic("Could not read request body")
		}
		defer r.Body.Close()

		for _, player := range game.PLAYERS {
			if player.Id == dto.PlayerId {
				game.MovePlayer(&player, dto.Direction)
			}
		}
	}
}

func main() {

	game = CreateNewGame()
	fmt.Println("Game: ", game.toJSON())

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})
	mux.HandleFunc("/echo", echo)

	mux.HandleFunc("/move-player", HandlePlayerMove)

	handler := corsMiddleware(mux)

	http.ListenAndServe(":3000", handler)
}
