package main

import (
	"fmt"
	"net/http"

	"github.com/AayushManocha/go-game-server/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Use(CorsMiddleware)

	r.Get("/echo", handlers.Echo)
	r.Post("/create-game", handlers.CreateGame)
	r.Post("/move-player", handlers.HandlePlayerMove)
	r.Post("/game-start", handlers.GameStart)
	r.Post("/game-pause", handlers.GamePause)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	fmt.Println("Starting Router...")
	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Println("Error starting router")
	}

}
