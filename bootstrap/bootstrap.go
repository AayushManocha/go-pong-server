package bootstrap

import "github.com/AayushManocha/go-game-server/game"

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
