package game

func GetGameById(gameId int, gameList []*Game) *Game {
	var foundGame *Game
	for _, g := range gameList {
		if g.Id == gameId {
			foundGame = g
		}
	}

	return foundGame
}

func GetPlayerById(playerId int, playerList []*Player) *Player {
	var foundPlayer *Player
	for _, p := range playerList {
		if p.Index == playerId {
			foundPlayer = p
		}
	}

	return foundPlayer
}
