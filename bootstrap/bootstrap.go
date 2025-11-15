package bootstrap

import (
	"github.com/AayushManocha/go-game-server/game"
)

type Application struct {
	LIVE_GAMES []*game.Game
}

var app *Application

func GetApp() *Application {
	if app == nil {
		app = &Application{
			LIVE_GAMES: []*game.Game{},
		}
	}

	return app
}

func (a *Application) AddGame(g *game.Game) {
	currentGames := a.LIVE_GAMES
	currentGames = append(currentGames, g)
}

func (a *Application) RemoveGame(id string) {
	currentGames := app.LIVE_GAMES
	var index int

	for i, g := range currentGames {
		if g.Id == id {
			index = i
			break
		}
	}

	s1 := currentGames[:index]
	s2 := currentGames[index+1:]

	app.LIVE_GAMES = append(s1, s2...)
}
